import * as Koa from "koa"
import router from "./services/router"
import * as json from "koa-json"
import * as logger from "koa-logger"

const app = new Koa();

app.use(json())
app.use(logger())

app.use(router.routes()).use(router.allowedMethods());

app.listen(3000, ()=> {
    console.log("AID Discovery service is running on port 3000")
})