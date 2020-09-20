// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

const overview_menu = [
    { icon: "mdi-view-dashboard", text: "Dashboard", link: "/" },
    { icon: "mdi-inbox", text: "Packages", link: "/overview/packages" },
]

const system_menu = [
    { icon: "mdi-poll-box", text: "Status" },
    { icon: "mdi-math-log", text: "Logs", link: "/system/logs" },
    { icon: "mdi-history", text: "Activity", link: "/system/activities" },
    { icon: "mdi-cog", text: "Preferences", link: "/system/preferences" },
    { icon: "mdi-monitor", text: "Monitor", link: "/system/monitor" },
]

const experiment_menu = [
    { icon: "mdi-database", text: "Dataset", link:"/experiment/dataset" },
    { icon: "mdi-jabber", text: "Experiment", link:"/experiment/new" }
]

const extension_menu = [
    { icon: "mdi-webhook", text: "Webhooks", link:"/extensions/webhooks" },
    { icon: "mdi-grain", text: "Apps", link:"/extensions/apps" }
]

const about_menu = [
    { icon: "mdi-github", text: "GitHub", link: "https://github.com/autoai-org/cvpm" },
    { icon: "mdi-database-search", text: "Registry" },
    { icon: "mdi-help", text: "Docs", link: "https://aid.autoai.org" },
    { icon: "mdi-information-outline", text: "AID Story", link: "https://yaonotes.org/blogs/why-aid.html" },
    { icon: "mdi-gift-outline", text: "Sponsors", link:"https://github.com/sponsors/xzyaoi"},
]



export {
    overview_menu,
    system_menu,
    extension_menu,
    experiment_menu,
    about_menu,
}