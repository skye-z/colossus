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

export const host = {
    add: (form) => {
        return post('/host/add', form)
    },
    edit: (id, form) => {
        return post('/host/' + id, form)
    },
    getList: (form) => {
        let query = buildQuery(form)
        return get('/host/list' + query)
    }
}

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