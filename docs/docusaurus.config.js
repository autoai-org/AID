module.exports = {
  title: 'A.I.D',
  tagline: 'Open Source MLOps Platform',
  url: 'https://aid.autoai.org',
  baseUrl: '/',
  favicon: 'img/favicon.ico',
  organizationName: 'autoai-org', // Usually your GitHub org/user name.
  projectName: 'aid', // Usually your repo name.
  themeConfig: {
    navbar: {
      title: 'AID',
      logo: {
        alt: 'AID',
        src: 'img/icon.png',
      },
      items: [
        {
          to: 'docs/overview',
          activeBasePath: 'docs',
          label: 'Docs',
          position: 'left',
        },
        {to: '/blog', label: 'Blog', position: 'left'},
        {to: '/docs/examples/showroom', label: 'Showroom', position: 'left'},
        {to: '/changelog', label: 'Changelog', position: 'left'},
        {to: '/docs/releases', label: 'Releases', position: 'left'},
        {
          href: 'https://hub.autoai.org',
          label: 'AIDHub',
          position: 'right',
        },
        {
          href: 'https://github.com/autoai-org/aid',
          label: 'GitHub',
          position: 'right',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [
        {
          title: 'Docs',
          items: [
            {
              label: 'Quick Start',
              to: 'docs/quickstart',
            },
            {
              label: 'Full Docs',
              to: 'docs/overview',
            },
            {
              label: 'Videos',
              to: 'https://www.youtube.com/playlist?list=PL2BhlRIBzYXIkKBOTjU4nzYvxpXqemb2n',
            },
          ],
        },
        {
          title: 'Community',
          items: [
            {
              label: 'Google Group',
              to: 'https://groups.google.com/forum/#!newtopic/autoai',
            },
            {
              label: 'Slack',
              to: 'https://join.slack.com/t/autoaiworkspace/shared_invite/zt-d0ibh5gj-_AIRcj1CedTBiXeqJsMwwQ',
            },
            {
              label: 'Telegram Group',
              to: 'https://t.me/autoai_org',
            },
          ],
        },
        {
          title: 'Social',
          items: [
            {
              label: 'Blog',
              to: 'blog',
            },
            {
              label: 'GitHub',
              to: 'https://github.com/autoai-org/aid',
            },
            {
              label: 'Twitter',
              to: 'https://twitter.com/aid_aiops',
            },
            {
              label: 'News',
              to: '/docs/news',
            },
          ],
        },
      ],
    },
  },
  presets: [
    [
      '@docusaurus/preset-classic',
      {
        docs: {
          sidebarPath: require.resolve('./sidebars.js'),
          editUrl:
            'https://github.com/facebook/docusaurus/edit/master/website/',
        },
        theme: {
          customCss: require.resolve('./src/css/custom.css'),
        },
      },
    ],
  ],
};
