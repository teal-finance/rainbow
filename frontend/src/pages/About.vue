<template>
  <div>
    <section class="text-foreground dark:text-foreground-dark">
      <div class="container flex flex-col items-center px-5 mx-auto py-14 md:flex-row">
        <div
          class="flex flex-col items-center text-center lg:flex-grow md:w-1/2 lg:pl-24 md:pl-16 md:items-start md:text-left">
          <h1 class="mb-4 font-bold text-9xl sm:text-4xl text-rainbow">Rainbow</h1>
          <p class="mb-8 leading-8 sm:w-full">Dashboard for DEX/CEX options trading highlighting market opportunities to
            ease strategies decision. To achieve this goal, Rainbow continuously scans CEX/DEX/DeFi market data.
            This project has been initiated by <a href="https://teal.finance">Teal.Finance</a> and has already won two
            prices during the Ethereum's EthGlobal Hackathon.</p>
        </div>
        <div class="w-1/2 mb-10 md:mb-0">
          <img class="object-cover object-center w-3/4 rounded" alt="Rainbow dashboard" src="@/assets/dashboard.svg">
        </div>
      </div>
    </section>
    <!-- Vision section -->
    <section>
      <div class="container flex flex-wrap px-5 py-24 mx-auto">
        <div class="flex flex-col w-full mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium">In progress</h1>
          <p class="mx-auto leading-relaxed lg:w-2/3">Our target is to be the most helpful possible,
            thus offering the highest possible user experience.Please open a GitHub issue or contact us at <a
              class="text-green-500" href="mailto:Teal.Finance@protonmail.com">Teal.Finance@protonmail.com</a> to
            suggest
            ideas or to share your feeling.To improve the decision-making, we are currently implementing more data
            retrieval and computing indicators to provide:</p>
        </div>
        <div class="flex flex-wrap w-full">
          <div class="md:w-1/2 md:pr-10 md:py-6">
            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 1
                </h2>
                <p class="leading-relaxed">fine asset information.</p>
              </div>
            </div>
            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 2
                </h2>
                <p class="leading-relaxed">easy browsing and comparison within the expiration dates</p>
              </div>
            </div>
            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 3
                </h2>
                <p class="leading-relaxed">grouping by same strike</p>
              </div>
            </div>
            <div class="relative flex pb-12">
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 4
                </h2>
                <p class="leading-relaxed">filtering and aggregate data across markets and providers
                </p>
              </div>
            </div>
          </div>
          <div class="md:w-1/2 md:pl-10 md:py-6">
            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div class="relative flex">
                <div
                  class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                  <i-bi-check-all />
                </div>
                <div class="flex-grow pl-4">
                  <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 5
                  </h2>
                  <p class="leading-relaxed">computing of indicators based on bid/ask orders</p>
                </div>
              </div>
            </div>

            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 6
                </h2>
                <p class="leading-relaxed">Mark IV (quoted bid and ask volatilities).</p>
              </div>
            </div>
            <div class="relative flex pb-12">
              <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
              </div>
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">STEP 7
                </h2>
                <p class="leading-relaxed">the <a class="text-green-500"
                    href="https://en.wikipedia.org/wiki/Greeks_(finance)">Greeks</a></p>
              </div>
            </div>
            <div class="relative flex pb-12">
              <div
                class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-indigo-800">
                <i-bi-check-all />
              </div>
              <div class="flex-grow pl-4">
                <h2 class="mb-1 text-sm font-medium tracking-wider">Final
                </h2>
                <p class="leading-relaxed">strategy outcome:(1-buy an option at the lowest price.
                  2-sell an option at the highest price. 3-arbitrage between two markets)</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Team section -->
    <section>
      <div class="container px-5 py-24 mx-auto">
        <div class="flex flex-col w-full mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium">OUR TEAM</h1>
          <p class="mx-auto leading-relaxed lg:w-2/3">The five members of the international Rainbow
            team:</p>
        </div>
        <div class="flex flex-wrap justify-center m-4">
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="team" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/o.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">0uep</h2>
                <p class="mb-4">0uep is a senior Dev, SysAdmin, and team leader with 20+ years XP.
                </p>
                <span class="inline-flex">
                  <a class="text-gray-500 hover:text-foreground" href="https://github.com/0uep">
                    <i-akar-icons-github-fill />
                  </a>
                </span>
              </div>
            </div>
          </div>
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="team" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/syn.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Synwx</h2>
                <p class="mb-4">Syn is a senior full-stack dev since the 90’s.</p>
                <span class="inline-flex">
                  <a class="text-gray-500 hover:text-foreground" href="https://github.com/synw">
                    <i-akar-icons-github-fill />
                  </a>
                </span>
              </div>
            </div>
          </div>
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="team" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/kahlys.jpeg">
              <div class="w-full">
                <h2 class="text-lg font-medium">kahlys</h2>
                <p class="mb-4">Kahlys is a very smart & pragmatic coder knowing many techs(Go, Rust,
                  Docker, web).</p>
                <span class="inline-flex">
                  <a class="text-gray-500 hover:text-foreground" href="https://github.com/kahlys">
                    <i-akar-icons-github-fill />
                  </a>
                </span>
              </div>
            </div>
          </div>
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="team" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/jy.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Cryptohazard</h2>
                <p class="mb-4">Cryptohazard is a blockchain researcher contributing to the success
                  of
                  the DeFi.</p>
                <span class="inline-flex">
                  <a class="text-gray-500 hover:text-foreground" href="https://github.com/cryptohazard">
                    <i-akar-icons-github-fill />
                  </a>
                </span>
              </div>
            </div>
          </div>
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="team" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/arad.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Arad</h2>
                <p class="mb-4">Arad is a talented junior front-end developer and UX specialist
                  living
                  in Iran.</p>
                <span class="inline-flex">
                  <a class="text-gray-500 hover:text-foreground" href="https://github.com/arad1990">
                    <i-akar-icons-github-fill />
                  </a>
                </span>
              </div>
            </div>
          </div>

        </div>
      </div>
    </section>
    <!-- Other tools section -->
    <section>
      <div class="container px-5 py-24 mx-auto">
        <div class="mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium sm:text-3xl">Our goals and
            tools</h1>
          <p class="mx-auto leading-relaxed xl:w-2/4 lg:w-3/4">in
            <a class="text-green-500" href="https://github.com/teal-finance/">Teal.Finance organization</a></p>
          <div class="flex justify-center mt-6">
            <div class="inline-flex w-16 h-1 rounded-full bg-primary"></div>
          </div>
        </div>
        <div class="flex flex-wrap -mx-4 -mt-4 -mb-10 space-y-6 sm:-m-4 md:space-y-0">
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full dark:bg-indigo-800 bg-primary">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">Teal.Finance</h2>
              <p class="leading-relaxed">We provide cryptocurrencies market data for trading. We also provide
                blockchain data and the tools to explore blockchains..</p>
              <a href="https://teal.finance" class="inline-flex items-center mt-3 text-green-500">Learn More
                <i-akar-icons-arrow-right />
              </a>
            </div>
          </div>
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full bg-primary dark:bg-indigo-800">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">FTX-NFT</h2>
              <p class="leading-relaxed">We provide cryptocurrencies market data for trading. We also provide
                blockchain data and the tools to explore blockchains..</p>
              <a href="https://teal.finance/nft" class="inline-flex items-center mt-3 text-green-500">Learn More
                <i-akar-icons-arrow-right />
              </a>
            </div>
          </div>
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full bg-primary dark:bg-indigo-800">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">Teal.Finance</h2>
              <p class="leading-relaxed">We provide cryptocurrencies market data for trading. We also provide
                blockchain data and the tools to explore blockchains..</p>
              <a href="https://teal.finance" class="inline-flex items-center mt-3 text-green-500">Learn More
                <i-akar-icons-arrow-right />
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {

}
</script>

<style lang="sass" scoped>
.text-rainbow
    font-size: 70px
    line-height: 70px
    background: linear-gradient(to right, #ee82ee, #ee82ee, #0000ff, #008000, #ffff00, #ffa500, #ff0000)
    -webkit-background-clip: text
    -webkit-text-fill-color: transparent
</style>