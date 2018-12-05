function getMenus () {
  return [
        { 'header': 'Admin' },
        { 'href': '/home', 'title': 'Home', 'icon': 'home' },
        { 'href': '/package', 'title': 'Packages', 'icon': 'grain' },
        { 'href': '/pretrained', 'title': 'Pre-trained', 'icon': 'storage' },
        { 'header': 'System' },
        { 'href': '/log', 'title': 'Log', 'icon': 'info' },
        { 'href': '/settings', 'title': 'Settings', 'icon': 'settings' },
        { 'href': '/login', 'icon': 'lock', 'title': 'Logout' }
  ]
}

export {
    getMenus
}
