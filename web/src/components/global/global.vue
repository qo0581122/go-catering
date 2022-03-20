<script>
const GetTypeOptions = [
  // {
  //   value: 0,
  //   label: "已撤销",
  // },
  {
    value: 1,
    label: "直接领取",
  },
  {
    value: 2,
    label: "系统派发",
  },
//   {
//     value: 3,
//     label: "好友转赠",
//   },
];
const GetTypeFormatOptions = [
  {
    value: 0,
    label: "已撤销",
  },
  {
    value: 1,
    label: "直接领取",
  },
  {
    value: 2,
    label: "系统派发",
  },
  {
    value: 3,
    label: "好友转赠",
  },
];
const ValidTimeTypeOptions = [
  { value: 1, label: "按固定时间" },
  { value: 2, label: "按领取后有效期" },
];

const UseTypeOptions = [
  { value: 1, label: "可并用" },
  { value: 2, label: "不可并用" },
];
const UseStatusOptions = [
  { value: 1, label: "未使用" },
  { value: 2, label: "已使用" },
];

const IntegrationChangeOptions = [
  { value: 1, label: "消费获取" },
  { value: 2, label: "活动获取" },
];

const OrderTypeOptions = [
  {value: 1, label:"堂食"},
  {value: 2, label:"外送"},
];

const OrderGetTypeOptions = [
  {value:1, label:"即取"},
  {value:2, label:"15分钟后自取"},
  {value:3, label:"30分钟后自取"},
  {value:4, label:"1小时后自取"},
];

const OrderStatusOptions = [
  {value:0, label:"已取消"},  
  {value:1, label:"待支付"},  
  {value:2, label:"配餐中"},  
  {value:3, label:"配送中"},  
  {value:4, label:"已送达"},  
  {value:5, label:"已完成"},
  {value:6, label:"已评价"}
];

const OrderPayTypeOptions = [
  {value:0,label:"微信"},
  {value:1,label:"支付宝"},
  {value:2,label:"银联卡"},
]

export default {
  GetTypeOptions,
  ValidTimeTypeOptions,
  UseTypeOptions,
  GetTypeFormatOptions,
  UseStatusOptions,
  IntegrationChangeOptions,
  OrderTypeOptions,
  OrderGetTypeOptions,
  OrderStatusOptions,
  OrderPayTypeOptions,
};
</script>