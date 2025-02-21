
        $.widget.bridge('uibutton', $.ui.button);
        $(function() {
            $( "#transactionFromDate" ).datepicker();
            $( "#transactionToDate" ).datepicker();
            $("#searchResult").dataTable();
        });
