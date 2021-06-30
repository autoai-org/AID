const languages = require('./supported-languages')

const language = process.env.LANGUAGE || 'en'

const t = require(`./translations/${language}.json`)

const repoUrl = "https://github.com/autoai-org/aid"

module.exports = {
    title: 'A.I.D',
    tagline: 'Open Source MLOps Platform',
    url: 'https://aid.autoai.org',
    baseUrl: `/`,
    favicon: 'img/favicon.ico',
    organizationName: 'autoai-org',
    projectName: 'aid',
    onBrokenLinks: 'ignore',
    themeConfig: {
        languages,
        language,
        t,
        algolia: {
            apiKey: '1cbcb76e9629ebcbee045b360838e212',
            indexName: 'autoai',
            searchParameters: {},
        },
        navbar: {
            title: 'AID',
            logo: {
                alt: 'AID',
                src: 'img/icon.png',
            },
            items: [
                { to: 'docs/about/intro', label: t.navbar.about, position: 'left' },
                {
                    to: 'docs/getting-started/intro',
                    label: t.navbar.docs,
                    position: 'left',
                },
                {
                    to: 'docs/pages/releases',
                    label: 'Releases',
                    position: 'left',
                },
                {
                    label: t.navbar.community,
                    position: 'left',
                    items: [
                        {
                            label: t.navbar.partners,
                            to: 'partners',
                        },
                        {
                            label: t.navbar.showcase,
                            to: 'showcase',
                        },
                    ],
                },
                {
                    href: repoUrl,
                    'aria-label': 'GitHub',
                    position: 'right',
                    className: 'header-github-link'
                },
            ],
        },
        footer: {
            style: 'dark',
            links: [
                {
                    title: t.navbar.docs,
                    items: [
                        {
                            label: t.navbar.gettingStarted,
                            to: 'docs/getting-started/intro',
                        },
                    ],
                },
                {
                    title: 'Community',
                    items: [
                        {
                            label: 'Email',
                            to: 'mailto:enquiry@autoai.org',
                        },
                        {
                            label: 'Partners',
                            to: '/partners',
                        },
                        {
                            label: 'Discord',
                            to: 'https://discord.gg/3BD3RzK2K2',
                        }
                    ],
                },
                {
                    title: 'Social',
                    items: [
                        {
                            label: 'GitHub',
                            to: 'https://github.com/autoai-org',
                        },
                        {
                            label: 'Twitter',
                            to: 'https://twitter.com/aid_aiops',
                        },
                        {
                            label: 'News',
                            to: '/en/docs/about/news',
                        },
                    ],
                },
            ],
            copyright: `Copyright © ${new Date().getFullYear()} AID Contributors. CC-BY / MIT`,
        },
    },
    presets: [
        [
            '@docusaurus/preset-classic',
            {
                docs: {
                    path: './docs/' + language,
                    sidebarPath: require.resolve('./sidebar.json'),
                    editUrl: 'https://github.com/facebook/docusaurus/edit/master/website/',
                },
                theme: {
                    customCss: require.resolve('./src/css/custom.css'),
                },
            },
        ],
    ],
};