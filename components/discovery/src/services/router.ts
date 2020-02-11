import Router from "koa-router";
import PingController from "../controllers/ping"

const router = new Router()

router.get("/ping", PingController.pong)

export default router