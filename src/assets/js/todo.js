// 初期表示の時の処理
$(function() {
    // 初期表示で警告表示領域を非表示にする
    clearDivAlert();

    // 期限のカレンダー表示を日本語にする
    $('#inputLimit').datepicker({
        language: 'ja'
    });

    // 追加ボタンを押した時の処理
    $("#btnAdd").on("click", function () {
        clearDivAlert();

        // 入力チェック
        if ($('#inputTask').val().trim().length == 0) {
            $('#divAlert').css('display', 'block');
            $('#inputAlert').text('タスク内容を入力してください');
            return;
        }

        if ($('#inputLimit').val().trim().length == 0) {
            $('#divAlert').css('display', 'block');
            $('#inputAlert').text('期限を入力してください');
            return;
        }

        // タスクの新規作成.
        $.ajax({
            type: 'POST',
            url: '/api/todo',
            data: {
                name: $('#inputTask').val(),
                deadline: $('#inputLimit').val()
            }
        }).done(function () {
            // reload.
            location.reload();
        }).fail(function () {
            $('#todoModal').modal('hide');
            alert('登録失敗.');
        });
    });

    $(document).on("click", ".btnDel", function () {
        const id = $(this).data("todo-id");

        // 完了ステータスへ更新.
        $.ajax({
            type: "POST",
            url: "/api/todo/" + id + "/complete"
        }).done(function() {
            location.reload();
        }).fail(function() {
            alert('更新失敗.');
        });
    });

    // 警告が出ていた場合に、中身を削除して非表示にする
    function clearDivAlert() {
        $('#divAlert').css('display', 'none');
        $('#inputAlert').text('');
    }
});