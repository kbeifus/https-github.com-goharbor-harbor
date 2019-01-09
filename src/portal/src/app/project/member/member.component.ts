
import { finalize } from 'rxjs/operators';
// Copyright (c) 2017 VMware, Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
import { Component, OnInit, ViewChild, OnDestroy, ChangeDetectionStrategy, ChangeDetectorRef } from "@angular/core";
import { ActivatedRoute, Router } from "@angular/router";
import { Subscription, forkJoin } from "rxjs";
import { TranslateService } from "@ngx-translate/core";
import { operateChanges, OperateInfo, OperationService, OperationState } from "@harbor/ui";

import { MessageHandlerService } from "../../shared/message-handler/message-handler.service";
import { ConfirmationTargets, ConfirmationState, ConfirmationButtons } from "../../shared/shared.const";
import { ConfirmationDialogService } from "../../shared/confirmation-dialog/confirmation-dialog.service";
import { ConfirmationMessage } from "../../shared/confirmation-dialog/confirmation-message";
import { SessionService } from "../../shared/session.service";
import { RoleInfo } from "../../shared/shared.const";
import { Project } from "../../project/project";
import { Member } from "./member";
import { SessionUser } from "../../shared/session-user";
import { AddGroupComponent } from './add-group/add-group.component';
import { MemberService } from "./member.service";
import { AddMemberComponent } from "./add-member/add-member.component";
import { AppConfigService } from "../../app-config.service";
import { UserPermissionService, USERSTATICPERMISSION, ErrorHandler } from "@harbor/ui";

@Component({
  templateUrl: "member.component.html",
  styleUrls: ["./member.component.scss"],
  changeDetection: ChangeDetectionStrategy.OnPush
})
export class MemberComponent implements OnInit, OnDestroy {

  members: Member[];
  projectId: number;
  roleInfo = RoleInfo;
  delSub: Subscription;

  currentUser: SessionUser;

  batchOps = 'delete';
  searchMember: string;
  selectedRow: Member[] = [];
  roleNum: number;
  isDelete = false;
  isChangeRole = false;
  loading = false;
  isLdapMode: boolean = false;

  isChangingRole = false;
  batchChangeRoleInfos = {};

  @ViewChild(AddMemberComponent)
  addMemberComponent: AddMemberComponent;

