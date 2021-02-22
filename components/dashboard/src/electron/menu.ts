import { Menu, app, BrowserWindow, shell } from 'electron'
const template: Array<any> = [
    {
        label: 'File',
        submenu: [
            {
                label: 'Open Hosts',
                click() {
                    BrowserWindow.getAllWindows()[0].webContents.send('disconnected')
                }
            }
        ],
    },
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
        label: 'About',
        submenu: [
            {
                label: 'Reload'
            },
            {
                label: 'Documents',
                click: () => { shell.openExternal('https://aid.autoai.org') }
            },
            {
                label: 'Source Code',
                click: () => { shell.openExternal('https://github.com/autoai-org/aid') }
            },
            {
                label: 'Donate',
                click: () => { shell.openExternal('https://github.com/sponsors/xzyaoi/') }
            },
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