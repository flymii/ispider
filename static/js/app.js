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
            if (inputValue === false) returnfalse; 
            
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

function user_edit(id){
    swal("编辑用户：" + id)
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
    swal("启用用户：" + id)
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