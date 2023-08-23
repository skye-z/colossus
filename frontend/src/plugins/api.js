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

export const host = {
    add: (form) => {
        return post('/host/add', form)
    }
}