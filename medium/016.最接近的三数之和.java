package view;

public class Class3 {
    public Class3() {
        super();
    }

    public static void main(String[] args) {
        Class3 class3 = new Class3();
        System.out.println(class3.threeSumClosest(new int[]{-1, 2, 1, -4}, 1));
    }
    
    public int threeSumClosest(int[] nums, int target){
        int diff = Integer.MAX_VALUE;
        int sum = 0;
        
        for(int i=0; i<nums.length; i++){
            for(int j=i+1; j<nums.length; j++){
                for(int k=j+1; k<nums.length; k++){
                    int m = Math.abs(nums[i] + nums[j] + nums[k] - target);
                    if(m < diff){
                        sum = nums[i] + nums[j] + nums[k];
                        diff = m;
                    }
                }
            }
        }
        
        return sum;
    }
}
