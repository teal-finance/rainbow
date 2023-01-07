<template>
  <div>
    <section class="text-foreground dark:text-foreground-dark">
      <div class="container flex flex-col items-center px-5 mx-auto py-14 md:flex-row">
        <div
          class="flex flex-col items-center text-center lg:flex-grow md:w-1/2 lg:pl-24 md:pl-16 md:items-start md:text-left">
          <h1 class="mb-4 font-bold text-9xl sm:text-4xl text-rainbow">Rainbow</h1>
          <p class="mb-8 leading-8 sm:w-full">
            Dashboard for DeFi options trading highlighting market opportunities
            to ease strategies decision. With Rainbow, users can scan CEX/DEX markets
            and fetch available options data.
            <br>
            This project is sponsored by
            <a class="text-green-500 hover:text-primary" href="https://teal.finance">Teal.Finance</a>,
            originally initiated to participate to the
            <a class="text-green-500 hover:text-primary" href="https://showcase.ethglobal.com/ethonline2021/rainbow">
              Ethereum's EthGlobal Hackathon
            </a>, the Solana's Ignition Hackathon and the
            <a class="text-green-500 hover:text-primary" href="https://www.encode.club/hack-defi">Encode Club's Hack DeFi</a>.
            <br>
          </p>
        </div>
        <div class="mb-10 md:mb-0">
          <img class="object-cover object-center rounded" alt="Rainbow dashboard" src="@/assets/dashboard.svg">
        </div>
      </div>
    </section>
    <!-- Inspired section -->
    <section>
      <div class="container flex flex-wrap px-5 py-24 mx-auto">
        <div class="flex flex-col w-full mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium">Rainbow was inspired by the following trends</h1>
          <div class="flex flex-wrap w-full">
            <div class="pr-6 my-10 md:w-1/2 md:py-6">
            <p class="p-6 leading-relaxed text-left shadow-md">
              The crypto derivatives markets are expected to grow a lot more in the upcoming years.
              Specifically, the options markets will see the biggest growth
              because in Finance, Option market is much bigger than the underlying spot market.
              This is lagging in Crypto, when we look at Deribit's volume,
              the main (centralized) venue for Crypto Options,
              compared to Perpetuals future.
              <a class="text-green-500 hover:text-primary" href="https://cryptonuage.com/posts/internet-of-decentralized-options/">
               Read more...
              </a>
            </p>
            </div>
            <div class="my-10 md:w-1/2 md:py-6">
              <div class="relative flex pb-5">
                <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                  <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
                </div>
                <div
                  class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-secondary">
                  <i-carbon:number-1 class="text-3xl"/>
                </div>
                <div class="flex-grow pl-4">
                  <p class="leading-relaxed text-left">Crypto-assets and DeFi becoming mainstream.</p>
                </div>
              </div>
              <div class="relative flex pb-5">
                <div class="absolute inset-0 flex items-center justify-center w-10 h-full">
                  <div class="w-1 h-full bg-gray-200 pointer-events-none"></div>
                </div>
                <div
                  class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-secondary">
                  <i-carbon:number-2 class="text-3xl"/>
                </div>
                <div class="flex-grow pl-4">
                  <p class="leading-relaxed text-left">
                    Advent of the Internet of Blockchains’ world with more
                    cross-chain applications and communications.
                  </p>
                </div>
              </div>
              <div class="relative flex pb-5">
                <div
                  class="relative z-10 inline-flex items-center justify-center flex-shrink-0 w-10 h-10 text-white rounded-full bg-primary dark:bg-secondary">
                  <i-carbon:number-3 class="text-3xl"/>
                </div>
                <div class="flex-grow pl-4">
                  <p class="leading-relaxed text-left">
                    Crypto options trading growing (hopefully) bigger than the spot
                    market, like in traditional Finance.
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section id="support">
      <div class="container flex flex-wrap px-5 py-24 mx-auto">
        <div class="flex flex-col w-full mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium">First providers implemented</h1>
          <div class="flex flex-wrap w-full">
            <div class="p-4 md:w-1/3">
              <div class="flex flex-col h-full p-8 bg-gray-100 rounded-lg dark:bg-primary-dark">
                <div class="flex items-center mb-3">
                  <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 mr-3 rounded-full">
                    <img src="@/assets/deribit.webp" alt="Deribit">
                  </div>
                  <h2 class="text-lg font-medium text-primary-dark dark:text-primary-r title-font">Deribit</h2>
                </div>
                <div class="flex-grow">
                  <p class="text-base leading-relaxed text-left">
                    Centralized exchange. Main venue for crypto Options trading.
                  </p>
                  <a href="https://www.deribit.com/" class="inline-flex items-center mt-3 text-green-500 hover:text-primary">
                    Visit Deribit
                  </a>
                </div>
              </div>
            </div>
            <div class="p-4 md:w-1/3">
              <div class="flex flex-col h-full p-8 bg-gray-100 rounded-lg dark:bg-primary-dark">
                <div class="flex items-center mb-3">
                  <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 mr-3 rounded-full">
                    <img src="@/assets/Opyn.png" alt="Opyn (Ethereum)">
                  </div>
                  <h2 class="text-lg font-medium text-primary-dark dark:text-primary-r title-font">
                    Opyn (Ethereum)
                  </h2>
                </div>
                <div class="flex-grow">
                  <p class="text-base leading-relaxed text-left">
                    DeFi protocol using TheGraph &amp; 0x.
                  </p>
                  <a href="https://www.opyn.co/" class="inline-flex items-center mt-3 text-green-500 hover:text-primary">
                    Visit Opyn
                  </a>
                </div>
              </div>
            </div>
            <div class="p-4 md:w-1/3">
              <div class="flex flex-col h-full p-8 bg-gray-100 rounded-lg dark:bg-primary-dark">
                <div class="flex items-center mb-3">
                  <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 mr-3 rounded-full">
                    <img src="@/assets/PsyOptions.png" alt="PsyOptions (Solana)">
                  </div>
                  <h2 class="text-lg font-medium text-primary-dark dark:text-primary-r title-font">PsyOptions (Solana)</h2>
                </div>
                <div class="flex-grow">
                  <p class="text-base leading-relaxed text-left">
                    DeFi protocol build on Serum order books.
                  </p>
                  <a href="https://www.psyoptions.io/" class="inline-flex items-center mt-3 text-green-500 hover:text-primary">
                    Visit PsyOptions
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
    <!-- Team section -->
    <section id="team">
      <div class="container px-5 py-24 mx-auto">
        <div class="flex flex-col w-full mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium sm:text-3xl">Our Team</h1>
          <p class="mx-auto leading-relaxed lg:w-2/3">
            The members of the international Rainbow team:
          </p>
        </div>
        <div class="flex flex-wrap justify-center m-4">
          <div class="p-4 lg:w-1/5 md:w-1/2">
            <div class="flex flex-col items-center h-full text-center">
              <img alt="Rainbow team member" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
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
              <img alt="Rainbow team member" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/syn.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Syn</h2>
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
              <img alt="Rainbow team member" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/kahlys.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Kahlys</h2>
                <p class="mb-4">Kahlys is a very smart &amp; pragmatic coder knowing many techs(Go, Rust,
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
              <img alt="Rainbow team member" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/jy.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Cryptohazard</h2>
                <p class="mb-4">
                  Cryptohazard is a blockchain PhD actively contributing to the success of the DeFi.
                </p>
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
              <img alt="Rainbow team member" class="flex-shrink-0 object-cover object-center w-full mb-4 rounded-lg h-58"
                src="@/assets/profile/arad.jpg">
              <div class="w-full">
                <h2 class="text-lg font-medium">Arad</h2>
                <p class="mb-4">
                  Arad is a talented junior frontend developer and UX/UI specialist.
                </p>
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
    <section id="projects">
      <div class="container px-5 py-24 mx-auto">
        <div class="mb-20 text-center">
          <h1 class="mb-4 text-2xl font-medium sm:text-3xl">Our projects</h1>
          <p class="mx-auto leading-relaxed xl:w-2/4 lg:w-3/4">in
            <a class="text-green-500" href="https://github.com/teal-finance/">Teal.Finance organization</a></p>
          <div class="flex justify-center mt-6">
            <div class="inline-flex w-16 h-1 rounded-full bg-primary"></div>
          </div>
        </div>
        <div class="flex flex-wrap -mx-4 -mt-4 -mb-10 space-y-6 sm:-m-4 md:space-y-0">
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full dark:bg-secondary bg-primary">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">Rainbow</h2>
              <p class="leading-relaxed">
                This is our flagship project started during hackathons.
                Rainbow enable traders to access in one place DeFi options data from various protocols.
                It provides a dashboard and an API to see live and historical options prices / best limits,
                currently from seven different exchanges:
                Deribit, PsyOptions, Opyn, Delta Exchange, Lyra, Thales, Zeta Markets and 0x.
              </p>
              <a href="https://teal.finance/rainbow" class="inline-flex items-center mt-3 text-green-500">See More
                <i-akar-icons-arrow-right />
              </a>
            </div>
          </div>
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full bg-primary dark:bg-secondary">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">Teal.NFT</h2>
              <p class="leading-relaxed">
                A fun side project where we explore FTX NFT data,
                even before they properly documented that API.
                We can now see which NFTs are the most popular on FTX & FTX US.
              </p>
              <a href="https://teal.finance/nft" class="inline-flex items-center mt-3 text-green-500">
                Learn More
                <i-akar-icons-arrow-right />
              </a>
            </div>
          </div>
          <div class="flex flex-col items-center p-4 text-center md:w-1/3">
            <div
              class="inline-flex items-center justify-center flex-shrink-0 w-20 h-20 mb-5 text-white rounded-full bg-primary dark:bg-secondary">
              <i-eos-icons-blockchain class="text-2xl" />
            </div>
            <div class="flex-grow">
              <h2 class="mb-3 text-lg font-medium">TealTicks</h2>
              <p class="leading-relaxed">
                TealTicks continuously collects market data across as much instruments as possible
                and provides a tick-by-tick oriented API with Jupiter support for traders.
              </p>
              <a href="https://teal.finance" class="inline-flex items-center mt-3 text-green-500">
                Learn More
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
    font-size: 4em
    line-height: 1em
    background: linear-gradient(to right, #ee82ee, #ee82ee, #0000ff, #008000, #ffff00, #ffa500, #ff0000)
    -webkit-background-clip: text
    -webkit-text-fill-color: transparent
</style>
