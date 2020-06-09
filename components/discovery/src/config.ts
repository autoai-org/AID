const config = {
    MONGO_URI: process.env.MONGO_URI,
    PORT: process.env.PORT || 3000,
    GITHUB_TOKEN: process.env.GITHUB_TOKEN,
    JWT_SECRET: process.env.JWT_SECRET || "autoai-aid"
}

export default config