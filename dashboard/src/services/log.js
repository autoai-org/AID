import sanitizeHtml from 'sanitize-html'

function parseLogMessage (msg) {
  let parsed = msg.replace(/warn/gi, "<span style='color:yellow'>warn</span>")
  parsed = parsed.replace(/error/gi, '<span style="color:red"><B>error</B></span>')
  parsed = parsed.replace(/info/gi, '<span style="color:green">info</span>')
  parsed = parsed.replace(/severe/gi, '<span style="color:red">severe</span>')
  parsed = parsed.replace(/failed/gi, '<span style="color:red">failed</span>')
  parsed = parsed.replace(/exception/gi, '<span style="color:red">exception</span>')
  return sanitizeHtml(parsed, {
    allowedTags: ['b', 'i', 'em', 'strong', 'span'],
    allowedAttributes: {
      a: ['href'],
      span: ['style']
    },
    allowedIframeHostnames: ['autoai.org']
  })
}

export {
  parseLogMessage
}
