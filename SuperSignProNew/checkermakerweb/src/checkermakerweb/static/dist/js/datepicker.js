/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 **/
'use strict';

  $("#translate").change(function() { 
        var lang = $('#translate :selected').val();
      
          // update localStorage key
          if('localStorage' in window){
               localStorage.setItem('uiLang', lang);
//               console.log( localStorage.getItem('uiLang') );
          }
          $(".lang").each(function(index, element) {
            $(this).text(arrLang[lang][$(this).attr("key")]);
            $(this).prop("placeholder",arrLang[lang][$(this).attr("key")]);
          });
        return
        });
	
$(document).ready(function() {
	
	  //console.log("lang vaue from datepicker.js",lang);
	

	

  
    
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
            format: "YYYY/MM/DD"
        }
    });

    $('.simple-date-range-picker').daterangepicker({
        autoUpdateInput: true,
        showDropdowns: true,
        linkedCalendars: false,
        startDate: moment().subtract(2, 'month').startOf('month').format('YYYY/MM/DD'),
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
            format: "YYYY/MM/DD",
            daysOfWeek: daysOfWeek,
            monthNames: monthNames
        },
    });
    
    $('.date-range-picker').daterangepicker({
        autoUpdateInput: true,
        showDropdowns: true,
        linkedCalendars: false,
        startDate: moment().subtract(2, 'month').startOf('month').format('YYYY/MM/DD'),
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
            format: "YYYY/MM/DD",
            daysOfWeek: daysOfWeek,
            monthNames: monthNames
        },
    });

    $('.simple-date-range-picker').on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('YYYY/MM/DD') + ' - ' + picker.endDate.format('YYYY/MM/DD'));
    });
    $('.simple-date-range-picker').on('cancel.daterangepicker', function(ev, picker) {
        $(this).val('');
    });

    $('input[name="simple-date-range-picker-callback"]').daterangepicker({
        opens: 'left'
    }, function(start, end, label) {
        swal("A new date selection was made", start.format('YYYY-MM-DD') + ' to ' + end.format('YYYY-MM-DD'), "success")
    });
    
     $('.date-range-picker').on('apply.daterangepicker', function(ev, picker) {
        $(this).val(picker.startDate.format('YYYY/MM/DD') + ' - ' + picker.endDate.format('YYYY/MM/DD'));
    });
    $('.date-range-picker').on('cancel.daterangepicker', function(ev, picker) {
        $(this).val('');
    });

    $('input[name="date-range-picker-callback"]').daterangepicker({
        opens: 'left'
    }, function(start, end, label) {
        swal("A new date selection was made", start.format('YYYY-MM-DD') + ' to ' + end.format('YYYY-MM-DD'), "success")
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
