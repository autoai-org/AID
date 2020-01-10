interface Log {
    Source: string,
    Title: string,
    ID: number,
    CreatedAt: string
}

interface LogContent {
    level: string,
    time: string,
    msg: string,
}

export { Log, LogContent }