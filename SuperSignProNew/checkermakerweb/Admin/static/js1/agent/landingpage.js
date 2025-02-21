	function switchTab(id){
					var $ = window.parent.$, jQuery = window.parent.jQuery;
                    $(".active").removeClass("active");
  				   $('ul li:nth-of-type('+id+')').addClass('active');
				}