export const CONFIRMATION_DIALOG_TEMPLATE: string = `
<clr-modal [(clrModalOpen)]="opened" [clrModalClosable]="false" [clrModalStaticBackdrop]="true">
    <h3 class="modal-title" class="confirmation-title" style="margin-top: 0px;">{{dialogTitle}}</h3>
    <div class="modal-body">
        <div class="confirmation-icon-inline">
            <clr-icon shape="warning" class="is-warning" size="64"></clr-icon>
        </div>
        <div class="confirmation-content">{{dialogContent}}</div>
    </div>
    <div class="modal-footer" [ngSwitch]="buttons">
       <ng-template [ngSwitchCase]="0">
        <button type="button" class="btn btn-outline" (click)="cancel()">{{'BUTTON.CANCEL' | translate}}</button>
        <button type="button" class="btn btn-primary" (click)="confirm()">{{ 'BUTTON.CONFIRM' | translate}}</button>
       </ng-template>
       <ng-template [ngSwitchCase]="1">
        <button type="button" class="btn btn-outline" (click)="cancel()">{{'BUTTON.NO' | translate}}</button>
        <button type="button" class="btn btn-primary" (click)="confirm()">{{ 'BUTTON.YES' | translate}}</button>
       </ng-template>
       <ng-template [ngSwitchCase]="2">
        <button type="button" class="btn btn-outline" (click)="cancel()">{{'BUTTON.CANCEL' | translate}}</button>
        <button type="button" class="btn btn-danger" (click)="confirm()">{{ 'BUTTON.DELETE' | translate}}</button>
       </ng-template>
       <ng-template [ngSwitchCase]="3">
        <button type="button" class="btn btn-primary" (click)="cancel()">{{'BUTTON.CLOSE' | translate}}</button>
       </ng-template>
    </div>
</clr-modal>
`;