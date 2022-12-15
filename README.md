# jh-weather
Weather project 
    This is a test project to check out the [Open Weather API](https://openweathermap.org/api/) . This takes in two query parameters , latitude and longitude and provides the weather at that location as a series of descriptions. 

Because this API has a limit of 1000 calls per day , it caches data for 90 seconds . 

Rationale for using 90 seconds as the cache limit is determined by two aspects.
1) weather rarely changes minute by minute 
2) The free tier for this API is 1000 calls per day . There are 1440 minutes in a day so roughly 1.5 minutes between calls will ensure that we do not transcend this tier.

Steps to run the application 
1) make jh
2) docker run -d -p <HOST PORT>:80 <IMG_ID> 