$(function(){
  $("input:file").change(function (){
     var fileName = $(this).val().split("\\").pop();
     $(".filename").html(fileName);
     $(".filename").addClass("selected");
     $(".error-label").addClass("hide");
   });

  $(".send-button input").click(function(e){
    var inputFileValue = $(".input-file").val();
    var fileExtension = $(".input-file").val().split(".").pop();
    if(inputFileValue === "" || fileExtension !== "log"){
      $(".error-label").removeClass("hide");
      return false;
    }

    // e.stopPropagation();
    // e.preventDefault();
    //
    // setTimeout(function(){
    //   $("#upload").submit();
    // }, 3000);
  });

  $(".filename").click(function(){
    $(".custom-file-upload").click();
  });

  var createdAT = new Date($(".created-at").html());
  var month = createdAT.getUTCMonth();
  var day = createdAT.getUTCDay();
  var year = createdAT.getUTCFullYear();
  $(".created-at").html(`${day}/${month}/${year}`);
  $(".created-at").removeClass("hide");


  var totalDistance = Math.round($(".total-distance").html() * 100) / 100;
  $(".total-distance").html(totalDistance);
  $(".total-distance").removeClass("hide");
  var onMiles = totalDistance * 0.621371;
  $(".miles-values").html(Math.round(onMiles * 100) / 100);

});
