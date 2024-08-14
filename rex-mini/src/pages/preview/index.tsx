import { View } from "@tarojs/components";
import { useLoad } from "@tarojs/taro";

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  return (
    <View className='container mx-auto'>
      <View className='p-2'>

      </View>
    </View>
  );
}
