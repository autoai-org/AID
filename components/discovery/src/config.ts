const config = {
    MONGO_URI: process.env.MONGO_URI,
    PORT: process.env.PORT || 3000,
    GITHUB_TOKEN: process.env.GITHUB_TOKEN,
    JWT_SECRET: process.env.JWT_SECRET || "autoai-aid",
    GOOGLE_CLIENT_ID: process.env.GOOGLE_CLIENT_ID || "",
    GOOGLE_SECRET:process.env.GOOGLE_SECRET || "",
}

export default config