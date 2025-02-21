/*Copyright (C) 2020-2030 Supernet Technologies India Private Limited. All rights reserved.
All material appearing on this repository, folder, sub folder and file(s) (“Content/Code etc”) is protected by copyright laws and is the property of Supernet Technologies India Pvt Ltd. You may not copy, reproduce, distribute, publish, display, perform, modify, create derivative works, transmit, or in any way exploit any such content, nor may you distribute any part of this content over any network, including a local area network, sell, offer it for sale, or use such content to construct any kind of database, you may not alter or remove any copyright or other notice from copies of contents on this file.
Copying or storing any contents except as provided by Supernet Technologies India Pvt Ltd is expressly prohibited without prior written permission of the Supernet Technologies.
For permission to use the content, please contact legal@supernet-india.com.
 **/
$(document).ready(function () {
                  
                  
      var name;
             
      $("#translate").change(function() { 
        
      lang = $(this).val();
      window.location.reload();
              
      return lang;
           
              
    });
                
       if(lang=="english"){
                  
            // alert(language)
         var validationMsg = {};
                
            validationMsg.name = "Please enter Name";
            validationMsg.email="Please enter Email";
            validationMsg.mobile="Please enter Mobile Number";
            validationMsg.add="Please enter Address";
            validationMsg.status="Please select Status";
            validationMsg.desc = "Please enter Description";
            validationMsg.code="Please enter Service Code";
            validationMsg.beneficiary = "Please enter Beneficiary";
            validationMsg.source = "Please enter Source";
            validationMsg.applicable = "Please enter Applicable";
            validationMsg.user="Please select User";
            validationMsg.category="Please select Category";
            validationMsg.service="Please select Service";
            validationMsg.biller="Please select Biller";
            validationMsg.channel="Please select Channel";
            validationMsg.acctype="Please select Account Type";
            validationMsg.balance="Please enter Balance";
            validationMsg.accnumber="Please enter Account Number";
            validationMsg.bankname="Please enter Bank Name";
            validationMsg.username="Please select User Name";
            validationMsg.ifcicode="Please enter IFCI Code";
            validationMsg.usertype="Please select User Type";
            validationMsg.add1="Please enter Address 1";
            validationMsg.add2="Please enter Address 2";
            validationMsg.lng="Please select Language";
            
            validationMsg.narration="Please enter Narration";
            validationMsg.ref="Please enter Reference Number";
            validationMsg.clients="Please select client";
            validationMsg.value="Please enter value";
            validationMsg.duration="Please duration in hours";
            validationMsg.old = "Please enter Old Password";
            validationMsg.new="Please enter New Password";
            validationMsg.confirm="Please confirm your New Password";
            
            validationMsg.bpool = "Please select Biller Pool Account";
            validationMsg.bcomm = "Please select Biller Commission Account";
            validationMsg.spool = "Please select Source Pool Account";
            validationMsg.scomm = "Please select Source Commission Account";
            validationMsg.op1 = "Please select Operation";
            validationMsg.type1 = "Please select Type";
   		    validationMsg.val1 = "Please enter Value";
		    validationMsg.maxamount = "Please enter max amount";
			validationMsg.btax = "Please select Beneficary Tax Account";
			validationMsg.stax = "Please select Source Tax Account";
			validationMsg.templatetype = "Please Select Template Type";
			validationMsg.url = "Please enter url";
			validationMsg.title = "Please Select title";
			validationMsg.template1 = "Please enter template 1";
			validationMsg.template2 = "Please enter template 2";
			validationMsg.template3 = "Please enter template 3";
			validationMsg.describeurl = "Please enter describe url";
			validationMsg.accholdname = "Please Enter Account Holder Name";
			validationMsg.phyaccnum = "Please Enter Physical Account Number";
			validationMsg.bankref = "Please Enter Bank Reference Number";
			validationMsg.settlementamount = "Please Enter Settlement Amount";
			validationMsg.operation = "Please Enter Operation";
			validationMsg.tdate = "Please Enter Transaction Date";
			validationMsg.ttime = "Please Enter Transaction Time";
			validationMsg.sfrom = "Please Enter From Date";
			validationMsg.sto = "Please Enter To Date";
			validationMsg.billermetadata = "Please Enter Biller Metadata";
			    validationMsg.rolename = "Please Enter Role Name";
				
			validationMsg.role = "Please Select Role";
			
			validationMsg.operatorname = "Please Enter Operator Name";
			validationMsg.operatorvalue = "Please Enter Operator Value";
			validationMsg.auth2 = "Please Select Minimum 1 Authenticator";
            
            
           
            
                  
       }else{
            // alert(language)
         var validationMsg = {};
                
            validationMsg.name = "Veuillez entrer le nom";
            validationMsg.email="Veuillez saisir un e-mail";
            validationMsg.mobile="Veuillez entrer le numéro de portable";
            validationMsg.add="Veuillez entrer l'adresse";
            validationMsg.status="Veuillez sélectionner le statut";
            validationMsg.desc = "Veuillez entrer la description";
            validationMsg.code="Veuillez entrer le code de service";
            validationMsg.beneficiary = "Veuillez entrer le bénéficiaire";
            validationMsg.source = "Veuillez saisir la source";
            validationMsg.applicable = "Veuillez entrer Applicable";
            validationMsg.user="Veuillez sélectionner l'utilisateur";
            validationMsg.category="Veuillez sélectionner la catégorie";
            validationMsg.service="Veuillez sélectionner Service";
            validationMsg.biller="Veuillez sélectionner Facturation";
            validationMsg.channel="Veuillez sélectionner la chaîne";
            validationMsg.acctype="Veuillez sélectionner le type de compte";
            validationMsg.balance="Veuillez entrer le solde";
            validationMsg.accnumber="Veuillez entrer le numéro de compte";
            validationMsg.bankname="Veuillez entrer le nom de la banque";
            validationMsg.username="Veuillez sélectionner le nom d'utilisateur";
             validationMsg.ifcicode="Veuillez entrer le code IFCI";
            validationMsg.usertype="Veuillez sélectionner le type d'utilisateur";
            validationMsg.add1="Veuillez entrer l'adresse 1";
            validationMsg.add2="Veuillez entrer l'adresse 2";
            validationMsg.lng="Veuillez sélectionner la langue";
            
            validationMsg.narration="Veuillez entrer la narration";
            validationMsg.ref="Veuillez entrer le numéro de référence";
            validationMsg.clients="Veuillez sélectionner un client";
            validationMsg.value="Veuillez entrer une valeur";
            validationMsg.duration="Veuillez durée en heures";
            validationMsg.old = "Veuillez saisir l'ancien mot de passe";
            validationMsg.new="Veuillez saisir un nouveau mot de passe";
            validationMsg.confirm="Veuillez confirmer Votre nouveau mot de passe";
            
             validationMsg.bpool = "Veuillez sélectionner le compte du pool de facturation";
            validationMsg.bcomm = "Veuillez sélectionner le compte de commission de facturation";
            validationMsg.spool = "Veuillez sélectionner le compte de pool source";
            validationMsg.scomm = "Veuillez sélectionner le compte de commission source";
            validationMsg.op1 = "Veuillez sélectionner l'opération";
			 validationMsg.type1 = "Veuillez sélectionner Type";
			validationMsg.val1 = "Veuillez saisir une valeur";
			validationMsg.maxamount = "Veuillez entrer le montant maximum";
			validationMsg.btax = "Veuillez sélectionner le compte fiscal du bénéficiaire";
			validationMsg.stax = "Veuillez sélectionner le compte de taxe source";
			validationMsg.templatetype = "Veuillez sélectionner le type de modèle";
			validationMsg.url = "Veuillez entrer l'url";
			validationMsg.title = "Veuillez sélectionner le titre";
			validationMsg.template1 = "Veuillez entrer le modèle 1";
			validationMsg.template2 = "Veuillez entrer le modèle 2";
			validationMsg.template3 = "Veuillez entrer le modèle 3";
			validationMsg.describeurl = "Veuillez saisir l'URL de description";
			validationMsg.accholdname = "Veuillez entrer le nom du titulaire du compte";
			validationMsg.phyaccnum = "Veuillez saisir le numéro de compte physique";
			validationMsg.bankref = "Veuillez entrer le numéro de référence bancaire";
			validationMsg.settlementamount = "Veuillez entrer le montant du règlement";
			validationMsg.operation = "Veuillez entrer l'opération";
			validationMsg.tdate = "Veuillez saisir la date de la transaction";
			validationMsg.ttime = "Veuillez saisir l'heure de la transaction";
			validationMsg.sfrom = "Veuillez saisir la date de début";
			validationMsg.sto = "Veuillez entrer à ce jour";
			validationMsg.billermetadata = "Veuillez saisir les métadonnées du fournisseur";
			validationMsg.rolename = "Veuillez saisir le nom du rôle";
			validationMsg.role = "Veuillez sélectionner un rôle";
			
			validationMsg.operatorname = "Veuillez entrer le nom de l'opérateur";
			validationMsg.operatorvalue = "Veuillez entrer la valeur de l'opérateur";
			validationMsg.auth2 = "Veuillez sélectionner au moins 1 authentificateur";
			
           
            
            
            
            
            
            
                  
            }
                
                $('#addUsersValidate').validate({
                rules: {
                	
					//system users
                input_full_name: {
                required: true,
                maxlength: 30
                },
                input_email: {
                required: true,
                maxlength: 50,
                },
                input_mobile: {
                required: true,
                maxlength: 10,
                },
                input_address: {
                required: true,
                maxlength: 40,
                },
                input_status: {
                required: true,
                maxlength: 40,
                },
                input_language: {
                required: true,
                maxlength: 40,
                },
                input_role: {
                required: true,
                maxlength: 40,
                },
                
                
                //user
                
                input_user_type: {
                required: true,
                maxlength: 40,
                },
                input_address1: {
                required: true,
                maxlength: 40,
                },
                input_address2: {
                required: true,
                maxlength: 40,
                },
                input_language: {
                required: true,
                maxlength: 40,
                },
                
                
                    
                
                
                //usertype
                
                input_first_name: {
                required: true,
                maxlength: 30
                },
                input_desc: {
                required: true,
                maxlength: 40
                },
                input_usertype: {
                required: true,
                maxlength: 40,
                },
                
                //Channel
                input_name: {
                required: true,
                maxlength: 30
                },
                input_description: {
                required: true,
                maxlength: 50,
                },
                
                //service
                input_service_code: {
                required: true,
                maxlength: 30,
                },
                
				//role
                
                input_role_name: {   
                required: true,
                maxlength: 30
                },
                
                //ChargeMnemonic
                input_beneficiary: {
                required: true,
                maxlength: 30
                },
                input_source: {
                required: true,
                maxlength: 30
                },
                input_applicable: {
                required: true,
                maxlength: 30
                },
                input_beneficary_pool: {
                required: true,
              
                },
                  input_beneficary_commision: {
                required: true,
               
                },
                  input_source_pool: {
                required: true,
            
                },
                  input_source_commision: {
                required: true,
              
                },
                  input_operation1: {
                required: true,
             
                },
                  input_operation2: {
                required: true,
          
                },
                  input_operation3: {
                required: true,
             
                },
                  input_type1: {
                required: true,
              
                },
                  input_type2: {
                required: true,
            
                },
                  input_type3: {
                required: true,
            
                },
                  input_value1: {
                required: true,
                maxlength: 10,
                },
                  input_value2: {
                required: true,
                maxlength: 10,
                },
                  input_value3: {
                required: true,
                maxlength: 10,
                },
                input_beneficary_tax: {
                required: true,
                maxlength: 25,
                },
                input_source_tax: {
                required: true,
                maxlength: 25,
                },
                
                //package
                
                input_user: {
                required: true,
                maxlength: 25,
              
                },
                  input_category: {
                required: true,
                maxlength: 25,
           
                },
                  input_service: {
                required: true,
                maxlength: 25,
          
                },
                  input_biller: {
                required: true,
                maxlength: 25,
             
                },
                  input_channel: {
                required: true,
                maxlength: 25,
          
                },
                input_client: {
                required: true,
                maxlength: 25,
          
                },
                
                
                //account
                
                input_user_name: {
                required: true,
                maxlength: 25,
              
                },
                  input_account_type: {
                required: true,
                maxlength: 25,
           
                },
                  input_balance: {
                required: true,
                maxlength: 25,
          
                },
                  input_account_number: {
                required: true,
                maxlength: 25,
             
                },
                  input_bank_name: {
                required: true,
                maxlength: 25,
          
                },
                
                input_narration: {
                required: true,
                maxlength: 50,
             
                },
                  input_reference_number: {
                required: true,
                maxlength: 25,
          
                },
                input_user_account_number: {
                required: true,
                maxlength: 25,
          
                },
                input_user_bank_name: {
                required: true,
                maxlength: 25,
          
                },
                input_user_ifci_code: {
                required: true,
                maxlength: 25,
          
                },
                
                //limits
                
                input_volumnechannel: {
                required: true,
                maxlength: 25,
              
                },
                  input_volume_value: {
                required: true,
                maxlength: 10,
           
                },
                  input_volume_duration: {
                required: true,
                maxlength: 25,
          
                },
                input_velocitychannel: {
                required: true,
          
                },
                
                input_velocity_value: {
                required: true,
                maxlength: 10,
             
                },
                  input_velocity_duration: {
                required: true,
                maxlength: 25,
          
                },
                input_max_amount: {
                required: true,
                maxlength: 10,
          
                },
                
                //changepassword
                
                input_oldpassword: {
                required: true,
              
                },
                input_newpassword: {
                required: true,
                
                },
                input_confirmnewpassword: {
                required: true,
                
                },
                
                //NotificationTemplate
                input_template_type: {
                required: true,
                maxlength: 25,
                
                },
                input_url: {
                required: true,
                
                },
                 input_title: {
                required: true,
                maxlength: 100,
                
                },
                input_template1: {
                required: true,
                maxlength: 200,
                
                },
                input_template2: {
                required: true,
                maxlength: 200,
                
                },
                input_template3: {
                required: true,
                maxlength: 200,
                
                },
                input_describe_url: {
                required: true,
                maxlength: 200,
                
                },
                
                //settlement
                
                input_account_holdername: {
                required: true,
                maxlength: 50,
                
                },
//                input_physical_account_number: {
//                required: true,
                
//                },
                input_settlement_amount: {
                required: true,
                 maxlength: 10,
                
                },
                input_bank_reference_number: {
                required: true,
                maxlength: 50,
                
                },
                input_operation: {
                required: true,
                maxlength: 25,
                
                },
                input_transaction_date: {
                required: true,
                
                },
                input_transaction_time: {
                required: true,
                
                },
                input_settlement_from: {
                required: true,
                
                },
                 input_settlement_to: {
                required: true,
                
                },
                    biller_metadata: {
                required: true,
                
                },
                
                 input_operator_name: {
                required: true,
                
                },
                 input_operator_value: {
                required: true,
                
                },
                 auth2: {
                required: true,
                
                },
                  
                
                
                },
                messages: {
                	
					//system users
                input_full_name: {
                required:validationMsg.name,
                },
                input_email: {
                required: validationMsg.email,
                },
                input_mobile: {
                required: validationMsg.mobile,
                },
                input_address: {
                required:validationMsg.add,
                },
                input_status: {
                required:validationMsg.status,
                },
                input_language: {
                required:validationMsg.lng,
                },
                input_role: {
                required:validationMsg.role,
                },
                
                
                //user
                input_user_type: {
                required:validationMsg.usertype,
                },
                input_address1: {
                required:validationMsg.add1,
                },
                input_address2: {
                required:validationMsg.add2,
                },
                
                
              
                //usertype
                input_first_name: {
                required:validationMsg.name,
                },
                 input_desc: {
                required:validationMsg.desc,
                },
                input_usertype: {
                required:validationMsg.usertype,
                },
                
                //channel
                input_name: {
                required: validationMsg.name,
                },
                input_description: {
                required: validationMsg.desc,
                },
                
                //service
                input_service_code: {
                required: validationMsg.code,
                },
                
              //role
                input_role_name: {
                required:validationMsg.name,
                },
                
                //ChargeMnemonic
                input_beneficiary: {
                required: validationMsg.beneficiary,
                },
                input_source: {
                required: validationMsg.source,
                },
                input_applicable: {
                required: validationMsg.applicable,
                },
                      input_beneficary_pool: {
                required: validationMsg.bpool,
                },
                   input_beneficary_commision: {
                required: validationMsg.bcomm,
                },
                   input_source_pool: {
                required: validationMsg.spool,
                },
                     input_source_commision: {
                required: validationMsg.scomm,
                },
                   input_operation1: {
                required: validationMsg.op1,
                },
                   input_operation2: {
                required: validationMsg.op1,
                },
                   input_operation3: {
                required: validationMsg.op1,
                },
                   input_type1: {
                required: validationMsg.type1,
                },
                   input_type2: {
                required: validationMsg.type1,
                },
                   input_type3: {
                required: validationMsg.type1,
                },
                   input_value1: {
                required: validationMsg.val1,
                },
                   input_value2: {
                required: validationMsg.val1,
                },
                   input_value3: {
                required: validationMsg.val1,
                },
                input_beneficary_tax: {
                required: validationMsg.btax,
                },
                input_source_tax: {
                required: validationMsg.stax,
                },

                
                //package
                input_user: {
                required: validationMsg.user,
                },
                input_category: {
                required: validationMsg.category,
                },
                input_service: {
                required: validationMsg.service,
                },
                input_biller: {
                required: validationMsg.biller,
                },
                input_channel: {
                required: validationMsg.channel,
                },
                input_client: {
                required: validationMsg.clients,
                },
                
                
                //account
                input_user_name: {
                required: validationMsg.username,
                },
                input_account_type: {
                required: validationMsg.acctype,
                },
                input_balance: {
                required: validationMsg.balance,
                },
                input_account_number: {
                required: validationMsg.accnumber,
                },
                input_bank_name: {
                required: validationMsg.bankname,
                },
                
                input_narration: {
                required: validationMsg.narration,
                },
                input_reference_number: {
                required: validationMsg.ref,
                },
                input_user_account_number: {
                required: validationMsg.accnumber,
                },
                input_user_bank_name: {
                required: validationMsg.username,
                },
                input_user_ifci_code: {
                required: validationMsg.ifcicode,
                },
                
                
                
                //limits
                input_volumnechannel: {
                required: validationMsg.channel,
                },
                input_volume_value: {
                required: validationMsg.value,
                },
                input_volume_duration: {
                required: validationMsg.duration,
                },
                input_velocitychannel: {
                required: validationMsg.channel,
                },
                
                input_velocity_value: {
                required: validationMsg.value,
                },
                input_velocity_duration: {
                required: validationMsg.duration,
                },
                input_max_amount: {
                required: validationMsg.maxamount,
                },
                
                //changepassword
                
                input_oldpassword: {
                required: validationMsg.old,
                },
                input_newpassword: {
                required: validationMsg.new,
                },
                input_confirmnewpassword: {
                required: validationMsg.confirm,
                },
                
                //NotificationTemplate
                input_template_type: {
                required: validationMsg.templatetype,
                },
                input_url: {
                required: validationMsg.url,
                },
                 input_title: {
                 required: validationMsg.title,
                
                },
                input_template1: {
                 required: validationMsg.template1,
                
                },
                input_template2: {
                 required: validationMsg.template2,
                
                },
                input_template3: {
                 required: validationMsg.template3,
                
                },
                input_describe_url: {
                 required: validationMsg.describeurl,
                
                },
                //settlement
                
                input_account_holdername: {
                required: validationMsg.accholdname,
                
                },
//                input_physical_account_number: {
//                required: validationMsg.phyaccnum,
                
//                },
                input_bank_reference_number: {
                required: validationMsg.bankref,
                
                },
                 input_settlement_amount: {
                required: validationMsg.settlementamount,
                
                },
                input_operation: {
                required: validationMsg.operation,
                
                },
                input_transaction_date: {
                required: validationMsg.tdate,
                
                },
                input_transaction_time: {
                required: validationMsg.ttime,
                
                },
                input_settlement_from: {
                required: validationMsg.sfrom,
                
                },
                 input_settlement_to: {
                required: validationMsg.sto,
                
                },
                biller_metadata: {
                required: validationMsg.billermetadata,
                
                },
                
                  input_operator_name: {
                required: validationMsg.operatorname,
                
                },
                 input_operator_value: {
                required: validationMsg.operatorvalue,
                
                },
                 auth2: {
               required: validationMsg.auth2,
                
                },
                
                
                
                },
                errorElement: 'span',
                errorPlacement: function (error, element) {
                error.addClass('invalid-feedback');
                //error.addClass('lang');
               // error.attr('key','FIELDEMPTY');
                element.closest('.form-group').append(error);
                },
                highlight: function (element, errorClass, validClass) {
                $(element).addClass('is-invalid');
                },
                unhighlight: function (element, errorClass, validClass) {
                $(element).removeClass('is-invalid');
                }
                });
                });
