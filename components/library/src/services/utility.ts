// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

export function classNames(...classes: any) {
    return classes.filter(Boolean).join(' ')
}
