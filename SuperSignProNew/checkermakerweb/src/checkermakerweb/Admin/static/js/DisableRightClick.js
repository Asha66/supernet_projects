document.oncontextmenu = function(){return false}
if(document.layers) {window.captureEvents(Event.MOUSEDOWN);window.onmousedown = function(e){if(e.target==document)return false;}}
else {}
var ctrlNameBKSPACE = "";
var ctrltypeBKSPACE = "";
window.history.forward(1);
document.onmousewheel = function()
{
	var key = (window.event) ? event.keyCode : (event.which) ? event.which : event.keyCode;
	if(event.shiftKey)
	{
		event.cancelBubble = true;
		event.returnValue = false;
		key = false; 
		return false;
	}
}
document.ondrag = function()
{
	var key = (window.event) ? event.keyCode : (evt.which) ? evt.which : evt.keyCode;event.cancelBubble = true;
	event.returnValue = false;
	key = false; 
	return false;
}
document.ondragenter = function()
{
	event.cancelBubble = true;
	event.returnValue = false;
	key = false;
	return false;
}
document.ondragover = function()
{
	event.cancelBubble = true;
	event.returnValue = false;
	key = false;
	return false;
}
document.ondrop = function()
{
	event.cancelBubble = true;
	event.returnValue = false;
	key = false; 
	return false;
}
document.ondragstart = function()
{
	event.cancelBubble = true;
	event.returnValue = false;
	key = false;
	return false;
}
document.onkeydown = function()
{
	var key = (window.event) ? event.keyCode : (event.which) ? event.which : event.keyCode;
	if ((key == 78 || key == 88) && (event.ctrlKey))
	{
		event.cancelBubble = true;
		event.returnValue = false;
		key = false;return false;
	
	}
/*ctrlNameBKSPACE = event.srcElement.form;
if ( key == 8) {	
ctrltypeBKSPACE = document.activeElement.getAttribute("type");
if(ctrlNameBKSPACE=="[object]" && (ctrltypeBKSPACE=="text" || ctrltypeBKSPACE=="textarea" || ctrltypeBKSPACE=="password") && (document.activeElement.getAttribute("readonly")==false))
{	return true;    	}
else{window.event.cancelBubble = true;window.event.returnValue = false;key= false ;return false; 	} }*/

/*ctrlNameBKSPACE = event.srcElement.form;
        if (window.event && window.event.keyCode == 8) {
        //alert("CHECK 3");
     ctrltypeBKSPACE = document.activeElement.getAttribute("type");
//alert(ctrltypeBKSPACE);
 
if((ctrlNameBKSPACE=="[object]"|| ctrlNameBKSPACE=="[object HTMLFormElement]") && ( ctrltypeBKSPACE=="text" || ctrltypeBKSPACE=="TEXT" || document.activeElement.tagName=="TEXTAREA" || ctrltypeBKSPACE=="password") && (document.activeElement.getAttribute("readonly")==null || document.activeElement.getAttribute("readonly")==false))
    {
      return true;
     }
     else{
      alert("This is 8");
       window.event.cancelBubble = true;
          window.event.returnValue = false;
          event.keyCode = false;
          return false;
     }
 
    }*/
if ( key == 112) {window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;  }
if ( key == 113) {window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;  }
if ( key == 114) {window.event.cancelBubble = true;window.event.returnValue = false;key = false; return false; }          
if ( key == 115) {window.event.cancelBubble = true; window.event.returnValue = false; key = false;  return false;}       
if ( key == 116) {window.event.cancelBubble = true; window.event.returnValue = false;key = false;return false;}
if ( key == 117) {window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;}
if ( key == 118) {window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;}
if ( key == 119) {window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false; }
if ( key == 120) { window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;    }
if ( key == 121) { window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;    }
if ( key == 122) { window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;    }
if ( key== 123) { window.event.cancelBubble = true;window.event.returnValue = false;key = false;return false;    }}
function setErrorInfo(errfield, msgval){
if(count == 0) 
if(firstErrorField == "") firstErrorField=errfield; count ++; msg += count + ".   "+ msgval + "<BR>";}
function displayPopup(url, ht, width, x, y)
{	if(x == null)		x = 250;	if(y == null)		y = 150;	popupHandle = null;	popupHandle = open(url,"mydisplay" , "resizable=no,status=no,location=no,scrollbars,height=" + ht + ",width=" + width + ",left=" + x + ",top="+y);	if (!popupHandle.opener) popupHandle.opener = self;}
