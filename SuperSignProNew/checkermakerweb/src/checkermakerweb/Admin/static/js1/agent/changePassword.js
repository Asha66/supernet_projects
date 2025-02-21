$(function()
  {
    $("#myTab a:last")
      .tab("show")
  }), $(document)
  .ready(function()
  {
    function s(s)
    {
      return s = s.replace(/\s+/, "")
    }
    $("#oldPassword")
      .focus(), $("#reset")
      .button()
      .click(function()
      {
        $("#oldpasstext")
          .nextAll()
          .remove(), $("#oldpasstext")
          .css("color", "black"), $("#newpasstext")
          .nextAll()
          .remove(), $("#newpasstext")
          .css("color", "black"), $("#confirmpasstext")
          .nextAll()
          .remove(), $("#confirmpasstext")
          .css("color", "black")
      }), $("#oldPassword")
      .focusout(function()
      {
        var e = $("#oldPassword")
          .val()
          .length;
        7 > e && ($("#oldpasstext")
          .nextAll()
          .remove(), $("#oldpasstext")
          .css("color", "black"), $("#oldpasstext")
          .after('<p style="color:red"> should be atleast 7 characters long </p')
          .css("color", "red")), e >= 7 && ($("#oldpasstext")
          .nextAll()
          .remove(), $("#oldpasstext")
          .css("color", "black"));
        var o = s($("#oldPassword")
            .val()
            .trim()),
          t = /^[ A-Za-z0-9_@.\/#&+-]*$/;
        return t.test(o) ? void $("#oldpasstext")
          .after('<p style="color:red"> Check password rules below </p')
          .css("color", "red") : ($("#oldpasstext")
            .nextAll()
            .remove(), $("#oldpasstext")
            .css("color", "black"), !1)
      }), $("#newPassword")
      .focusout(function()
      {
        var e = $("#newPassword")
          .val()
          .length;
        7 > e && ($("#newpasstext")
          .nextAll()
          .remove(), $("#newpasstext")
          .css("color", "black"), $("#newpasstext")
          .after('<p style="color:red"> should be atleast 7 characters long </p')
          .css("color", "red")), e >= 7 && ($("#newpasstext")
          .nextAll()
          .remove(), $("#newpasstext")
          .css("color", "black"));
        var o = /^[ A-Za-z0-9_@.\/#&+-]*$/,
          t = s($("#newPassword")
            .val()
            .trim());
        return o.test(t) ? void $("#newpasstext")
          .after('<p style="color:red"> Check password rules below </p')
          .css("color", "red") : ($("#newpasstext")
            .nextAll()
            .remove(), $("#newpasstext")
            .css("color", "black"), !1)
      }), $("#confirmPassword")
      .focusout(function()
      {
        var e = s($("#newPassword")
            .val()
            .trim()),
          o = s($("#confirmPassword")
            .val()
            .trim());
        return e === o ? ($("#confirmpasstext")
            .nextAll()
            .remove(), $("#confirmpasstext")
            .css("color", "black"), !1) : void $("#confirmpasstext")
          .after('<p style="color:red"> Passwords no not match ! </p')
          .css("color", "red")
      })
  });