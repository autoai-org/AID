// Copyright (c) 2020 Xiaozhe Yao & AICAMP.CO.,LTD
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import * as Converter from 'ansi-to-html'
let converter = new Converter();

function ansi2html(ansiLogs: string){
    return converter.toHtml(ansiLogs)
}

export {
    ansi2html
}