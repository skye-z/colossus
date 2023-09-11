import request from './request'

function post(url, data) {
    return request({
        url: url,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: JSON.stringify(data)
    })
}

function get(url) {
    return request({
        url: url,
        method: 'GET'
    })
}

function buildQuery(form) {
    let query = ''
    for (let key in form) {
        query += key + '=' + form[key] + '&'
    }
    if (query == '') return query
    else return '?' + query.substring(0, query.length - 1)
}

// 应用配置接口
export const config = {
    getAll: () => {
        return get('/config/all')
    },
    update: (key, value) => {
        return post('/config/update', {
            key, value
        })
    },
}

// 主机信息接口
export const host = {
    add: form => {
        return post('/host/add', form)
    },
    edit: (id, form) => {
        return post('/host/' + id, form)
    },
    getList: form => {
        let query = buildQuery(form)
        return get('/host/list' + query)
    },
    getInfo: id => {
        return get('/host/' + id)
    },
    remove: id => {
        return request({
            url: '/host/' + id,
            method: 'DELETE'
        })
    }
}

// 主机文件接口
export const file = {
    getHome: id => {
        return post('/file/home', {
            id
        })
    },
    getList: (id, path, hide) => {
        return post('/file/list', {
            id,
            path,
            hide
        })
    },
    getInfo: (id, path) => {
        return post('/file/info', {
            id,
            path
        })
    },
    download: (id, model, localPath, serverPath, fileName) => {
        return post('/file/down', { id, model, localPath, serverPath, fileName })
    },
    upload: (id, model, localPath, serverPath, fileName) => {
        return post('/file/up', { id, model, localPath, serverPath, fileName })
    },
    move: (id, oldPath, newPath) => {
        return post('/file/move', { id, localPath: oldPath, serverPath: newPath })
    },
    remove: (id, path, fileName) => {
        return post('/file/remove', { id, serverPath: path, fileName })
    },
    createDirectory: (id, serverPath) => {
        return post('/file/directory', { id, serverPath })
    },
}

// 命令接口
export const code = {
    add: form => {
        return post('/code/add', form)
    },
    edit: (id, form) => {
        return post('/code/' + id, form)
    },
    getList: form => {
        let query = buildQuery(form)
        return get('/code/list' + query)
    },
    remove: id => {
        return request({
            url: '/code/' + id,
            method: 'DELETE'
        })
    }
}

// 凭证接口
export const cert = {
    initAdd: level => {
        return get('/cert/add?level=' + level)
    },
    add: form => {
        return post('/cert/add', form)
    },
    edit: (id, form) => {
        return post('/cert/' + id, form)
    },
    getList: form => {
        let query = buildQuery(form)
        return get('/cert/list' + query)
    },
    remove: id => {
        return request({
            url: '/cert/' + id,
            method: 'DELETE'
        })
    }
}