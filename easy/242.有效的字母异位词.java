package com.ranbo;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class Main {

    public static void main(String[] args) {
	    Main main = new Main();
	    String str1 = "StatueOfLiberty";
	    String str2 = "BuiltToStayFree";
	    System.out.println(main.isAnagram2(str1, str2));
    }
    private boolean isAnagram2(String str1, String str2){
        if(str1.length() != str2.length()){
            return false;
        }

        if(str1.length()<2){
            return true;
        }
        str1 = str1.toLowerCase();
        str2 = str2.toLowerCase();

        Map<Character, Integer> map1 = this.countString(str1);
        Map<Character, Integer> map2 = this.countString(str2);

        for(Map.Entry<Character, Integer> entry : map1.entrySet()){
            if(!entry.getValue().equals(map2.get(entry.getKey()))){
                return false;
            }
        }
        return true;
    }

    private Map<Character, Integer> countString(String str){
        Map<Character, Integer> retMap = new HashMap<>();
        for(char c : str.toCharArray()){
            retMap.merge(c, 1, Integer::sum);
        }
        return retMap;
    }

    private boolean isAnagram1(String str1, String str2){
        str1 = str1.toLowerCase();
        str2 = str2.toLowerCase();

        Integer[] arr = new Integer[str1.length()];
        Arrays.fill(arr, 0);
        for(char c : str2.toCharArray()){
            int idx = str1.indexOf(c);
            while(idx > -1){
                if(arr[idx] == 1){
                    int tmp = idx;
                    idx = str1.substring(tmp+1).indexOf(c);
                    if(idx == -1) break;
                    idx = tmp + idx + 1;
                }else{
                    arr[idx] = 1;
                    break;
                }
            }
        }
        for(int i : arr){
            if(i==0) return false;
        }
        return true;
    }
}
