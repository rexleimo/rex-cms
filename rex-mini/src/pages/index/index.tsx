import { View } from "@tarojs/components";
import { useLoad } from "@tarojs/taro";
import "./index.less";

export default function Index() {
  useLoad(() => {
    console.log("Page loaded.");
  });

  return (
    <View className='index-page'>
      <View></View>
    </View>
  );
}
