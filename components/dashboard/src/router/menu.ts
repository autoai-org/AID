// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

const overview_menu = [
    { icon: "mdi-view-dashboard", text: "Dashboard", link: "/"},
    { icon: "mdi-inbox", text: "Packages", link: "/overview/packages" },
]

const system_menu = [
    { icon: "mdi-poll-box", text: "Status" },
    { icon: "mdi-math-log", text: "Logs", link: "/system/logs"},
    { icon: "mdi-puzzle", text: "Extensions" },
    { icon: "mdi-settings-outline", text: "Preferences",link: "/system/preferences" },
]

const about_menu = [
    { icon: "mdi-github-circle", text: "GitHub", link: "https://github.com/autoai-org/cvpm"},
    { icon: "mdi-database-search", text: "Registry" },
    { icon: "mdi-help", text: "Docs", link: "https://aid.autoai.org"},
    { icon: "mdi-information-outline", text: "AID Story" },
]

export {
    overview_menu,
    system_menu,
    about_menu,
}