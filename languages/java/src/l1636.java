import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

public class l1636 {
    public static void main(String[] args) {
      int[] nums = new int[]{1,1,2,2,2,3};
        System.out.println(Arrays.toString(frequencySort(nums)));
    }
    public static int[] frequencySort(int[] nums) {
        Map<Integer,Integer> mp = new HashMap<>();
        for (int num:nums) {
            mp.put(num,mp.getOrDefault(num,0)+1);
        }
        return Arrays.stream(nums).boxed().sorted(
                (a,b)->{
                    if(mp.get(a) != mp.get(b)){
                        return mp.get(a) - mp.get(b);
                    }
                    return b-a;
                }).mapToInt(Integer::valueOf).toArray();

    }
}
