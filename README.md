# alfred-slack-workflow

Alfred workflow to open slack channel.

## Install and setup

1. Download latest version from [here](https://github.com/monda00/alfred-slack-workflow/releases) and import workflow

2. Get slack token
   1. Create a slack app [here](https://api.slack.com/apps/new)

       ![creat_app](https://user-images.githubusercontent.com/25761599/114042525-793b3c00-98c0-11eb-9414-c3a7a0d1f1a8.png)

   2. Add `channnels:read`, `groups:read`, `im:read`, `mpim:read` and `team:read` to Bot Token Scopes

       ![oauth](https://user-images.githubusercontent.com/25761599/114047810-f072cf00-98c4-11eb-95c9-c93a057c11de.png)

   3. Install app to workspace

       ![install_app](https://user-images.githubusercontent.com/25761599/114044687-60cc2100-98c2-11eb-9c53-2a5cfeca003f.png)

   4. Copy bot OAuth token

       ![copy_token](https://user-images.githubusercontent.com/25761599/114048616-b2c27600-98c5-11eb-93f1-bebe2815a1f1.png)

3. Set token to alfred workflow environment variables

    ![set_token](https://user-images.githubusercontent.com/25761599/114048868-e56c6e80-98c5-11eb-9789-3a75b4d27bbd.png)
    ![input_token](https://user-images.githubusercontent.com/25761599/114049347-57dd4e80-98c6-11eb-8100-382f68874670.png)

## How to use

### Update channels.

Update channels list.

`update channels`

<img width="629" alt="update_channel" src="https://user-images.githubusercontent.com/25761599/114295936-71a7ad00-9ae3-11eb-8e40-fc67d729770e.png">

### Open channels.

Open channel in the Slack app.

`sl {channel}`

<img width="630" alt="open_channel" src="https://user-images.githubusercontent.com/25761599/114295940-740a0700-9ae3-11eb-9bd5-18ca053b9ab5.png">
