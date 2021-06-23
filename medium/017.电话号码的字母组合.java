package view;

import java.util.ArrayList;
import java.util.List;

public class Class2 {
    public Class2() {
        super();
    }

    public static void main(String[] args) {
        Class2 class2 = new Class2();
        System.out.println(class2.letterCombinations(""));
    }
    
    private List<String> letterCombinations(String digits){
        String[][] chars = new String[][]{new String[]{"a", "b", "c"},
                                          new String[]{"d", "e", "f"},
                                          new String[]{"g", "h", "i"},
                                          new String[]{"j", "k", "l"},
                                          new String[]{"m", "n", "o"},
                                          new String[]{"p", "q", "r", "s"},
                                          new String[]{"t", "u", "v"},
                                          new String[]{"w", "x", "y", "z"},
                                          };
        
        
        List<String> retList = new ArrayList<>();
        
        if(digits != null && digits != ""){
            List<String> tmpList = new ArrayList<>();
            tmpList.add("");
            
            for(int i=0; i<digits.length(); i++){
                int idx = Integer.valueOf(String.valueOf(digits.charAt(i))) - 2;
                retList.clear();
                for(int j=0; j<tmpList.size(); j++){
                    for(int k=0; k<chars[idx].length; k++){
                        retList.add(tmpList.get(j) + chars[idx][k]);
                    }
                }
                tmpList.clear();
                for(int m=0; m<retList.size(); m++){
                    tmpList.add(retList.get(m));
                }
            }
        }
        
        return retList;
    }
}
