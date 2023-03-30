# Factory Method (工廠方法模式)
:label: Creational Pattern

## Factory Method Intent (工廠方法模式定義)
- 一種創造性的模式，定義了一個創造物件的介面(interface)，並且該介面定義了一個創造物件的方法，讓繼承並實作該介面的類別來決定要創造什麼樣的物件或實例化哪一個類別
- 工廠方法可以讓一個類別將實例化的工作推遲到繼承該介面的子類別

## Factory Method Structure (工廠方法模式結構)
![image](./structure.png)

## Factory Method Known Uses (工廠方法的適用場景)
* ### 在Coding的「過程」中，如果當下還沒有辦法知道程式要用到的「切確類別」及其依賴關係時：
  - 工廠方法將**創建物件「TheObject」的程式**與**實際使用「TheObject」的程式**分離，進而在不影響其它程式的情況下，擴充**創造「TheObject**的程式
  - 如果我們需要一個新的「TheObject」:「ConcreteObjectC」的時候，我們只需要新增一個實作「Creator」介面的工廠物件即可
* ### 如果我們要讓使用端來擴充「package or 框架」的內部組件時：
  - 繼承可能是擴展「軟體 or 框架」預設行為最簡單的方法，但是當我們使用子類別替代標準組件時，框架該如何分辦出該子類別？
    - 將各框架中的構造組件程式碼集中到單個工廠方法中，除了繼承該組件之外，可以允許任何人對該方法進行覆寫
  - 假設我們使用一個Open Source UI 框架編寫自已的應用程式，我們希望在程式中使用圓形的按鈕，但是原來的框架只支援方形按鈕：
    - 我們可以使用圓形按鈕的類別來繼承**按鈕類別**，但是我們需要告訴UI框架類別使用新的按鈕類別來代替預設按鈕
    - 為了實現這個功能，我們可以根據基礎框架類別開發子類別：**圓形按鈕**，並且覆寫框架子類別的CreateButton創建按鈕方法。
    - 基礎類別的CreateButton方法return「按鈕Interface」，而我們開發的框架子類別就可以return**圓形按鈕**
    - 接著我們就可以使用圓形按鈕UI類別來代替UI框架類別
* ### 如果我們希望重複利用現有物件來節省系統資源，而不是每次都要重新創建物件時：
  - 在處理大型資源密集物件，比如：Database連接、文件系統or網路資源時，我們會經常碰到這種資源需求
  - 所以要從「重複使用現有物件」的角度去思考
    * 我們需要創建儲存空間來存放所有已經創建的物件
    * 當外部使用者呼叫請求一個物件時，程式將在物件池中尋找可用的物件，然後將這個可用物件return給外部使用者
    * 如果物件池中沒有可供使用的物件，程式則要創建一個新物件，並把這個新物件放到物件池中
    * 這些程式碼需要位在同一個地方，確保重複的程式碼不會污染程式
    * 最直接的方式，就是將這些程式碼放置在我們試重複使用物件類別的建構式中，但是從定義上來說，建構式始終return的是新物件，無法return現有的實例
    * 因此，我們需要一個「既能創建新物件，又可以重複利用現有物件」的普通方法，這看起來非常像工廠方法



## Factory Method Motivation (情境,移到example裡面)

## Factory Method Applicability



## Factory Method Participants

## Factory Method Collaborations

## Factory Method Consequences

## Factory Method Implementations/Simple Code

## Factory Method Related Patterns