  @ViewChild(AddGroupComponent)
  addGroupComponent: AddGroupComponent;
  hasCreateMemberPermission: boolean;
  hasUpdateMemberPermission: boolean;
  hasDeleteMemberPermission: boolean;
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private memberService: MemberService,
    private translate: TranslateService,
    private messageHandlerService: MessageHandlerService,
    private OperateDialogService: ConfirmationDialogService,
    private session: SessionService,
    private operationService: OperationService,
    private appConfigService: AppConfigService,
    private userPermissionService: UserPermissionService,
    private errorHandler: ErrorHandler,
    private ref: ChangeDetectorRef) {

    this.delSub = OperateDialogService.confirmationConfirm$.subscribe(message => {
      if (message &&
        message.state === ConfirmationState.CONFIRMED &&
        message.source === ConfirmationTargets.PROJECT_MEMBER) {
        if (this.batchOps === 'delete') {
          this.deleteMembers(message.data);
        }
      }
    });
    let hnd = setInterval(() => ref.markForCheck(), 100);
    setTimeout(() => clearInterval(hnd), 1000);
  }

  ngOnDestroy() {
    if (this.delSub) {
      this.delSub.unsubscribe();
    }
  }

  ngOnInit() {
    // Get projectId from route params snapshot.
    this.projectId = +this.route.snapshot.parent.params["id"];
    // Get current user from registered resolver.
    this.currentUser = this.session.getCurrentUser();
    this.retrieve(this.projectId, "");
    if (this.appConfigService.isLdapMode()) {
      this.isLdapMode = true;
    }
    // get member permission rule
    this.getMemberPermissionRule(this.projectId);
  }

  doSearch(searchMember: string) {
    this.searchMember = searchMember;
    this.retrieve(this.projectId, this.searchMember);
  }

  refresh() {
    this.retrieve(this.projectId, "");
  }

  retrieve(projectId: number, username: string) {
    this.loading = true;
    this.selectedRow = [];
    this.memberService
      .listMembers(projectId, username).pipe(
        finalize(() => this.loading = false))
      .subscribe(
        response => {
          this.members = response;
          let hnd = setInterval(() => this.ref.markForCheck(), 100);
          setTimeout(() => clearInterval(hnd), 1000);
        },
        error => {
          this.router.navigate(["/harbor", "projects"]);
          this.messageHandlerService.handleError(error);
        });
  }

  get onlySelf(): boolean {
    if (this.selectedRow.length === 1 &&
      this.selectedRow[0].entity_type === 'u' &&
      this.selectedRow[0].entity_id === this.currentUser.user_id) {
      return true;
    }
    return false;
  }

  member_type_toString(user_type: string) {
    if (user_type === 'u') {
      return 'MEMBER.USER_TYPE';
    } else {
      return 'MEMBER.GROUP_TYPE';
    }
  }

  // Add member
  openAddMemberModal() {
    this.addMemberComponent.openAddMemberModal();
  }

  addedMember(result: boolean) {
    this.searchMember = "";
    this.retrieve(this.projectId, "");
  }

  // Add group
  openAddGroupModal() {
    this.addGroupComponent.open();
  }
  addedGroup(result: boolean) {
    this.searchMember = "";
    this.retrieve(this.projectId, "");
  }

  changeMembersRole(members: Member[], roleId: number) {
    if (!members) {
      return;
    }

    let changeOperate = (projectId: number, member: Member, ) => {
      return this.memberService
        .changeMemberRole(projectId, member.id, roleId)
        .then(() => this.batchChangeRoleInfos[member.id] = 'done')
        .catch(error => this.messageHandlerService.handleError(error + ": " + member.entity_name));
    };

    // Preparation for members role change
    this.batchChangeRoleInfos = {};
    let RoleChangePromises: Promise<any>[] = [];
    members.forEach(member => {
      if (member.entity_type === 'u' && member.entity_id === this.currentUser.user_id) {
        return;
      }
      this.batchChangeRoleInfos[member.id] = 'pending';
      RoleChangePromises.push(changeOperate(this.projectId, member));
    });

    Promise.all(RoleChangePromises).then(() => {
      this.retrieve(this.projectId, "");
    });
  }

  ChangeRoleOngoing(entity_id: number) {
    return this.batchChangeRoleInfos[entity_id] === 'pending';
  }

  // Delete members
  openDeleteMembersDialog(members: Member[]) {
    this.batchOps = 'delete';
    let nameArr: string[] = [];
    if (members && members.length) {
      members.forEach(data => {
        nameArr.push(data.entity_name);
      });
      let deletionMessage = new ConfirmationMessage(
        "MEMBER.DELETION_TITLE",
        "MEMBER.DELETION_SUMMARY",
        nameArr.join(","),
        members,
        ConfirmationTargets.PROJECT_MEMBER,
        ConfirmationButtons.DELETE_CANCEL
      );
      this.OperateDialogService.openComfirmDialog(deletionMessage);
    }
  }

  deleteMembers(members: Member[]) {
    if (!members) { return; }
    let memberDeletingPromises: Promise<any>[] = [];

    // Function to delete specific member
    let deleteMember = (projectId: number, member: Member) => {
      let operMessage = new OperateInfo();
      operMessage.name = 'OPERATION.DELETE_MEMBER';
      operMessage.data.id = member.id;
      operMessage.state = OperationState.progressing;
      operMessage.data.name = member.entity_name;

      this.operationService.publishInfo(operMessage);
      if (member.entity_type === 'u' && member.entity_id === this.currentUser.user_id) {
        this.translate.get("BATCH.DELETED_FAILURE").subscribe(res => {
          operateChanges(operMessage, OperationState.failure, res);
        });
        return null;
      }

      return this.memberService
        .deleteMember(projectId, member.id)
        .then(response => {
          this.translate.get("BATCH.DELETED_SUCCESS").subscribe(res => {
            operateChanges(operMessage, OperationState.success);
          });
        })
        .catch(error => {
          this.translate.get("BATCH.DELETED_FAILURE").subscribe(res => {
            operateChanges(operMessage, OperationState.failure, res);
          });
        });
    };

    // Deleting member then wating for results
    members.forEach(member => memberDeletingPromises.push(deleteMember(this.projectId, member)));

    Promise.all(memberDeletingPromises).then(() => {
      this.selectedRow = [];
      this.batchOps = 'idle';
      this.retrieve(this.projectId, "");
    });
  }
  getMemberPermissionRule(projectId: number): void {
    let hasCreateMemberPermission = this.userPermissionService.getPermission(projectId,
      USERSTATICPERMISSION.MEMBER.KEY, USERSTATICPERMISSION.MEMBER.VALUE.CREATE);
    let hasUpdateMemberPermission = this.userPermissionService.getPermission(projectId,
      USERSTATICPERMISSION.MEMBER.KEY, USERSTATICPERMISSION.MEMBER.VALUE.UPDATE);
    let hasDeleteMemberPermission = this.userPermissionService.getPermission(projectId,
      USERSTATICPERMISSION.MEMBER.KEY, USERSTATICPERMISSION.MEMBER.VALUE.DELETE);
    forkJoin(hasCreateMemberPermission, hasUpdateMemberPermission, hasDeleteMemberPermission).subscribe(MemberRule => {
      this.hasCreateMemberPermission = MemberRule[0] as boolean;
      this.hasUpdateMemberPermission = MemberRule[1] as boolean;
      this.hasDeleteMemberPermission = MemberRule[2] as boolean;
    }, error => this.errorHandler.error(error));
  }
}
