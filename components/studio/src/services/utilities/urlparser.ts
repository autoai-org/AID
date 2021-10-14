function parseURL(url: string) : Object {
    const urlParams = new URLSearchParams(url)
    let results:{ [key: string]: string} = {}
    for (const entry of urlParams.entries()) {
        results[entry[0]] = entry[1]
    }
    return results
}

export {
    parseURL
}