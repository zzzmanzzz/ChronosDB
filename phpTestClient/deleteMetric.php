<?php
$ch = curl_init();
$host =  "http://127.0.0.1:13000";
$inputUrl = $host . "/api/v1/metric?metric=wyatt_new";


curl_setopt($ch, CURLOPT_URL ,$inputUrl);
curl_setopt($ch, CURLOPT_HEADER, true);
curl_setopt($ch,CURLOPT_CUSTOMREQUEST, "DELETE");
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
$response = curl_exec($ch);
$header_size = curl_getinfo($ch, CURLINFO_HEADER_SIZE);
$header = substr($response, 0, $header_size);
$body = substr($response, $header_size);
var_dump($response);
var_dump(json_decode($body));
curl_close($ch);



