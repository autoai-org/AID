function drawBBox(ctx: CanvasRenderingContext2D, item:any, im_w:number,im_h:number,cv_w:number,cv_h:number) {
    if (ctx!= null) {
        ctx.beginPath()
        for(let i=0;i<item.length;i++) {
            ctx.lineWidth=0.5
            let x_0 = item[i].loc[0]
            let x_1 = item[i].loc[1]
            let y_0 = item[i].loc[2]
            let y_1 = item[i].loc[3]
            x_0 = x_0 * cv_w/im_w
            y_0 = y_0 * cv_h/im_h
            x_1 = x_1 * cv_w/im_w
            y_1 = y_1 * cv_h/im_h
            ctx.rect(x_0,y_0, x_1-x_0,y_1-y_0);
            ctx.stroke()
        }    
    }
}

export {
    drawBBox
}