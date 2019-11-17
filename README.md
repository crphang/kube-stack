# Introduction

Basic Microservice setup with kubernetes. 

1. API Gateway using Ambassador that routes to authentication service
1. Authentication service that does nothing but checks headers
1. Simple App that returns the index page with nothing but header information from auth service
