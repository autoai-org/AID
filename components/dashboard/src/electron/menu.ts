import {Menu, app} from 'electron'

const template:Array<any> = [
    {
        label: 'Edit App',
        submenu: [
            {
                label: 'Undo'
            },
            {
                label: 'Redo'
            }
        ]
    },
    {
        label: 'View App',
        submenu: [
            {
                label: 'Reload'
            },
            {
                label: 'Toggle Full Screen'
            }
        ]
    }
];

if (process.platform === 'darwin') {
    template.unshift({
        label: app.getName(),
        submenu: [
            {
                label: 'Quit',
                accelerator: 'CmdOrCtrl+Q',
                click() {
                    app.quit();
                }
            }
        ]
    });
}

const appMenu = Menu.buildFromTemplate(template)

export default appMenu