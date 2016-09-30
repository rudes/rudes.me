autocmd BufWritePost * silent !docker cp static/css/base.css rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/css/base.css
autocmd BufWritePost * silent !docker cp static/js/base.js rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/js/base.js
autocmd BufWritePost * silent !docker cp static/templates/index.tmpl rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/templates/index.tmpl
autocmd BufWritePost * silent !docker cp static/templates/header.tmpl rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/templates/header.tmpl
autocmd BufWritePost * silent !docker cp static/templates/base.tmpl rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/templates/base.tmpl
autocmd BufWritePost * silent !docker cp static/templates/post.tmpl rudesme_web_1:/go/src/github.com/rudes/rudes.me/static/templates/post.tmpl
