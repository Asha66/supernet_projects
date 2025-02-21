$(document).ready(function($) {
  $("a.btn-navbar").hide();
  setTimeout(function() {
    $('body').addClass('loaded');
  }, 1000);
  //change navbar menu in mobile devices
  var isMobile = window.matchMedia("only screen and (max-width: 480px)");
  if (isMobile.matches) {
    //Conditional script here
    $(".navbutton").hide();
$(".menubutton").show();
    $("#iFrameLandingPage").height(" 1000px");
  }
  //Load iframes - navbar navigation
  //change navbar menu in mobile devices
  var isMobile = window.matchMedia("only screen and (min-width: 768px)");
  if (isMobile.matches) {
    //Conditional script here
    $(".menubutton").hide();
  }
  var isMobile = window.matchMedia("only screen and (max-width: 767px)");
  if (isMobile.matches) {
    //Conditional script here
    $("#iFrameLandingPage").height(" 1000px");
  }

});


$('li').click(function() {
  $('li', $(this).parent()).removeClass('active');
  $(this).addClass('active');
});
