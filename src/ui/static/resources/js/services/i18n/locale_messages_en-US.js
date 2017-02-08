/*
    Copyright (c) 2016 VMware, Inc. All Rights Reserved.
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at
        
        http://www.apache.org/licenses/LICENSE-2.0
        
    Unless required by applicable law or agreed to in writing, software
    distributed under the License is distributed on an "AS IS" BASIS,
    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
    See the License for the specific language governing permissions and
    limitations under the License.
*/
var locale_messages = {
  'sign_in':  'Sign In',
  'sign_up': 'Sign Up',
  'forgot_password': 'Forgot Password',
  'login_now': 'Login Now',
  'its_easy_to_get_started': 'It\'s easy to get started ...',
  'icon_label_1': 'Anonymous repository access',
  'icon_label_2': 'Repositories managed by project',
  'icon_label_3': 'Role based access control',
  'why_use_harbor': 'Why use Harbor?',
  'index_desc': 'Project Harbor is an enterprise-class registry server, which extends the open source Docker Registry server by adding the functionality usually required by an enterprise, such as security, control, and management. Harbor is primarily designed to be a private registry - providing the needed security and control that enterprises require. It also helps minimize bandwidth usage, which is helpful to both improve productivity as well as performance.',
  'index_desc_1': '<strong>Security:</strong> Keep their intellectual properties within their organizations.',
  'index_desc_2': '<strong>Efficiency:</strong> A private registry server is set up within the organization\'s network and can reduce significantly the internet traffic to the public service. ',
  'index_desc_3': '<strong>Access Control:</strong> RBAC (Role Based Access Control) is provided. User management can be integrated with existing enterprise identity services like AD/LDAP. ',
  'index_desc_4': '<strong>Audit:</strong> All access to the registry are logged and can be used for audit purpose.',
  'index_desc_5': '<strong>GUI:</strong> User friendly single-pane-of-glass management console.',
  'index_desc_6': '<strong>Image Replication:</strong> Replicate images between instances.',
  'view_all': 'View all...',
  'repositories': 'Repositories',
  'project_repo_name': 'Project/Repository Name',
  'creation_time': 'Creation Time',
  'author': 'Author',
  'username': 'Username',
  'username_is_required': 'Username is required.',
  'username_has_been_taken': 'Username has been taken.',
  'username_is_too_long': 'Username is too long. (maximum 20 characters)',
  'username_contains_illegal_chars': 'Username contains illegal character(s).',
  'email': 'Email',
  'email_desc': 'The Email address will be used for resetting password.',
  'email_is_required': 'Email is required.',
  'email_has_been_taken': 'Email has been taken.',
  'email_content_illegal': 'Email format is illegal.',	
  'email_does_not_exist': 'Email does not exist.',
  'email_is_too_long': 'Email is to long. (maximum 50 characters)',
  'full_name': 'Full Name',
  'full_name_desc': 'First name & Last name',
  'full_name_is_required': 'Full name is required.',
  'full_name_is_too_long': 'Full name is too long. (maximum 20 characters)',
  'full_name_contains_illegal_chars': 'Full name contains illegal character(s).', 
  'password': 'Password',
  'password_desc': 'At least 8 characters, less than 20 characters with 1 lowercase letter, 1 capital letter and 1 numeric character.',
  'password_is_required': 'Password is required.',
  'password_is_invalid': 'Password is invalid. At least 8 characters, less than 20 characters with 1 lowercase letter, 1 capital letter and 1 numeric character.', 
  'confirm_password': 'Confirm Password',
  'password_does_not_match': 'Passwords do not match.',
  'comments': 'Comments',
  'comment_is_too_long': 'Comment is too long. (maximum 20 characters)',
  'forgot_password_description': 'Please input the Email used when you signed up, a reset password Email will be sent to you.',
  'reset_password': 'Reset Password',
  'successful_reset_password': 'Password has been reset successfully.',
  'failed_to_change_password': 'Failed to change password.',
  'summary': 'Summary',
  'projects': 'Projects',
  'public_projects': 'Public Projects', 
  'public': 'Public',
  'public_repositories': 'Public Repositories',
  'my_project_count': 'My Projects',
  'my_repo_count': 'My Repositories',
  'public_project_count': 'Public Projects',
  'public_repo_count': 'Public Repositories',
  'total_project_count': 'Total Projects',
  'total_repo_count': 'Total Repositories',
  'top_10_repositories': 'Top 10 Repositories',
  'repository_name': 'Repository Name',
  'size': 'Size',
  'count': 'Downloads',
  'creator': 'Creator',
  'no_top_repositories': 'No data, start with Harbor now!',
  'logs': 'Logs',
  'task_name': 'Task Name',
  'details': 'Details',
  'user': 'User',
  'no_user_logs': 'No data, start with Harbor now!',
  'users': 'Users',
  'my_projects': 'My Projects', 
  'project_name': 'Project Name',
  'role': 'Role',
  'publicity': 'Publicity',
  'button_on': 'On',
  'button_off': 'Off',
  'new_project': 'New Project', 
  'save': 'Save',
  'cancel': 'Cancel',
  'confirm': 'Confirm',
  'total': 'Total',
  'items': 'item(s)',  
  'add_member': 'Add Member',
  'operation': 'Operation',
  'advanced_search': 'Advanced Search',
  'all': 'All',
  'others': 'Others',
  'search': 'Search',
  'duration': 'Duration',
  'from': 'From',
  'to': 'To',
  'timestamp': 'Timestamp',
  'dashboard': 'Dashboard',
  'admin_options': 'Admin Options',
  'account_setting': 'Account Settings',
  'log_out': 'Log Out',
  'registration_time': 'Registration Time',
  'system_management': 'System Management', 
  'change_password': 'Change Password',
  'search_result': 'Search Result', 
  'old_password': 'Old Password',
  'old_password_is_required': 'Old password is required.',
  'old_password_is_incorrect': 'Old password is incorrect.',
  'new_password_is_required': 'New password is required.',
  'new_password_is_invalid': 'New password is invalid. At least 8 characters, less than 20 characters with 1 lowercase letter, 1 capital letter and 1 numeric character.',
  'new_password': 'New Password',
  'username_already_exist': 'Username already exist.',
  'username_does_not_exist': 'Username does not exist.',
  'username_or_password_is_incorrect': 'Username or password is incorrect',
  'username_and_password_are_required': 'Both username and password are required.',
  'username_email': 'Username/Email',
  'project_name_is_required': 'Project name is required',
  'project_already_exist': 'Project already exist',
  'project_name_is_invalid': 'Project name is invalid, it should be all lowercase and with no space.',
  'project_name_is_too_short': 'Project name is too short, it should be greater than 4 characters.',
  'project_name_is_too_long': 'Project name is too long, it should be less than 30 characters.',
  'search_projects_or_repositories': 'Search projects or repositories',
  'tag': 'Tag',  
  'image_details': 'Image Details',
  'pull_command': 'Pull Command',
  'alert_delete_repo_title': 'Confirm Deletion',
  'alert_delete_repo': 'All tags under this repository will be deleted. ' +
      'The space of this repository will be recycled during garbage collection.<br/>' +
      '<br/>Delete repository "$0" now?',
  'alert_delete_tag_title': 'Confirm Deletion',
  'alert_delete_tag': 'Note: All tags under this repository will be deleted if they are pointing to this image.<br/><br/>Delete tag "$0" now?',
  'alert_delete_selected_tag': 'Note: All selected tags under this repository will be deleted if they are pointing to this image.<br/><br/>Delete selected tags now?',
  'close': 'Close',
  'ok': 'OK',
  'welcome': 'Welcome to Harbor!',
  'continue' : 'Continue',
  'no_projects_add_new_project': 'No projects available now.',
  'no_repositories': 'No repositories found, please use "docker push" to upload images.',
  'failed_to_add_member': 'Project member can not be added, insuffient permissions.',
  'failed_to_change_member': 'Project member can not be changed, insuffient permissions.',
  'failed_to_delete_member': 'Project member can not be deleted, insuffient permissions.',
  'failed_to_delete_project': 'Project contains repositories or replication policies can not be deleted.', 
  'failed_to_delete_project_insuffient_permissions': 'Project can not be deleted, insuffient permissions.',
  'confirm_delete_project_title': 'Project Deletion',
  'confirm_delete_project': 'Are you sure to delete the project "$0" ?',
  'confirm_delete_user_title': 'User Deletion',
  'confirm_delete_user': 'Are you sure to delete the user "$0" ?',
  'confirm_delete_policy_title': 'Replication Policy Deletion',
  'confirm_delete_policy': 'Are you sure to delete the replication policy "$0" ?',
  'confirm_delete_destination_title': 'Destination Deletion',
  'confirm_delete_destination': 'Are you sure to delete the destination "$0" ?',
  'replication': 'Replication',
  'name': 'Name',
  'description': 'Description',
  'destination': 'Destination',
  'start_time': 'Start Time',
  'last_start_time': 'Last Start Time',
  'end_time': 'End Time',
  'activation': 'Activation',
  'replication_jobs': 'Replication Jobs',
  'actions': 'Actions',
  'status': 'Status',
  'logs' : 'Logs',
  'enabled': 'Enabled',
  'enable': 'Enable',
  'disabled': 'Disabled',
  'disable': 'Disable',
  'no_replication_policies_add_new': 'No replication policies, please add new policy.',
  'no_replication_policies': 'No replication policies.',
  'no_replication_jobs': 'No replication jobs.',
  'no_destinations': 'No destinations, please add new destination.',
  'name_is_required': 'Name is required.',
  'name_is_too_long': 'Name is too long. (maximum 20 characters)',
  'description_is_too_long': 'Description is too long. ',
  'general_setting': 'General',
  'destination_setting': 'Destination Settings',
  'endpoint': 'Destination URL',
  'endpoint_is_required': 'Destination URL is required.',
  'test_connection': 'Test connection',
  'add_new_destination': 'New Destination',
  'edit_destination': 'Edit Destination',
  'successful_changed_password': 'Password has been changed successfully.',
  'change_profile': 'Change Profile',
  'successful_changed_profile': 'User profile has been changed successfully.',
  'administrator': 'Administrator',
  'popular_repositories': 'Popular Repositories',
  'harbor_intro_title': 'About Harbor',
  'mail_has_been_sent': 'Password resetting Email has been sent.',
  'send': 'Send',
  'successful_signed_up': 'Signed up successfully.',
  'add_new_policy': 'Add New Policy',
  'edit_policy': 'Edit Policy',
  'delete_policy': 'Delete Policy',
  'add_new_title': 'Add User',
  'add_new': 'Add',
  'successful_added': 'New user added successfully.',
  'copyright': 'Copyright',
  'all_rights_reserved': 'All Rights Reserved.',
  'pinging_target': 'Testing connection ...',
  'successful_ping_target': 'Connection tested successfully.',
  'failed_to_ping_target': 'Connetion test failed, please check your settings.',
  'policy_already_exists': 'Policy already exists.',
  'destination_already_exists': 'Destination already exists.',
  'refresh': 'Refresh',
  'select_all': 'Select All',
  'delete_tag': 'Delete Tag',
  'delete_repo': 'Delete Repo',
  'delete_selected_tag': 'Delete Selected Tag(s)',
  'download_log': 'View Logs',
  'edit': 'Edit',
  'delete': 'Delete',
  'transfer': 'Transfer',
  'all': 'All',
  'pending': 'Pending',
  'running': 'Running',
  'finished': 'Finished',
  'canceled': 'Canceled',
  'stopped': 'Stopped',
  'retrying': 'Retrying',
  'error': 'Error',
  'about': 'About',
  'about_harbor': 'About Harbor',
  'current_version': '<label>Version:</label>&nbsp;&nbsp;<span>$0</span>',
  'current_storage': '<label>Storage:</label>&nbsp;&nbsp;<span>$0 GB available of $1 GB.</span>',
  'default_root_cert': '<label>Default Root Cert.:</label>&nbsp;&nbsp;<span><a href="$0">$1</a></span>',
  'download': 'Download',
  'failed_to_get_project_member': 'Failed to get current project member.',
  'failed_to_delete_repo': 'Failed to delete repository.',
  'failed_to_delete_repo_insuffient_permissions': 'Failed to delete repository: insuffient permissions.',
  'failed_to_get_repo': 'Failed to get repositories.',
  'failed_to_get_tag': 'Failed to get tag.',
  'failed_to_get_log': 'Failed to get logs.',
  'failed_to_get_project': 'Failed to get projects.',
  'failed_to_update_user': 'Failed to update user.',
  'failed_to_get_stat': 'Failed to get stat data.',
  'failed_to_get_top_repo': 'Failed to get top repositories.',
  'failed_to_get_user_log': 'Failed to get user logs.',
  'failed_to_send_email': 'Failed to send email.',
  'failed_to_reset_pasword': 'Failed to reset password.',
  'failed_in_search': 'Failed in search.',
  'failed_to_sign_up': 'Failed to sign up.',
  'failed_to_add_user': 'Failed to add user.',
  'failed_to_delete_user': 'Failed to delete user.',
  'failed_to_list_user': 'Failed to list user data.',
  'failed_to_toggle_admin': 'Failed to change admin role.',
  'failed_to_list_destination': 'Failed to list destinations.',
  'failed_to_list_replication': 'Failed to list replication policies.',
  'failed_to_toggle_policy': 'Failed to change status of replication policy.',
  'failed_to_create_replication_policy': 'Failed to create replication policy.',
  'failed_to_get_destination': 'Failed to get destination.',
  'failed_to_get_destination_policies': 'Failed to get policies of the destination.',
  'failed_to_get_replication_policy': 'Failed to get replication policy.',
  'failed_to_update_replication_policy': 'Failed to update replication policy.',
  'failed_to_delete_replication_enabled': 'Cannot delete policy: policy has unfinished job(s) or policy is enabled.',
  'failed_to_delete_replication_policy': 'Failed to delete replication policy.',
  'failed_to_delete_destination': 'Failed to delete destination.',
  'failed_to_create_destination': 'Failed to create destination.',
  'failed_to_update_destination': 'Failed to update destination.',
  'failed_to_toggle_publicity_insuffient_permissions': 'Failed to change project publicity: insuffient permissions.',
  'failed_to_toggle_publicity': 'Failed to change project publicity.',
  'failed_to_sign_in': 'Failed to sign in.',
  'project_does_not_exist': 'Project does not exist.',
  'project_admin': 'Project Admin',
  'developer': 'Developer',
  'guest': 'Guest',
  'inline_help_role_title': '<strong>The Definition of Roles</strong>',
  'inline_help_role': '<strong>Project Admin</strong>: Project Admin has read/write and member management privileges to the project.<br/>' +
      '<strong>Developer</strong>: Developer has read and write privileges to the project.<br/>' +
      '<strong>Guest</strong>: Guest has read-only privilege for a specified project.',
  'inline_help_publicity_title': '<strong>Publicity of Project</strong>',
  'inline_help_publicity': 'When a project is set to public, anyone has read permission to the repositories under this project, and the user does not need to run "docker login" before pulling images under this project.',
  'alert_job_contains_error': 'Found errors in the replication job(s), please check.',
  'found_error_in_replication_job': 'Found $0 error(s).',
  'caution': 'Caution',
  'confirm_to_toggle_enabled_policy_title': 'Enable Policy',
  'confirm_to_toggle_enabled_policy': 'After enabling the replication policy, all repositories under the project will be replicated to the destination registry. Please confirm to continue.',
  'confirm_to_toggle_disabled_policy_title': 'Disable Policy',
  'confirm_to_toggle_disabled_policy': 'After disabling the policy, all unfinished replication jobs of this policy will be stopped and canceled. Please confirm to continue.',
  'begin_date_is_later_than_end_date': 'Begin date should not be later than end date.'
};