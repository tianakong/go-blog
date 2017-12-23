layui.define(['form', 'element'], function (exports) {
    var form = layui.form,
        $ = layui.jquery,
        layer = layui.layer;

    form.on('submit(reg)', function (data) {
        $.ajax({
            type: 'post',
            data: data.field,
            url: $(data.form).attr('action'),
            dataType: 'json',
            success: function (res) {
                if(res.status) {
                    window.location.href="/register"
                } else {
                    layer.msg(res.info);
                }
            },
            error: function (res) {

            }
        })
        return false
    })
    form.on('submit(login)', function (data) {
        $.ajax({
            type: 'post',
            data: data.field,
            url: $(data.form).attr('action'),
            dataType: 'json',
            success: function (res) {
                if(res.status) {
                    window.location.href="/"
                } else {
                    layer.msg(res.info);
                }
            },
            error: function (res) {

            }
        })
        return false
    })

    exports('index');
})