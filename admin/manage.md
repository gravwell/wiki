# Managing User-Created Objects

Users can create a variety of objects within the Gravwell system:

* Resources
* Saved/backgrounded searches
* Scheduled searches/scripts
* Dashboards
* Templates
* Userfiles

At this time, there are no GUI utilities for managing these objects as the administrator. However, the [Gravwell command-line client](#!cli/cli.md) can list, delete, and modify all of these object types using options in the **admin** sub-menu.

To access these management options, run the client, log in as an administrator user, and enter the admin menu:

```
$ ./client -s gravwell.example.org
Username:  admin
Password:  
#>  admin
admin>  help
add_user            Add a new user
impersonate_user    Impersonate an existing users
del_user            Delete an existing user
get_user            Get an existing users details
update_user         Update an existing user
list_users          List all users
lock_user           Lock a user account
user_activity       Show a specific users activity
user_sessions       Show all open sessions
change_user_pwd     Change a users password
change_admin        Set a users admin status
add_group           Create a new group
del_group           Delete an existing group
list_groups         Lists all existing groups
list_group_users    Lists all members of an existing group
update_group        Update an existing group
add_users_group     Add users to an existing group
del_users_group     Delete users from an existing group
add_user_groups     Add user to existing groups
del_user_groups     Delete a user from groups
get_log_level       Get the webservers current logging level
set_log_level       Set the webservers current logging level
all_dashboards      Get all dashboards for all users
del_dashboard       Delete a dashboard owned by another user
license_info        Display license information
license_sku         Display license SKU
license_serial      Display license Serial Number
license_update      Upload a new license
list_queries        List all queries (active and saved) for all users
delete_queries      Delete any query (active or saved) for any user
list_users_storage  List all users current storage usage
add_indexer         Add another indexer to the configuration
list_extractions    List installed autoextractors
add_extraction      Add a new autoextractor
delete_extraction   Delete an installed autoextractor
update_extraction   Update an installed autoextractor
sync_extractions    Force a sync of installed autoextractors to indexers
resource            Create and manage resources
scheduled_search    Manage scheduled searches
templates           Manage templates
pivots              Manage pivots
userfiles           Manage user files
kits                Manage and upload kits
admin>
```

The rest of this section will briefly describe management options for each object type.

## Managing Dashboards

To list all dashboards on the system, from the **admin** menu run the `all_dashboards` command.

To delete a dashboard, run the `del_dashboard` command from the **admin** menu.

## Managing Searches

To list all searches on the system (saved, backgrounded, or active), run the `list_queries` command from the **admin** menu.

To delete a query, run the `delete_queries` command.

## Managing Resources

The admin sub-menu contains its own sub-menu for managing resources with commands mirroring those available in the regular resource menu:

```
admin>  resource
resource>  help
list                	List available resources
create              	Create a new resource
update              	Upload new data to a resource
delete              	Delete a resource
updatemeta          	Update resource metadata
resource>  
```

From this menu, the administrator can list *all* resources on the system, modify a resource's contents, change its name/description/ownership, or delete it.

## Managing Scheduled Searches

The admin sub-menu contains its own sub-menu for managing scheduled searches:

```
admin>  scheduled_search
scheduled search>  help
list                	List saved searches
listall             	List all saved searches
create              	Create a new scheduled search
createscript        	Create a new scheduled search w/ script
delete              	Delete a scheduled search
```

From this menu, the administrator can manage *all* scheduled searches on the system, not just his/her own.

## Managing templates/pivots

Templates and pivots each have a sub-menu within the admin menu (`templates` and `pivots`) with an identical set of commands for administrators:

```
admin>  templates
template>  help
list                	List templates
create              	Create a new template
update              	Upload new contents to a template
delete              	Delete a template
print               	Print template contents
updatemeta          	Update template metadata
template>  quit
admin>  pivots
pivot>  help
list                	List pivots
create              	Create a new pivot
update              	Upload new contents to a pivot
delete              	Delete a pivot
print               	Print pivot contents
updatemeta          	Update pivot metadata
pivot>
```

These commands can be used to affect any template or pivot on the system.

## Managing User Files

As with templates, resources, etc., user files also have a sub-menu within the admin menu for admin management. Commands executed within the admin menu can operate on any user file in the whole system.

```
admin>  userfiles
userfile>  help
list                	List available userfiles
add                 	Add a new userfile
update              	Update an existing userfile
del                 	Delete a userfile
get                 	Download a userfile
userfile> 
```
