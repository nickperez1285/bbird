// src/components/Weather.js
import React, { useState, useEffect } from 'react';
import axios from 'axios';

const Weather = () => {
    const [weather, setWeather] = useState(null);
    const apiUrl = 'https://default-blackbird-nicks-organization-b73f0-0.blackbird-relay.a8r.io/newspaper-api/weather';

    useEffect(() => {
        axios
            .get(apiUrl)
            .then((response) => {
                setWeather(response.data);
            })
            .catch((error) => {
                console.error('Error fetching weather data:', error);
            });
    }, []);

    return (
        <div>
            <h2>Weather</h2>
            {weather ? (
                <div>
                    <p><strong>Humidity:</strong> {weather.humidity}</p>
                    <p><strong>Temperature:</strong> {weather.temperature}°C</p>
                    <p><strong>Condition:</strong> {weather.condition}</p>
                    <p><strong>Wind:</strong> {weather.windSpeed}</p>
                </div>
            ) : (
                <p>Loading weather data...</p>
            )}
        </div>
    );
};

export default Weather;
