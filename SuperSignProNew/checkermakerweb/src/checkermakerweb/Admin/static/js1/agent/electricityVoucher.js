    $(function() {     
    function favoriteRecharge(operator, amount) {
      var selectOperator = operator;
      var selectAmount = amount;
      document.getElementById('operator').value = selectOperator;
      document.getElementById('amount').value = selectAmount;
    }
    $(document).ready(function() {
      $("#operator").click(function() {
        $("#errmsgSelect").hide();
      });
      $("#amount").click(function() {
        $("#errmsg").hide();
      });
      $("#proceedsubmit").click(function() {
        if ($("#operator").val() == 'select') {
          $("#errmsgSelect").html("<sup>*<\/sup> Select the Operator").show();
        }
        if ($("#amount").val() == '') {
          $("#errmsg").html("<sup>*<\/sup> Enter the Amount").show();
        } else {
          $("#transactionPasswordDiv").show();
          $("#transactionPassword").focus();
        }
      });
      //called when key is pressed in textbox
      $("#amount").keypress(function(e) {
        //if the letter is not digit then display error and don't type anything
        if (e.which != 8 && e.which != 0 && (e.which < 48 || e.which > 57)) {
          //display error message
          $("#errmsg").html("<sup>*<\/sup> Enter Numbers Only").show();
          return false;
        } else {
          $("#errmsg").html("<sup>*<\/sup> Enter Numbers Only").hide();
        }
      });
      $("#transactionProceed").click(function() {
        if ($("#transactionPassword").val() == '') {
          $("#errmsgProceed").html("<sup>*<\/sup> Enter your Transaction Password").show();
          $("#transactionPassword").focus();
        } else {
          $("#transactionPasswordDiv").show();
        }
      });
    });
    function Submit() {
      document.getElementById('ElectricityVoucher').submit();
    }