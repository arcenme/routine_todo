$(document).ready(function () {
    // tooltip
    $('[data-toggle="tooltip"]').tooltip();

    $("#deadline").datepicker({
        format: "dd MM yyyy",
        locale: 'id',
        autoclose: true,
        todayHighlight: true,
        startDate: moment(new Date()).format('DD MMMM yyyy'),
        orientation: "bottom right"
    });

    // render task
    function renderTask(tasks) {
        const html = tasks.map(task => {
            return `<div class="row px-3 mb-2 align-items-center todo-item rounded">
                        <div class="col-auto m-1 p-0 d-flex align-items-center">
                            <h2 class="m-0 p-0">
                                <i class="fa ${task.is_done == '1' ? 'fa-check-square-o' : 'fa-square-o'} text-primary btn m-0 p-0 btn-status" data-id="${task.id}" data-isdone="${task.is_done}" data-toggle="tooltip" data-placement="bottom" title="${task.is_done ? 'Mark as todo' : 'Masrk as done'}"></i>
                            </h2>
                        </div>
                        <div class="col px-1 m-1 d-flex align-items-center">
                            <input type="text" class="form-control form-control-plaintext form-control-lg border-0 edit-todo-input rounded px-3" value="${task.task_name}" />
                        </div>
                        <div class="col-auto m-1 p-0 px-3">
                            <div class="row">
                                <div class="col-auto d-flex align-items-center rounded bg-white border border-warning">
                                    <i class="fa fa-hourglass-2 my-2 px-2 text-warning btn" data-toggle="tooltip" data-placement="bottom" title="" data-original-title="Due on date"></i>
                                    <h6 class="text my-2 pr-2">${moment(task.deadline).format('DD MMM YYYY')}</h6>
                                </div>
                            </div>
                        </div>
                        <div class="col-auto my-1 ml-4 mr-1  p-0 todo-actions" style="width:100px">
                            <div class="row d-flex align-items-center justify-content-end">
                                <h5 class="m-0 p-0 px-2">
                                    <i class="fa fa-pencil text-info btn m-0 p-0 btn-edit" data-id="${task.id}" data-toggle="tooltip" data-placement="bottom" title="Edit todo"></i>
                                </h5>
                                <h5 class="m-0 p-0 px-2">
                                    <i class="fa fa-trash-o text-danger btn m-0 p-0 btn-delete" data-id="${task.id}" data-toggle="tooltip" data-placement="bottom" title="Delete todo"></i>
                                </h5>
                            </div>
                            <div class="row todo-created-info" style="float:right">
                                <div class="col-auto d-flex align-items-center pr-2">
                                    <label class="date-label my-2 text-black-50">${task.assignee}</label>
                                    <i class="fa fa-info-circle ml-2 mr-0 px-0 text-black-50 btn" data-toggle="tooltip" data-placement="bottom" title="Assignee" data-original-title="Assignee"></i>
                                </div>
                            </div>
                        </div>
                    </div>`
        })

        $('.render-task').empty().append(html)
    }

    // get all task 
    function getTask() {
        $.ajax({
            contex: this,
            url: "/api/routine",
            type: 'GET',
            datatype: 'json',
            success: function (res) {
                renderTask(res.data)
            }, error: function (err) {
                iziToast.error({
                    title: 'Error',
                    message: 'Something went wrong!'
                });
            }
        })
    }

    getTask()

    // modal add
    $('#btn-add').click(function () {
        $('#form-add-edit').trigger('reset')
        $('#form-add-edit input').removeClass('is-invalid')
        $('.modal-add-edit-title').text('Add Task')
        $('#submit-add-edit').val('add')
        $('#modalAddAndEdit').modal('show')
    })

    // submit add/edit task
    $('#submit-add-edit').click(function () {
        const data = {
            "task_name": $('#task_name').val(),
            "assignee": $('#assignee').val(),
            "deadline": $('#deadline').val() ? moment($('#deadline').val()).format('YYYY-MM-DD') : ''
        }

        if ($(this).val() == 'edit')
            data['id'] = $('#id').val()

        $.ajax({
            contex: this,
            type: 'POST',
            url: '/api/routine',
            contentType: "application/json; charset=utf-8",
            dataType: 'json',
            data: JSON.stringify(data),
            beforeSend: function () {
                $(this).text('...saving')
                $(this).prop('disabled', true)
            }, success: function (res) {
                $(this).text('Save changes')
                $(this).prop('disabled', false)

                $('#modalAddAndEdit').modal('hide')
                getTask()
            }, error: function (err) {
                $(this).text('Save changes')
                $(this).prop('disabled', false)

                iziToast.error({
                    title: 'Error',
                    message: 'Internal server error'
                });
            }
        })
    })
})