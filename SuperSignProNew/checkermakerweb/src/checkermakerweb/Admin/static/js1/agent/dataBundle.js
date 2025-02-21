
function favoriteRecharge(amount) {
    var selectAmount = amount;
    document.getElementById('netone_amount').value = selectAmount;
    document.getElementById('netone_mob').focus();
    $("#errmsgamount").html("<sup>*<\/sup> Enter the Amount").hide();
    
    document.getElementById('telecel_amount').value = selectAmount;
    document.getElementById('telecel_mob').focus();
    $("#errmsgamount2").html("<sup>*<\/sup> Enter the Amount").hide();

    document.getElementById('econet_amount').value = selectAmount;
    document.getElementById('econet_mob').focus();
    $("#errmsgamount3").html("<sup>*<\/sup> Enter the Amount").hide();
}

function Submit1() {    
    var denomination = document.getElementById('netone_amount').value;
    var num = document.getElementById('netone_mob').value; 
    var pass = document.getElementById('netone_transactionPassword').value; 
    document.getElementById("NETONE").submit();
}

function Submit2() {    
    var denomination = document.getElementById('telecel_amount').value;
    var num = document.getElementById('telecel_mob').value; 
    var pass = document.getElementById('telecel_transactionPassword').value; 
    document.getElementById("TELECEL").submit();
}

function Submit3() {    
    var denomination = document.getElementById('econet_amount').value;
    var num = document.getElementById('econet_mob').value; 
    var pass = document.getElementById('econet_transactionPassword').value;
    document.getElementById("ECONET").submit();
}


function validate(evt) {
    var theEvent = evt || window.event;
    var key = theEvent.keyCode || theEvent.which;
    key = String.fromCharCode(key);
    var regex = /[0-9]|\./;
    if (!regex.test(key)) {
        theEvent.returnValue = false;
        if (theEvent.preventDefault) theEvent.preventDefault();
    }
}

$(document).ready(function() 
{
    $("#transactionProceed").click(function() {
        if ($("#netone_transactionPassword").val() == '') {
            $("#errmsgProceed").html("<sup>*<\/sup> Enter your Transaction Password").show();
            $("#netone_transactionPassword").focus();
        } else {
            $("#transactionPasswordDiv").show();
        }
    });
    $("#transactionProceed2").click(function() {
        if ($("#telecel_transactionPassword").val() == '') {
            $("#errmsgProceed2").html("<sup>*<\/sup> Enter your Transaction Password").show();
            $("#telecel_transactionPassword").focus();
        } else {
            $("#transactionPasswordDiv2").show();
        }
    });
    $("#transactionProceed3").click(function() {
        if ($("#transactionPassword3").val() == '') {
            $("#errmsgProceed3").html("<sup>*<\/sup> Enter your Transaction Password").show();
            $("#econet_transactionPassword").focus();
        } else {
            $("#transactionPasswordDiv3").show();
        }
    });

    $('a[data-toggle="tab"]').on('shown.bs.tab', function(e) {
        var currentTab = $(e.target).text(); // get current tab  
        $(".li").removeClass('active');
        $(e.target).addClass('active');
        $(".formData").val(currentTab);
         $("#netone_amount").val("select");
     $("#telecel_amount").val("select");
      $("econet_amount").val("select");

    });
    $("#netone_amount").click(function() {

        $("#errmsgamount").hide();
    });
    $("#procedsubmit").click(function() 
    {
        var e = $("#netone_amount").val();
        if (e == "select") 
        {
            $("#errmsgamount").html("<sup>*<\/sup> Enter the Amount").show();
            $("#netone_amount").focus();
        } else {
            $("#errmsgamount").html("<sup>*<\/sup> Enter the Amount").hide();
        }
        var number = $("netone_mob").val();
        if (number = " " || number.length < 9) {
            $("#errmsg").html("<sup>*<\/sup> Please Enter a Valid Number ").show();
        }
        var number = $("netone_mob").val();
        if (number.length > 8) 
        {
            $("#errmsg").html("<sup>*<\/sup> Please Enter a Valid Number ").hide();
            $("#transactionPasswordDiv").show();
            $("netone_transactionPassword").focus();
        }
    });

    $("#telecel_amount").click(function() {

        $("#errmsgamount2").hide();
    });

    $("#procedsubmit2").click(function() {
        var e = $("#telecel_amount").val();
        if (e == "select") 
        {
            $("#errmsgamount2").html("<sup>*<\/sup> Enter the Amount").show();
            $("#telecel_amount").focus();
        } else {
            $("#errmsgamount2").html("<sup>*<\/sup> Enter the Amount").hide();
        }
        var number = $("#telecel_mob").val();
        if (number = " " || number.length < 9) 
        {
            $("#errmsg2").html("<sup>*<\/sup> Please Enter a Valid Number ").show();
        }
        var number = $("#telecel_mob").val();
        if (number.length > 8) 
        {
            $("#errmsg2").html("<sup>*<\/sup> Please Enter a Valid Number ").hide();
            $("#transactionPasswordDiv2").show();
            $("#telecel_transactionPassword").focus();
        }
    });

    $("#econet_amount").click(function() {

        $("#errmsgamount3").hide();
    });

    $("#procedsubmit3").click(function() {
        var e = $("#econet_amount").val();
        if (e == "select") 
        {
            $("#errmsgamount3").html("<sup>*<\/sup> Enter the Amount").show();
            $("#econet_amount").focus();
        } else {
            $("#errmsgamount3").html("<sup>*<\/sup> Enter the Amount").hide();
        }
        var number = $("#econet_mob").val();
        if (number = " " || number.length < 9) 
        {
            $("#errmsg3").html("<sup>*<\/sup> Please Enter a Valid Number ").show();
        }
        var number = $("#econet_mob").val();
        if (number.length > 8) 
        {
            $("#errmsg3").html("<sup>*<\/sup> Please Enter a Valid Number ").hide();
            $("#transactionPasswordDiv3").show();
            $("#econet_transactionPassword").focus();
        }
    });
	var receivedOperator = {{.receivedOperator}};
	if (receivedOperator == 'NETONE')
		{
			$("#tab_1").addClass('active');
			$("#tab_2").removeClass('active');
			$("#tab_3").removeClass('active');
			$("#N1").addClass('active').siblings().removeClass('active');
		}
		else if (receivedOperator == 'TELECEL')
		{
			$("#tab_1").removeClass('active');
			$("#tab_2").addClass('active');
			$("#tab_3").removeClass('active');
			$("#T1").addClass('active').siblings().removeClass('active');
		}
		else if (receivedOperator == 'ECONET')
		{
			$("#tab_1").removeClass('active');
			$("#tab_2").removeClass('active');
			$("#tab_3").addClass('active');
			    $("#E1").addClass('active').siblings().removeClass('active');
		}
});