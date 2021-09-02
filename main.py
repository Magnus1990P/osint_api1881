#!/usr/bin/env python3
#! coding: utf-8

#import http.client, urllib.request, urllib.parse, urllib.error, base64
from requests import get
import urllib.parse
import base64
import bs4 as bs
import json
import configparser

#soup = bs.BeautifulSoup(source, "html")

config = configparser.ConfigParser()
config.read("config.ini")





class api1881:
	API_KEYA = None
	API_KEYB = None
	API_URL  = None

	HTTP_HEAD = { 'Ocp-Apim-Subscription-Key': None,}

	def __init__(self, key_a=None, key_b=None, url="https://services.api1881.no/" ):
		if None not in [key_a, key_b]:
			self.API_KEYA = key_a
			self.API_KEYB = key_b
		self.API_URL = url

		self.HTTP_HEAD['Ocp-Apim-Subscription-Key'] = self.API_KEYA

	def run_query(self, url_context=None, params=None, data=None):
		if params:
			params = urllib.parse.urlencode( params, safe="" )

		if data:
			data = urllib.parse.urlencode( data )

		url = self.API_URL + url_context
		if params:
			url += "?{}".format(params)
		
		resp = get(url, headers=self.HTTP_HEAD)
		if resp.status_code == 200:
			print("{} - {}".format(resp.status_code, json.loads(resp.text)))
			return json.loads(resp.content)
		print("{} - {}".format(resp.status_code, resp.text))
		return None


	def address_to_coordinates(self,address=None):
		if not address:
			return None
		params = { "address" : address }

		res = self.run_query(url_context="/geocoding/address", params=params)
		return res
		

	def coordinates_to_address(self, lat=None, lon=None):
		if not X or not Y:
			return None
		params = { "latitude" : lat, "longtitude": lon }
		res = self.run_query(url_context="/geocoding/address", params=params)
		return res

	def orgnumber_lookup(self, orgnr=None):
		if not orgnr or not orgnr.isdigit():
			return None
		
		res = self.run_query(url_context="/lookup/organizationnumber/{}".format(orgnr))
		return res

	def phonenumber_lookup(self, phone=None):
		if not phone:
			return None
		
		res = self.run_query(url_context="/lookup/phonenumber/{}".format(phone))
		return res


	def __str__(self,):
		out = "{:<30}{}\n".format("API URL", self.API_URL)
		out += "{:<30}{}\n".format("API KEY", self.API_KEYA)
		out += "{:<30}{}\n".format("API KEY", self.API_KEYB)
		for k in self.HTTP_HEAD:
			out += "{:<30}{}\n".format( k, self.HTTP_HEAD[k] ) 
		return out


A = api1881(key_a=config["API"]["KEY_A"], key_b=config["API"]["KEY_B"],
        url=config["API"]["URL"])



