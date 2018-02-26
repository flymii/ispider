/*------users.html start------*/
function user_add(){
    swal({ 
            title: "Add New User", 
            text: "Please input user name：",
            type: "input", 
            showCancelButton: true, 
            closeOnConfirm: false, 
            animation: "slide-from-top", 
            inputPlaceholder: "user name" 
        },
        function(inputValue){ 
            if (inputValue === false) return false; 
            
            if (inputValue === "") { 
                swal.showInputError("Please input user name！");
                return false 
            } 

            $.ajax({
                url:"/admin/service/user_add",
                dataType:'json',
                type:"POST",
                cache:false,
                data:{username:inputValue},
                success:function(data){
                    if(data == null){
                        swal("添加失败", "服务器错误", "error");
                        return;
                    }
                    if (data.status != 0){
                        swal("添加失败", data.msg, "error");
                        return;
                    }
                    swal({title:"添加成功!",text: "添加用户:【" + inputValue + "】成功!", type:"success" }, function(){
                        location.reload();
                    });
                }
            })
        }
    );
}

function user_edit(){
    var level = $('select[name=level]').val();
    var id = $('input[name=id]').val();
    $.ajax({
        url:'/admin/service/user_edit',
        dataType:'json',
        type:'POST',
        cache:false,
        data:{id:id, level:level},
        success:function(data){
            if(data == null){
                swal("修改失败", "服务器错误", "error");
                return;
            }
            if (data.status != 0){
                swal("修改失败", data.msg, "error");
                return;
            }
            swal({title:"修改成功!",text: data.msg, type:"success" }, function(){
                window.location.href = document.referrer
            });
        }
    })
}

function user_delete(id){
    swal({ 
        title: "确定删除吗？", 
        text: "你将无法恢复该用户！", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定删除！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_delete',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("删除失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("删除失败", data.msg, "error");
                    return;
                }
                swal({title:"删除成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function user_enable(id){
    swal({ 
        title: "确定启用该用户吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定启用！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_enable',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("启用失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("启用失败", data.msg, "error");
                    return;
                }
                swal({title:"启用成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}

function user_disable(id){
    swal({ 
        title: "确定禁用该用户吗？", 
        text: "", 
        type: "warning",
        showCancelButton: true, 
        confirmButtonColor: "#DD6B55",
        confirmButtonText: "确定禁用！", 
        closeOnConfirm: false
    },
    function(){
        $.ajax({
            url:'/admin/service/user_disable',
            dataType:'json',
            type:'POST',
            cache:false,
            data:{id:id},
            success:function(data){
                if(data == null){
                    swal("禁用失败", "服务器错误", "error");
                    return;
                }
                if (data.status != 0){
                    swal("禁用失败", data.msg, "error");
                    return;
                }
                swal({title:"禁用成功!",text: data.msg, type:"success" }, function(){
                    location.reload();
                });
            }
        })
    });
}
/*------users.html end------*/