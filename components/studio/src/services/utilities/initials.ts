function getInitials(text: string):string {
    if (text==null) {
        return ""
    }
    return text.match(/(^\S\S?|\b\S)?/g)!.join("").match(/(^\S|\S$)?/g)!.join("").toUpperCase()
}

export {
    getInitials
}