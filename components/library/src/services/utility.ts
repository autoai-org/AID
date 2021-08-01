// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

export function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}

export function parseQuery(keyword: string): string {
    const urlParams = new URLSearchParams(window.location.search);
    return urlParams.get(keyword) || ""
}

export function getInitials(text: string):string {
    if (text==null) {
        return ""
    }
    return text.match(/(^\S\S?|\b\S)?/g)!.join("").match(/(^\S|\S$)?/g)!.join("").toUpperCase()
}