module.exports = {
    dest: 'docs/dist',
    locales: {
      '/': {
        lang: 'en(US)',
        title: 'CVPM',
        description: 'Package Manager for Computer Vision'
      },
      '/zh-CN/': {
        lang: 'zh-CN',
        title: 'CVPM',
        description: '用于计算机视觉的包管理'
      }
    },
    head: [
      ['link', { rel: 'icon', href: `/logo.png` }],
      ['link', { rel: 'manifest', href: '/manifest.json' }],
      ['meta', { name: 'theme-color', content: '#3eaf7c' }],
      ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
      ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }],
      ['link', { rel: 'apple-touch-icon', href: `/icons/apple-touch-icon-152x152.png` }],
      ['link', { rel: 'mask-icon', href: '/icons/safari-pinned-tab.svg', color: '#3eaf7c' }],
      ['meta', { name: 'msapplication-TileImage', content: '/icons/msapplication-icon-144x144.png' }],
      ['meta', { name: 'msapplication-TileColor', content: '#000000' }]
    ],
    serviceWorker: true,
    themeConfig: {
      repo: 'unarxiv/cvpm',
      editLinks: false,
      lastUpdated: 'Last Updated',
      docsDir: 'docs',
      locales: {
        '/': {
          label: 'English',
          selectText: 'Languages',
          editLinkText: 'Edit this page on GitHub',
          nav: [
            {
              text: 'Guide',
              link: '/en-US/guide/',
            },
            {
              text: 'Blog',
              link: 'https://blog.cvtron.xyz/'
            },
            {
        
              text: 'Status',
              link: 'https://status.cvtron.xyz/'
            },
            {
              text: 'Forum',
              link: 'https://forum.cvtron.xyz/'
            },
            {
              text: 'Community',
              items: [
                { text: 'CVTron', link: 'https://github.com/unarxiv/CVTron' }
              ]
            }
          ],
          sidebar: {
            '/en-US/guide/': genSidebarConfig('Guide')
          }
        },
        '/zh-CN/': {
          label: '简体中文',
          selectText: '选择语言',
          lastUpdated: '上次更新',
          editLinkText: '在 GitHub 上编辑此页',
          nav: [
            {
              text: '指南',
              link: '/zh-CN/guide/',
            },
            {
              text: '博客',
              link: 'https://blog.cvtron.xyz/'
            },
            {
        
              text: '状态',
              link: 'https://status.cvtron.xyz/'
            },
            {
        
              text: '论坛',
              link: 'https://forum.cvtron.xyz/'
            },
            {
              text: '社群',
              items: [
                { text: 'CVTron', link: 'https://github.com/unarxiv/CVTron' },
              ]
            }
          ],
          sidebar: {
            '/zh-CN/guide/': genSidebarConfig('指南')
          }
        }
      }
    }
  }
  
  function genSidebarConfig (title) {
    return [
      {
        title,
        collapsable: true,
        children: [
          '',
          'getting-started',
        ]
      }
    ]
  }