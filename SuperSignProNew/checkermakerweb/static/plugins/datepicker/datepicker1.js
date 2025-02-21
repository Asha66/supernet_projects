'use strict';

  // The default language is English
var lang = "english";
// Check for localStorage support
if('localStorage' in window){
   
   var usrLang = localStorage.getItem('uiLang');
   if(usrLang) {
       lang = usrLang
   }

}
	
$(document).ready(function() {
	
	 // console.log("lang vaue from datepicker.js",lang);
	
        $("#translate").change(function() { 
        var lang = $('#translate :selected').val();
          // update localStorage key
          if('localStorage' in window){
               localStorage.setItem('uiLang', lang);
//               console.log( localStorage.getItem('uiLang') );
          }
        
          window.location.reload();
        //var lang = $('#translate :selected').val();
       
	
return
        
        });
  
    
    var TodayLabel, YesterdayLabel, thisMonthLabel, lastMonthLabel, last7DaysLabel, last30DaysLabel, customRangeLabel, applyLabel, cancelLabel, daysOfWeek, monthNames, last3MonthLabel;
    if (lang == "english") {
        TodayLabel = "Today";
        YesterdayLabel = "Yesterday";
        thisMonthLabel = "This Month";
        lastMonthLabel = "Last Month";
        last3MonthLabel = "Last 3 Month";
        last7DaysLabel = "Last 7 Days";
        last30DaysLabel = "Last 30 Days";
        customRangeLabel = "Custom Range";
        applyLabel = "Apply";
        cancelLabel = "Cancel";
        daysOfWeek = ["Su", "Mo", "Tu", "We", "Th", "Fr", "Sa"];
        monthNames = [
            "January",
            "February",
            "March",
            "April",
            "May",
            "June",
            "July",
            "August",
            "September",
            "October",
            "November",
            "December"
        ];
    } else{
        TodayLabel = "Aujourd'hui";
        YesterdayLabel = "Hier";
        thisMonthLabel = "Ce mois-ci";
        lastMonthLabel = "Le mois dernier";
        last3MonthLabel = "Le 3 mois dernier";
        last7DaysLabel = "Les 7 derniers jours";
        last30DaysLabel = "Les 30 derniers jours";
        customRangeLabel = "Période";
        applyLabel = "Valider";
        cancelLabel = "Annuler";
        daysOfWeek = ["Do", "Lu", "Ma", "Me", "Je", "Ve", "Sa"];
        monthNames = [
            "Janvier",
            "Février",
            "Mars",
            "Avril",
            "Mai",
            "Juin",
            "Juillet",
            "Août",
            "Septembre",
            "Octobre",
            "Novembre",
            "Décembre"
        ];
    }

    $('input[name="single-date-picker"]').daterangepicker({
        singleDatePicker: true,
        showDropdowns: true,
        locale: {
            format: "MM/DD/YYYY"
        }
    });

    $('.simple-date-range-picker').daterangepicker({
        autoUpdateInput: true,
        showDropdowns: true,
        linkedCalendars: false,
        startDate: moment().subtract(2, 'month').startOf('month').format('MM/DD/YYYY'),
        ranges: {
            [TodayLabel]: [moment(), moment()],
            [YesterdayLabel]: [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
            [last7DaysLabel]: [moment().subtract(6, 'days'), moment()],
            [last30DaysLabel]: [moment().subtract(29, 'days'), moment()],
            [thisMonthLabel]: [moment().startOf('month'), moment().endOf('month')],
            [lastMonthLabel]: [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
            [last3MonthLabel]: [moment().subtract(2, 'month').startOf('month'), moment().endOf('month')]
        },
        locale: {
            customRangeLabel: customRangeLabel,
            applyLabel: applyLabel,
            cancelLabel: cancelLabel,
            format: "MM/DD/YYYY",
            daysOfWeek: daysOfWeek,
            monthNames: monthNames
        },
    });
    
    $('.date-range-picker').daterangepicker({
        autoUpdateInput: true,
        showDropdowns: true,
        linkedCalendars: false,
        startDate: moment().subtract(2, 'month').startOf('month').format('MM/DD/YYYY'),
        ranges: {
            [TodayLabel]: [moment(), moment()],
            [YesterdayLabel]: [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
            [last7DaysLabel]: [moment().subtract(6, 'days'), moment()],
            [last30DaysLabel]: [moment().subtract(29, 'days'), moment()],
            [thisMonthLabel]: [moment().startOf('month'), moment().endOf('month')],
            [lastMonthLabel]: [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')],
            [last3MonthLabel]: [moment().subtract(2, 'month').startOf('month'), moment().endOf('month')]
        },
        locale: {
            customRangeLabel: customRangeLabel,
            applyLabel: applyLabel,
            cancelLabel: cancelLabel,
            format: "MM/DD/YYYY",
            daysOfWeek: daysOfWeek,
            monthNames: monthNames
        },
    });

    $('.simple-date-range-picker').on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('MM/DD/YYYY') + ' - ' + picker.endDate.format('MM/DD/YYYY'));
    });
    $('.simple-date-range-picker').on('cancel.daterangepicker', function(ev, picker) {
        $(this).val('');
    });

    $('input[name="simple-date-range-picker-callback"]').daterangepicker({
        opens: 'left'
    }, function(start, end, label) {
        swal("A new date selection was made", start.format('MM/DD/YYYY') + ' to ' + end.format('MM/DD/YYYY'), "success")
    });
    
     $('.date-range-picker').on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('MM/DD/YYYY') + ' - ' + picker.endDate.format('MM/DD/YYYY'));
    });
    $('.date-range-picker').on('cancel.daterangepicker', function(ev, picker) {
        $(this).val('');
    });

    $('input[name="date-range-picker-callback"]').daterangepicker({
        opens: 'left'
    }, function(start, end, label) {
        swal("A new date selection was made", start.format('MM/DD/YYYY') + ' to ' + end.format('MM/DD/YYYY'), "success")
    });

    $('input[name="datetimes"]').daterangepicker({
        timePicker: true,
        startDate: moment().startOf('hour'),
        endDate: moment().startOf('hour').add(32, 'hour'),
        locale: {
            format: 'M/DD hh:mm A'
        }
    });

    /**
     * datefilter
     */
    var datefilter = $('input[name="datefilter"]');
    datefilter.daterangepicker({
        autoUpdateInput: false,
        locale: {
            cancelLabel: 'Clear'
        }
    });

    datefilter.on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('MM/DD/YYYY') + ' - ' + picker.endDate.format('MM/DD/YYYY'));
    });

    $('input.create-event-datepicker').daterangepicker({
        singleDatePicker: true,
        showDropdowns: true,
        autoUpdateInput: false
    }).on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('MM/DD/YYYY'));
    });
 
        
});

