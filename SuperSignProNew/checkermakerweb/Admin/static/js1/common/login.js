$(document).ready(function($) {
  $("a.btn-navbar").hide();
  $("#username").focus();
  setTimeout(function() {
    $('body').addClass('loaded');
    $('h1').css('color', '#222222');
  }, 3000);
});
$('#Register').click(function() {
  window.location.href = 'signup.html';
  return false;
});
$('#Submit').click(function() {
  if ($.trim($('#username').val()) == "") {
    $("#usernameMessage").html("Please enter your user identity").css('color', 'red');
    return false;
  }
});
$('#Submit').click(function() {
  if ($.trim($('#password').val()) == "") {
    $("#passwordMessage").html("Please enter your password").css('color', 'red');
    return false;
  }
});