package main

var quotes = []*Quote{
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Don't cry because it's over, smile because it happened.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1511992603p2/8630.jpg",
		Content: "I'm selfish, impatient and a little insecure. I make mistakes, I am out of control and at times hard to handle. But if you can't handle me at my worst, then you sure as hell don't deserve me at my best.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Be yourself; everyone else is already taken.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Two things are infinite: the universe and human stupidity; and I'm not sure about the universe.",
	},
	{
		Author: "Frank Zappa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1315160559p2/22302.jpg",
		Content: "So many books, so little time.",
	},
	{
		Author: "Bernard M. Baruch",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331977133p2/5768330.jpg",
		Content: "Be who you are and say what you feel, because those who mind don't matter, and those who matter don't mind.",
	},
	{
		Author: "Marcus Tullius Cicero",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197881416p2/13755.jpg",
		Content: "A room without books is like a body without a soul.",
	},
	{
		Author: "William W. Purkey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282396130p2/1744830.jpg",
		Content: "You've gotta dance like there's nobody watching,Love like you'll never be hurt,Sing like there's nobody listening,And live like it's heaven on earth.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "You know you're in love when you can't fall asleep because reality is finally better than your dreams.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "You only live once, but if you do it right, once is enough.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Be the change that you wish to see in the world.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "In three words I can sum up everything I've learned about life: it goes on.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "If you want to know what a man's like, take a good look at how he treats his inferiors, not his equals.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Don’t walk in front of me… I may not followDon’t walk behind me… I may not leadWalk beside me… just be my friend",
	},
	{
		Author: "C.S. LewisThe Four Loves",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1387503187p2/10554.jpg",
		Content: "Friendship ... is born at the moment when one man says to another \"What! You too? I thought that no one but myself . . .",
	},
	{
		Author: "Eleanor RooseveltThis is My Story",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "No one can make you feel inferior without your consent.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "If you tell the truth, you don't have to remember anything.",
	},
	{
		Author: "Elbert Hubbard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216826209p2/114059.jpg",
		Content: "A friend is someone who knows all about you and still loves you.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "I've learned that people will forget what you said, people will forget what you did, but people will never forget how you made them feel.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Always forgive your enemies; nothing annoys them so much.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Live as if you were to die tomorrow. Learn as if you were to live forever.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "To live is the rarest thing in the world. Most people exist, that is all.",
	},
	{
		Author: "Martin Luther King Jr.A Testament of Hope: The Essential Writings and Speeches",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Darkness cannot drive out darkness: only light can do that. Hate cannot drive out hate: only love can do that.",
	},
	{
		Author: "Oscar WildeThe Happy Prince and Other Stories",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I am so clever that sometimes I don't understand a single word of what I am saying.",
	},
	{
		Author: "Friedrich NietzscheTwilight of the Idols",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "Without music, life would be a mistake.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "To be yourself in a world that is constantly trying to make you something else is the greatest accomplishment.",
	},
	{
		Author: "Rob Siltanen",
		AuthorAvatar: "",
		Content: "Here's to the crazy ones. The misfits. The rebels. The troublemakers. The round pegs in the square holes. The ones who see things differently. They're not fond of rules. And they have no respect for the status quo. You can quote them, disagree with them, glorify or vilify them. About the only thing you can't do is ignore them. Because they change things. They push the human race forward. And while some may see them as the crazy ones, we see genius. Because the people who are crazy enough to think they can change the world, are the ones who do.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1387506524p2/2534.jpg",
		Content: "We accept the love we think we deserve.",
	},
	{
		Author: "Narcotics Anonymous",
		AuthorAvatar: "",
		Content: "Insanity is doing the same thing, over and over again, but expecting different results.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1512016299p2/12379.jpg",
		Content: "I believe that everything happens for a reason. People change so that you can learn to let go, things go wrong so that you appreciate them when they're right, you believe lies so you eventually learn to trust no one but yourself, and sometimes good things fall apart so better things can fall together.",
	},
	{
		Author: "Steve Martin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1410018030p2/7103.jpg",
		Content: "A day without sunshine is like, you know, night.",
	},
	{
		Author: "William ShakespeareAll's Well That Ends Well",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Love all, trust a few, do wrong to none.",
	},
	{
		Author: "Neil GaimanThe Kindly Ones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Have you ever been in love? Horrible isn't it? It makes you so vulnerable. It opens your chest and it opens up your heart and it means that someone can get inside you and mess you up.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Never put off till tomorrow what may be done day after tomorrow just as well.",
	},
	{
		Author: "Kent M. KeithThe Silent Revolution: Dynamic Leadership in the Student Council",
		AuthorAvatar: "",
		Content: "\n  The Paradoxical Commandments\nPeople are illogical, unreasonable, and self-centered.Love them anyway.If you do good, people will accuse you of selfish ulterior motives.Do good anyway.If you are successful, you will win false friends and true enemies.Succeed anyway.The good you do today will be forgotten tomorrow.Do good anyway.Honesty and frankness make you vulnerable.Be honest and frank anyway.The biggest men and women with the biggest ideas can be shot down by the smallest men and women with the smallest minds.Think big anyway.People favor underdogs but follow only top dogs.Fight for a few underdogs anyway.What you spend years building may be destroyed overnight.Build anyway.People really need help but may attack you if you do help them.Help people anyway.Give the world the best you have and you'll get kicked in the teeth.Give the world the best you have anyway.",
	},
	{
		Author: "Dr. SeussOh, The Places You'll Go!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "You have brains in your head. You have feet in your shoes. You can steer yourself any direction you choose. You're on your own. And you know what you know. And YOU are the one who'll decide where to go...",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "That which does not kill us makes us stronger.",
	},
	{
		Author: "George R.R. MartinA Dance with Dragons",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "A reader lives a thousand lives before he dies, said Jojen. The man who never reads lives only one.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "For every minute you are angry you lose sixty seconds of happiness.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "If you judge people, you have no time to love them.",
	},
	{
		Author: "Lao Tzu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "Being deeply loved by someone gives you strength, while loving someone deeply gives you courage.",
	},
	{
		Author: "George Eliot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525078524p2/173.jpg",
		Content: "It is never too late to be what you might have been.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "This life is what you make it. No matter what, you're going to mess up sometimes, it's a universal truth. But the good part is you get to decide how you're going to mess it up. Girls will be your friends - they'll act like it anyway. But just remember, some come, some go. The ones that stay with you through everything - they're your true best friends. Don't let go of them. Also remember, sisters make the best friends in the world. As for lovers, well, they'll come and go too. And baby, I hate to say it, most of them - actually pretty much all of them are going to break your heart, but you can't give up because if you give up, you'll never find your soulmate. You'll never find that half who makes you whole and that goes for everything. Just because you fail once, doesn't mean you're gonna fail at everything. Keep trying, hold on, and always, always, always believe in yourself, because if you don't, then who will, sweetie? So keep your head high, keep your chin up, and most importantly, keep smiling, because life's a beautiful thing and there's so much to smile about.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It takes a great deal of bravery to stand up to our enemies, but just as much to stand up to our friends.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "There is never a time or place for true love. It happens accidentally, in a heartbeat, in a single flashing, throbbing moment.",
	},
	{
		Author: "Douglas AdamsThe Salmon of Doubt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "I love deadlines. I love the whooshing noise they make as they go by.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "My thoughts are stars I cannot fathom into constellations.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "If you can't explain it to a six year old, you don't understand it yourself.",
	},
	{
		Author: "Robert A. HeinleinStranger in a Strange Land",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192826560p2/205.jpg",
		Content: "Love is that condition in which the happiness of another person is essential to your own.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "I'm not upset that you lied to me, I'm upset that from now on I can't believe you.",
	},
	{
		Author: "Garrison Keillor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1259697704p2/2014.jpg",
		Content: "Anyone who thinks sitting in church can make you a Christian must also think that sitting in a garage can make you a car.",
	},
	{
		Author: "Maya AngelouI Know Why the Caged Bird Sings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "There is no greater agony than bearing an untold story inside you.",
	},
	{
		Author: "Jorge Luis Borges",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1392103489p2/7572.jpg",
		Content: "I have always imagined that Paradise will be a kind of library.",
	},
	{
		Author: "Robert A. Heinlein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192826560p2/205.jpg",
		Content: "Women and cats will do as they please, and men and dogs should relax and get used to the idea.",
	},
	{
		Author: "Pablo Picasso",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536109p2/3253.jpg",
		Content: "Everything you can imagine is real.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Sometimes the questions are complicated and the answers are simple.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You love me. Real or not real?\"I tell him, \"Real.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "We don't see things as they are, we see them as we are.",
	},
	{
		Author: "William NicholsonShadowlands",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1397581529p2/9590.jpg",
		Content: "We read to know we're not alone.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "You may not be her first, her last, or her only. She loved before she may love again. But if she loves you now, what else matters? She's not perfect—you aren't either, and the two of you may never be perfect together but if she can make you laugh, cause you to think twice, and admit to being human and making mistakes, hold onto her and give her the most you can. She may not be thinking about you every second of the day, but she will give you a part of her that she knows you can break—her heart. So don't hurt her, don't change her, don't analyze and don't expect more than she can give. Smile when she makes you happy, let her know when she makes you mad, and miss her when she's not there.",
	},
	{
		Author: "H. Jackson Brown Jr.P.S. I Love You",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288560924p2/33394.jpg",
		Content: "Twenty years from now you will be more disappointed by the things that you didn't do than by the ones you did do. So throw off the bowlines. Sail away from the safe harbor. Catch the trade winds in your sails. Explore. Dream. Discover.",
	},
	{
		Author: "André GideAutumn Leaves",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424119198p2/7617.jpg",
		Content: "It is better to be hated for what you are than to be loved for what you are not.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "Only once in your life, I truly believe, you find someone who can completely turn your world around. You tell them things that you’ve never shared with another soul and they absorb everything you say and actually want to hear more. You share hopes for the future, dreams that will never come true, goals that were never achieved and the many disappointments life has thrown at you. When something wonderful happens, you can’t wait to tell them about it, knowing they will share in your excitement. They are not embarrassed to cry with you when you are hurting or laugh with you when you make a fool of yourself. Never do they hurt your feelings or make you feel like you are not good enough, but rather they build you up and show you the things about yourself that make you special and even beautiful. There is never any pressure, jealousy or competition but only a quiet calmness when they are around. You can be yourself and not worry about what they will think of you because they love you for who you are. The things that seem insignificant to most people such as a note, song or walk become invaluable treasures kept safe in your heart to cherish forever. Memories of your childhood come back and are so clear and vivid it’s like being young again. Colours seem brighter and more brilliant. Laughter seems part of daily life where before it was infrequent or didn’t exist at all. A phone call or two during the day helps to get you through a long day’s work and always brings a smile to your face. In their presence, there’s no need for continuous conversation, but you find you’re quite content in just having them nearby. Things that never interested you before become fascinating because you know they are important to this person who is so special to you. You think of this person on every occasion and in everything you do. Simple things bring them to mind like a pale blue sky, gentle wind or even a storm cloud on the horizon. You open your heart knowing that there’s a chance it may be broken one day and in opening your heart, you experience a love and joy that you never dreamed possible. You find that being vulnerable is the only way to allow your heart to feel true pleasure that’s so real it scares you. You find strength in knowing you have a true friend and possibly a soul mate who will remain loyal to the end. Life seems completely different, exciting and worthwhile. Your only hope and security is in knowing that they are a part of your life.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "All that is gold does not glitter,Not all those who wander are lost;The old that is strong does not wither,Deep roots are not reached by the frost.From the ashes a fire shall be woken,A light from the shadows shall spring;Renewed shall be blade that was broken,The crownless again shall be king.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1374084212p2/12415.jpg",
		Content: "It is our choices, Harry, that show what we truly are, far more than our abilities.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Imperfection is beauty, madness is genius and it's better to be absolutely ridiculous than absolutely boring.",
	},
	{
		Author: "Jane AustenNorthanger Abbey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "The person, be it gentleman or lady, who has not pleasure in a good novel, must be intolerably stupid.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "There are only two ways to live your life. One is as though nothing is a miracle. The other is as though everything is a miracle.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It does not do to dwell on dreams and forget to live.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "As he read, I fell in love the way you fall asleep: slowly, and then all at once.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Good friends, good books, and a sleepy conscience: this is the ideal life.",
	},
	{
		Author: "Maurice Switzer",
		AuthorAvatar: "",
		Content: "It is better to remain silent at the risk of being thought a fool, than to talk and remove all doubt of it.",
	},
	{
		Author: "William ShakespeareAs You Like It",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "The fool doth think he is wise, but the wise man knows himself to be a fool.",
	},
	{
		Author: "Allen Saunders",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368887014p2/276029.jpg",
		Content: "Life is what happens to us while we are making other plans.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Whenever you find yourself on the side of the majority, it is time to reform (or pause and reflect).",
	},
	{
		Author: "Oscar WildeLady Windermere's Fan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "We are all in the gutter, but some of us are looking at the stars.",
	},
	{
		Author: "Bil Keane",
		AuthorAvatar: "https://images.gr-assets.com/authors/1320885123p2/3230608.jpg",
		Content: "Yesterday is history, tomorrow is a mystery, today is a gift of God, which is why we call it the present.",
	},
	{
		Author: "Neil GaimanCoraline",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Fairy tales are more than true: not because they tell us that dragons exist, but because they tell us that dragons can be beaten.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "A woman is like a tea bag; you never know how strong it is until it's in hot water.",
	},
	{
		Author: "Thomas A. Edison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400354079p2/3091287.jpg",
		Content: "I have not failed. I've just found 10,000 ways that won't work.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The man who does not read has no advantage over the man who cannot read.",
	},
	{
		Author: "Elie Wiesel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1255518412p2/1049.jpg",
		Content: "The opposite of love is not hate, it's indifference. The opposite of art is not ugliness, it's indifference. The opposite of faith is not heresy, it's indifference. And the opposite of life is not death, it's indifference.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "It is not a lack of love, but a lack of friendship that makes unhappy marriages.",
	},
	{
		Author: "Gordon A. Eadie",
		AuthorAvatar: "",
		Content: "If you don't stand for something you will fall for anything.",
	},
	{
		Author: "Douglas AdamsThe Long Dark Tea-Time of the Soul",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "I may not have gone where I intended to go, but I think I have ended up where I needed to be.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "I solemnly swear that I am up to no good.",
	},
	{
		Author: "Groucho MarxThe Essential Groucho: Writings For By And About Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "Outside of a dog, a book is man's best friend. Inside of a dog it's too dark to read.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "I am enough of an artist to draw freely upon my imagination. Imagination is more important than knowledge. Knowledge is limited. Imagination encircles the world.",
	},
	{
		Author: "Pablo Neruda100 Love Sonnets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "I love you without knowing how, or when, or from where. I love you simply, without problems or pride: I love you in this way because I do not know any other way of loving but this, in which there is no I or you, so intimate that your hand upon my chest is my hand, so intimate that when I fall asleep your eyes close.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "I like nonsense, it wakes up the brain cells. Fantasy is a necessary ingredient in living.",
	},
	{
		Author: "Charles M. Schulz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590750p2/209672.jpg",
		Content: "All you need is love. But a little chocolate now and then doesn't hurt.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Sometimes, you read a book and it fills you with this weird evangelical zeal, and you become convinced that the shattered world will never be put back together unless and until all living humans read the book.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "If you only read the books that everyone else is reading, you can only think what everyone else is thinking.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1376429298p2/13146.jpg",
		Content: "Never trust anyone who has not brought a book with them.",
	},
	{
		Author: "Jim Henson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208316138p2/4427.jpg",
		Content: "Beauty is in the eye of the beholder and it may be necessary from time to time to give a stupid or misinformed beholder a black eye.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "You can never get a cup of tea large enough or a book long enough to suit me.",
	},
	{
		Author: "Dr. SeussHappy Birthday to You!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Today you are You, that is truer than true. There is no one alive who is Youer than You.",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It is impossible to live without failing at something, unless you live so cautiously that you might as well not have lived at all - in which case, you fail by default.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Of course it is happening inside your head, Harry, but why on earth should that mean that it is not real?",
	},
	{
		Author: "Woody Allen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501608198p2/10356.jpg",
		Content: "I'm not afraid of death; I just don't want to be there when it happens.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Have you fallen in love with the wrong person yet?'Jace said, \"Unfortunately, Lady of the Haven, my one true love remains myself.\"...\"At least,\" she said, \"you don't have to worry about rejection, Jace Wayland.\"\"Not necessarily. I turn myself down occasionally, just to keep it interesting.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "To the well-organized mind, death is but the next great adventure.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "If you want your children to be intelligent, read them fairy tales. If you want them to be more intelligent, read them more fairy tales.",
	},
	{
		Author: "Paul   Terry",
		AuthorAvatar: "",
		Content: "Whenever I feel the need to exercise, I lie down until it goes away.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1488565330p2/6856.jpg",
		Content: "If one cannot enjoy reading a book over and over again, there is no use in reading it at all.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Logic will get you from A to Z; imagination will get you everywhere.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I'm in love with you,\" he said quietly.\"Augustus,\" I said.\"I am,\" he said. He was staring at me, and I could see the corners of his eyes crinkling. \"I'm in love with you, and I'm not in the business of denying myself the simple pleasure of saying true things. I'm in love with you, and I know that love is just a shout into the void, and that oblivion is inevitable, and that we're all doomed and that there will come a day when all our labor has been returned to dust, and I know the sun will swallow the only earth we'll ever have, and I am in love with you.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I am nothing special, of this I am sure. I am a common man with common thoughts and I've led a common life. There are no monuments dedicated to me and my name will soon be forgotten, but I've loved another with all my heart and soul, and to me, this has always been enough..",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You don't get to choose if you get hurt in this world...but you do have some say in who hurts you. I like my choices.",
	},
	{
		Author: "Dr. SeussI Can Read With My Eyes Shut!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "The more that you read, the more things you will know. The more that you learn, the more places you'll go.",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "Life isn't about finding yourself. Life is about creating yourself.",
	},
	{
		Author: "Markus ZusakI Am the Messenger",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Sometimes people are beautiful.Not in looks.Not in what they say.Just in what they are.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "Folks are usually about as happy as they make their minds up to be.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "One good thing about music, when it hits you, you feel no pain.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "The truth is, everyone is going to hurt you. You just got to find the ones worth suffering for.",
	},
	{
		Author: "William ShakespeareA Midsummer Night's Dream",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Love looks not with the eyes, but with the mind,And therefore is winged Cupid painted blind.",
	},
	{
		Author: "Joan PowersPooh's Little Instruction Book",
		AuthorAvatar: "",
		Content: "If you live to be a hundred, I want to live to be a hundred minus one day so I never have to live without you.",
	},
	{
		Author: "Theodore Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1359838481p2/188.jpg",
		Content: "Do what you can, with what you have, where you are.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Not all of us can do great things. But we can do small things with great love.",
	},
	{
		Author: "John Green",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Saying 'I notice you're a nerd' is like saying, 'Hey, I notice that you'd rather be intelligent than be stupid, that you'd rather be thoughtful than be vapid, that you believe that there are things that matter more than the arrest record of Lindsay Lohan. Why is that?' In fact, it seems to me that most contemporary insults are pretty lame. Even 'lame' is kind of lame. Saying 'You're lame' is like saying 'You walk with a limp.' Yeah, whatever, so does 50 Cent, and he's done all right for himself.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "There is no pretending,\" Jace said with absolute clarity. \"I love you, and I will love you until I die, and if there is life after that, I'll love you then.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "A wise girl kisses but doesn't love, listens but doesn't believe, and leaves before she is left.",
	},
	{
		Author: "A.A. MilneThe House at Pooh Corner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Piglet sidled up to Pooh from behind. \"Pooh!\" he whispered.\"Yes, Piglet?\"\"Nothing,\" said Piglet, taking Pooh's paw. \"I just wanted to be sure of you.",
	},
	{
		Author: "Bill WattersonThe Complete Calvin and Hobbes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "Reality continues to ruin my life.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "Things change. And friends leave. Life doesn't stop for anybody.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Not all those who wander are lost.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "The only way out of the labyrinth of suffering is to forgive.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I am good, but not an angel. I do sin, but I am not the devil. I am just a small girl in a big world trying to find someone to love.",
	},
	{
		Author: "Bessie Anderson StanleyMore Heart Throbs Volume Two in Prose and Verse Dear to the American People And by them contributed as a Supplement to the original $10,000 Prize Book HEART THROBS",
		AuthorAvatar: "",
		Content: "He has achieved success who has lived well, laughed often, and loved much;Who has enjoyed the trust of pure women, the respect of intelligent men and the love of little children;Who has filled his niche and accomplished his task;Who has never lacked appreciation of Earth's beauty or failed to express it;Who has left the world better than he found it,Whether an improved poppy, a perfect poem, or a rescued soul;Who has always looked for the best in others and given them the best he had;Whose life was an inspiration;Whose memory a benediction.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Do not pity the dead, Harry. Pity the living, and, above all those who live without love.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "The reason I talk to myself is because I’m the only one whose answers I accept.",
	},
	{
		Author: "Benjamin Franklin Wade",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445333361p2/4184579.jpg",
		Content: "Go to heaven for the climate and hell for the company.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "One must always be careful of books,\" said Tessa, \"and what is inside them, for words have the power to change us.",
	},
	{
		Author: "C.S. LewisThe Four Loves",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "To love at all is to be vulnerable. Love anything and your heart will be wrung and possibly broken. If you want to make sure of keeping it intact you must give it to no one, not even an animal. Wrap it carefully round with hobbies and little luxuries; avoid all entanglements. Lock it up safe in the casket or coffin of your selfishness. But in that casket, safe, dark, motionless, airless, it will change. It will not be broken; it will become unbreakable, impenetrable, irredeemable. To love is to be vulnerable.",
	},
	{
		Author: "Rosemarie Urquico",
		AuthorAvatar: "",
		Content: "You should date a girl who reads.Date a girl who reads. Date a girl who spends her money on books instead of clothes, who has problems with closet space because she has too many books. Date a girl who has a list of books she wants to read, who has had a library card since she was twelve.Find a girl who reads. You’ll know that she does because she will always have an unread book in her bag. She’s the one lovingly looking over the shelves in the bookstore, the one who quietly cries out when she has found the book she wants. You see that weird chick sniffing the pages of an old book in a secondhand book shop? That’s the reader. They can never resist smelling the pages, especially when they are yellow and worn.She’s the girl reading while waiting in that coffee shop down the street. If you take a peek at her mug, the non-dairy creamer is floating on top because she’s kind of engrossed already. Lost in a world of the author’s making. Sit down. She might give you a glare, as most girls who read do not like to be interrupted. Ask her if she likes the book.Buy her another cup of coffee.Let her know what you really think of Murakami. See if she got through the first chapter of Fellowship. Understand that if she says she understood James Joyce’s Ulysses she’s just saying that to sound intelligent. Ask her if she loves Alice or she would like to be Alice.It’s easy to date a girl who reads. Give her books for her birthday, for Christmas, for anniversaries. Give her the gift of words, in poetry and in song. Give her Neruda, Pound, Sexton, Cummings. Let her know that you understand that words are love. Understand that she knows the difference between books and reality but by god, she’s going to try to make her life a little like her favorite book. It will never be your fault if she does.She has to give it a shot somehow.Lie to her. If she understands syntax, she will understand your need to lie. Behind words are other things: motivation, value, nuance, dialogue. It will not be the end of the world.Fail her. Because a girl who reads knows that failure always leads up to the climax. Because girls who read understand that all things must come to end, but that you can always write a sequel. That you can begin again and again and still be the hero. That life is meant to have a villain or two.Why be frightened of everything that you are not? Girls who read understand that people, like characters, develop. Except in the Twilight series.If you find a girl who reads, keep her close. When you find her up at 2 AM clutching a book to her chest and weeping, make her a cup of tea and hold her. You may lose her for a couple of hours but she will always come back to you. She’ll talk as if the characters in the book are real, because for a while, they always are.You will propose on a hot air balloon. Or during a rock concert. Or very casually next time she’s sick. Over Skype.You will smile so hard you will wonder why your heart hasn’t burst and bled out all over your chest yet. You will write the story of your lives, have kids with strange names and even stranger tastes. She will introduce your children to the Cat in the Hat and Aslan, maybe in the same day. You will walk the winters of your old age together and she will recite Keats under her breath while you shake the snow off your boots.Date a girl who reads because you deserve it. You deserve a girl who can give you the most colorful life imaginable. If you can only give her monotony, and stale hours and half-baked proposals, then you’re better off alone. If you want the world and the worlds beyond it, date a girl who reads.Or better yet, date a girl who writes.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Dumbledore watched her fly away, and as her silvery glow faded he turned back to Snape, and his eyes were full of tears.\"After all this time?\"\"Always,\" said Snape.",
	},
	{
		Author: "Charles M. SchulzThe Complete Peanuts, Vol. 5: 1959-1960",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1376733794p2/1024.jpg",
		Content: "I love mankind ... it's people I can't stand!!",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You don’t forget the face of the person who was your last hope.",
	},
	{
		Author: "Jess C. ScottThe Intern",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374157564p2/2980674.jpg",
		Content: "When someone loves you, the way they talk about you is different. You feel safe and comfortable.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "There is nothing to writing. All you do is sit down at a typewriter and bleed.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "You may say I'm a dreamer, but I'm not the only one. I hope someday you'll join us. And the world will live as one.",
	},
	{
		Author: "W.C. Fields",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331712829p2/82951.jpg",
		Content: "I am free of all prejudice. I hate everyone equally. ",
	},
	{
		Author: "Toni Morrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494211316p2/3534.jpg",
		Content: "If there's a book that you want to read, but it hasn't been written yet, then you must write it.",
	},
	{
		Author: "Neil GaimanThe Kindly Ones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I've been making a list of the things they don't teach you at school. They don't teach you how to love somebody. They don't teach you how to be famous. They don't teach you how to be rich or how to be poor. They don't teach you how to walk away from someone you don't love any longer. They don't teach you how to know what's going on in someone else's mind. They don't teach you what to say to someone who's dying. They don't teach you anything worth knowing.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Remember, we're madly in love, so it's all right to kiss me anytime you feel like it.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "The marks humans leave are too often scars.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "′Classic′ - a book which people praise and don't read.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What a slut time is. She screws everybody.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Finish each day and be done with it. You have done what you could. Some blunders and absurdities no doubt crept in; forget them as soon as you can. Tomorrow is a new day. You shall begin it serenely and with too high a spirit to be encumbered with your old nonsense.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "I have never let my schooling interfere with my education.",
	},
	{
		Author: "Nicholas SparksA Walk to Remember",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Love is like the wind, you can't see it but you can feel it.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "I find television very educating. Every time somebody turns on the set, I go into the other room and read a book.",
	},
	{
		Author: "Shel Silverstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "Listen to the mustn'ts, child. Listen to the don'ts. Listen to the shouldn'ts, the impossibles, the won'ts. Listen to the never haves, then listen close to me... Anything can happen, child. Anything can be.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "There is no friend as loyal as a book.",
	},
	{
		Author: "Helen Keller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1266096039p2/7275.jpg",
		Content: "When one door of happiness closes, another opens; but often we look so long at the closed door that we do not see the one which has been opened for us.",
	},
	{
		Author: "Walter M. Miller Jr.A Canticle for Leibowitz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1339680458p2/6025722.jpg",
		Content: "You don’t have a soul, Doctor. You are a soul. You have a body, temporarily.",
	},
	{
		Author: "Alexandre Dumas fils",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457825404p2/3186713.jpg",
		Content: "The difference between genius and stupidity is: genius has its limits.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "People think a soul mate is your perfect fit, and that's what everyone wants. But a true soul mate is a mirror, the person who shows you everything that is holding you back, the person who brings you to your own attention so you can change your life. A true soul mate is probably the most important person you'll ever meet, because they tear down your walls and smack you awake. But to live with a soul mate forever? Nah. Too painful. Soul mates, they come into your life just to reveal another layer of yourself to you, and then leave. A soul mates purpose is to shake you up, tear apart your ego a little bit, show you your obstacles and addictions, break your heart open so new light can get in, make you so desperate and out of control that you have to transform your life, then introduce you to your spiritual master...",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It matters not what someone is born, but what they grow to be.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "If you can make a woman laugh, you can make her do anything.",
	},
	{
		Author: "Jodi PicoultMy Sister's Keeper",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "You don't love someone because they're perfect, you love them in spite of the fact that they're not.",
	},
	{
		Author: "Douglas AdamsThe Restaurant at the End of the Universe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "The story so far:In the beginning the Universe was created.This has made a lot of people very angry and been widely regarded as a bad move.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "I would venture to guess that Anon, who wrote so many poems without signing them, was often a woman.",
	},
	{
		Author: "Robert FulghumTrue Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198520878p2/19630.jpg",
		Content: "We’re all a little weird. And life is a little weird. And when we find someone whose weirdness is compatible with ours, we join up with them and fall into mutually satisfying weirdness—and call it love—true love.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "Success is not final, failure is not fatal: it is the courage to continue that counts.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Life is like riding a bicycle. To keep your balance, you must keep moving.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Some infinities are bigger than other infinities.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "So, this is my life. And I want you to know that I am both happy and sad and I'm still trying to figure out how that could be.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Love never dies a natural death. It dies because we don't know how to replenish its source. It dies of blindness and errors and betrayals. It dies of illness and wounds; it dies of weariness, of witherings, of tarnishings.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1381874709p2/25106.jpg",
		Content: "Do one thing every day that scares you.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Some people never go crazy. What truly horrible lives they must lead.",
	},
	{
		Author: "Jane AustenNorthanger Abbey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "There is nothing I would not do for those who are really my friends. I have no notion of loving people by halves, it is not my nature.",
	},
	{
		Author: "Marthe Troly-CurtinPhrynette Married",
		AuthorAvatar: "",
		Content: "Time you enjoy wasting is not wasted time.",
	},
	{
		Author: "Terry PratchettDiggers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "The trouble with having an open mind, of course, is that people will insist on coming along and trying to put things in it.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "The real lover is the man who can thrill you by kissing your forehead or smiling into your eyes or just staring into space.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "What really knocks me out is a book that, when you're all done reading it, you wish the author that wrote it was a terrific friend of yours and you could call him up on the phone whenever you felt like it. That doesn't happen much, though.",
	},
	{
		Author: "Stephenie MeyerEclipse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "He's like a drug for you, Bella.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "It is what you read when you don't have to that determines what you will be when you can't help it.",
	},
	{
		Author: "Alfred Tennyson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454788521p2/13638502.jpg",
		Content: "If I had a flower for every time I thought of you...I could walk through my garden forever.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "Who are you to judge the life I live?I know I'm not perfect-and I don't live to be-but before you start pointing fingers...make sure you hands are clean!",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "And, when you want something, all the universe conspires in helping you to achieve it.",
	},
	{
		Author: "John GreenAn Abundance of Katherines",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Books are the ultimate Dumpees: put them down and they’ll wait for you forever; pay attention to them and they always love you back.",
	},
	{
		Author: "Linda Grayson",
		AuthorAvatar: "",
		Content: "There is nothing better than a friend, unless it is a friend with chocolate.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I declare after all there is no enjoyment like reading! How much sooner one tires of any thing than of a book! -- When I have a house of my own, I shall be miserable if I have not an excellent library.",
	},
	{
		Author: "Ayn Rand",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "The question isn't who is going to let me; it's who is going to stop me.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "To die will be an awfully big adventure.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1394570348p2/806.jpg",
		Content: "It takes courage to grow up and become who you really are.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "I wish it need not have happened in my time,\" said Frodo.\"So do I,\" said Gandalf, \"and so do all who live to see such times. But that is not for them to decide. All we have to decide is what to do with the time that is given us.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Anyone who has never made a mistake has never tried anything new.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "You never really understand a person until you consider things from his point of view... Until you climb inside of his skin and walk around in it.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "We believe in ordinary acts of bravery, in the courage that drives one person to stand up for another.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "There are infinite numbers between 0 and 1. There's .1 and .12 and .112 and an infinite collection of others. Of course, there is a bigger infinite set of numbers between 0 and 2, or between 0 and a million. Some infinities are bigger than other infinities. A writer we used to like taught us that. There are days, many of them, when I resent the size of my unbounded set. I want more numbers than I'm likely to get, and God, I want more numbers for Augustus Waters than he got. But, Gus, my love, I cannot tell you how thankful I am for our little infinity. I wouldn't trade it for the world. You gave me a forever within the numbered days, and I'm grateful.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "A lady's imagination is very rapid; it jumps from admiration to love, from love to matrimony in a moment.",
	},
	{
		Author: "Marianne WilliamsonA Return to Love: Reflections on the Principles of \"A Course in Miracles\"",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198520601p2/17297.jpg",
		Content: "Our deepest fear is not that we are inadequate. Our deepest fear is that we are powerful beyond measure. It is our light, not our darkness that most frightens us. We ask ourselves, 'Who am I to be brilliant, gorgeous, talented, fabulous?' Actually, who are you not to be? You are a child of God. Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory of God that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "And in that moment, I swear we were infinite.",
	},
	{
		Author: "Jimi HendrixJimi Hendrix - Axis: Bold as Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207679011p2/7268.jpg",
		Content: "I'm the one that's got to die when it's time for me to die, so let me live my life the way I want to.",
	},
	{
		Author: "Madeleine L'Engle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1305256804p2/106.jpg",
		Content: "You have to write the book that wants to be written. And if the book will be too difficult for grown-ups, then you write it for children.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "So it's not gonna be easy. It's going to be really hard; we're gonna have to work at this everyday, but I want to do that because I want you. I want all of you, forever, everyday. You and me... everyday.",
	},
	{
		Author: "Khaled Hosseini",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "But better to get hurt by the truth than comforted with a lie.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Some day you will be old enough to start reading fairy tales again.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Remember, if the time should come when you have to make a choice between what is right and what is easy, remember what happened to a boy who was good, and kind, and brave, because he strayed across the path of Lord Voldemort. Remember Cedric Diggory.",
	},
	{
		Author: "Lewis Carroll",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "It’s no use going back to yesterday, because I was a different person then.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "The truth.\" Dumbledore sighed. \"It is a beautiful and terrible thing, and should therefore be treated with great caution.",
	},
	{
		Author: "Charles J. SykesDumbing Down Our Kids: Why American Children Feel Good About Themselves But Can't Read, Write or Add",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306363644p2/1962.jpg",
		Content: "Be nice to nerds. You may end up working for them. We all could.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "The books that the world calls immoral are books that show the world its own shame.",
	},
	{
		Author: "Kurt VonnegutMother Night",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "We are what we pretend to be, so we must be careful about what we pretend to be.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "... a mind needs books as a sword needs a whetstone, if it is to keep its edge.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I believe in Christianity as I believe that the sun has risen: not only because I see it, but because by it I see everything else.",
	},
	{
		Author: "Oscar WildeThe Importance of Being Earnest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "The truth is rarely pure and never simple.",
	},
	{
		Author: "Nicholas SparksAt First Sight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Just when you think it can't get any worse, it can. And just when you think it can't get any better, it can.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Numbing the pain for a while will make it worse when you finally feel it.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "The one you love and the one who loves you are never, ever the same person.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "I became insane, with long intervals of horrible sanity.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "If you're gonna be two-faced at least make one of them pretty.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You can't live your life for other people. You've got to do what's right for you, even if it hurts some people you love.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "An eye for an eye will only make the whole world blind.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Some people don't understand the promises they're making when they make them,\" I said.\"Right, of course. But you keep the promise anyway. That's what love is. Love is keeping the promise anyway.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "So I walked back to my room and collapsed on the bottom bunk, thinking that if people were rain, I was drizzle and she was a hurricane.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Is this the part where you start tearing off strips of your shirt to bind my wounds?\"\"If you wanted me to rip my clothes off, you should have just asked.",
	},
	{
		Author: "Roald Dahl",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1368665403p2/174690.jpg",
		Content: "Those who don't believe in magic will never find it.",
	},
	{
		Author: "Socrates",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390145726p2/275648.jpg",
		Content: "The only true wisdom is in knowing you know nothing.",
	},
	{
		Author: "Helen Keller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1266096039p2/7275.jpg",
		Content: "I would rather walk with a friend in the dark, than alone in the light.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "You can never be overdressed or overeducated.",
	},
	{
		Author: "Charles William Eliot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1312422304p2/4398096.jpg",
		Content: "Books are the quietest and most constant of friends; they are the most accessible and wisest of counselors, and the most patient of teachers.",
	},
	{
		Author: "C.E.M. Joad",
		AuthorAvatar: "https://images.gr-assets.com/authors/1345505812p2/5991490.jpg",
		Content: "Creativity is knowing how to hide your sources",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "I speak to everyone in the same way, whether he is the garbage man or the president of the university.",
	},
	{
		Author: "Joe KlaasTwelve Steps to Happiness",
		AuthorAvatar: "",
		Content: "The truth will set you free, but first it will piss you off.",
	},
	{
		Author: "Sarah DessenSomeone Like You",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Life is an awful, ugly place to not have a best friend.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I wish I could freeze this moment, right here, right now and live in it forever.",
	},
	{
		Author: "Rita Mae Brown",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209493600p2/23511.jpg",
		Content: "The statistics on sanity are that one out of every four people is suffering from a mental illness. Look at your 3 best friends. If they're ok, then it's you.",
	},
	{
		Author: "William StyronConversations with William Styron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1490146642p2/7565.jpg",
		Content: "A great book should leave you with many experiences, and slightly exhausted at the end. You live several lives while reading.",
	},
	{
		Author: "Lemony Snicket",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Fate is like a strange, unpopular restaurant filled with odd little waiters who bring you things you never asked for and don't always like.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "Until I feared I would lose it, I never loved to read. One does not love breathing.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I don't want to be a man,\" said Jace. \"I want to be an angst-ridden teenager who can't confront his own inner demons and takes it out verbally on other people instead.\"\"Well,\" said Luke, \"you're doing a fantastic job.",
	},
	{
		Author: "Joseph ConradChance",
		AuthorAvatar: "https://images.gr-assets.com/authors/1403814208p2/3345.jpg",
		Content: "Being a woman is a terribly difficult trade since it consists principally of dealings with men.",
	},
	{
		Author: "Victor Hugo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "Music expresses that which cannot be put into words and that which cannot remain silent",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "Perhaps one did not want to be loved so much as to be understood.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Keep away from people who try to belittle your ambitions. Small people always do that, but the really great make you feel that you, too, can become great.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Happiness is when what you think, what you say, and what you do are in harmony.",
	},
	{
		Author: "Laurel Thatcher UlrichWell-Behaved Women Seldom Make History",
		AuthorAvatar: "https://images.gr-assets.com/authors/1225603389p2/9639.jpg",
		Content: "Well-behaved women seldom make history.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The fear of death follows from the fear of life. A man who lives fully is prepared to die at any time.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "A lie can travel half way around the world while the truth is putting on its shoes.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Never tell the truth to people who are not worthy of it.",
	},
	{
		Author: "Alfred TennysonIn Memoriam",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454788521p2/13638502.jpg",
		Content: "Tis better to have loved and lostThan never to have loved at all.",
	},
	{
		Author: "Stephen KingOn Writing: A Memoir of the Craft",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Books are a uniquely portable magic.",
	},
	{
		Author: "William ShakespeareTwelfth Night",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Be not afraid of greatness. Some are born great, some achieve greatness, and others have greatness thrust upon them.",
	},
	{
		Author: "Chuck PalahniukDiary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "It's so hard to forget pain, but it's even harder to remember sweetness. We have no scar to show for happiness. We learn so little from peace.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Do you remember me telling you we are practicing non-verbal spells, Potter?\"\"Yes,\" said Harry stiffly.\"Yes, sir.\"\"There's no need to call me \"sir\" Professor.\"The words had escaped him before he knew what he was saying.",
	},
	{
		Author: "Haruki MurakamiKafka on the Shore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Memories warm you up from the inside. But they also tear you apart.",
	},
	{
		Author: "Maya AngelouWouldn't Take Nothing for My Journey Now",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "What you're supposed to do when you don't like a thing is change it. If you can't change it, change the way you think about it. Don't complain.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I can never read all the books I want; I can never be all the people I want and live all the lives I want. I can never train myself in all the skills I want. And why do I want? I want to live and feel all the shades, tones and variations of mental and physical experience possible in my life. And I am horribly limited.",
	},
	{
		Author: "Dr. SeussHorton Hears a Who!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "A person's a person, no matter how small.",
	},
	{
		Author: "Nicolas Chamfort",
		AuthorAvatar: "https://images.gr-assets.com/authors/1420504765p2/308646.jpg",
		Content: "A day without laughter is a day wasted.",
	},
	{
		Author: "Aristotle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390143800p2/2192.jpg",
		Content: "Knowing yourself is the beginning of all wisdom.",
	},
	{
		Author: "Pablo Neruda100 Love Sonnets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "I love you as certain dark things are to be loved, in secret, between the shadow and the soul.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "When adults say, \"Teenagers think they are invincible\" with that sly, stupid smile on their faces, they don't know how right they are. We need never be hopeless, because we can never be irreparably broken. We think that we are invincible because we are. We cannot be born, and we cannot die. Like all energy, we can only change shapes and sizes and manifestations. They forget that when they get old. They get scared of losing and failing. But that part of us greater than the sum of our parts cannot begin and cannot end, and so it cannot fail.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "It's the possibility of having a dream come true that makes life interesting.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I finally understood what true love meant...love meant that you care for another person's happiness more than your own, no matter how painful the choices you face might be.",
	},
	{
		Author: "Ernest HemingwayThe Garden of Eden",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "Happiness in intelligent people is the rarest thing I know.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Oh, I wouldn't mind, Hazel Grace. It would be a privilege to have my heart broken by you.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "He’s not perfect. You aren’t either, and the two of you will never be perfect. But if he can make you laugh at least once, causes you to think twice, and if he admits to being human and making mistakes, hold onto him and give him the most you can. He isn’t going to quote poetry, he’s not thinking about you every moment, but he will give you a part of him that he knows you could break. Don’t hurt him, don’t change him, and don’t expect for more than he can give. Don’t analyze. Smile when he makes you happy, yell when he makes you mad, and miss him when he’s not there. Love hard when there is love to be had. Because perfect guys don’t exist, but there’s always one guy that is perfect for you.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Thomas Edison's last words were \"It's very beautiful over there\". I don't know where there is, but I believe it's somewhere, and I hope it's beautiful.",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Jesus!\" Luke exclaimed.\"Actually, it's just me,\" said Simon. \"Although I've been told the resemblance is startling.",
	},
	{
		Author: "Markus Herz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424985471p2/5249608.jpg",
		Content: "Be careful about reading health books. Some fine day you'll die of a misprint.",
	},
	{
		Author: "Ralph Waldo EmersonEmerson in His Journals",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "It is one of the blessings of old friends that you can afford to be stupid with them.",
	},
	{
		Author: "Oscar WildeThe Critic as Artist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Yes: I am a dreamer. For a dreamer is one who can only find his way by moonlight, and his punishment is that he sees the dawn before the rest of the world.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Just because you have the emotional range of a teaspoon doesn't mean we all have.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Only the very weak-minded refuse to be influenced by literature and poetry.",
	},
	{
		Author: "Terry JohnsonInsignificance",
		AuthorAvatar: "",
		Content: "Have you ever noticed how ‘What the hell’ is always the right decision to make?",
	},
	{
		Author: "Henri J.M. NouwenOut of Solitude: Three Meditations on the Christian Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198543648p2/4837.jpg",
		Content: "When we honestly ask ourselves which person in our lives mean the most to us, we often find that it is those who, instead of giving advice, solutions, or cures, have chosen rather to share our pain and touch our wounds with a warm and tender hand. The friend who can be silent with us in a moment of despair or confusion, who can stay with us in an hour of grief and bereavement, who can tolerate not knowing, not curing, not healing and face with us the reality of our powerlessness, that is a friend who cares.",
	},
	{
		Author: "Frederick Lewis Donaldson",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1435957881p2/32234.jpg",
		Content: "The Seven Social Sins are: Wealth without work. Pleasure without conscience. Knowledge without character. Commerce without morality. Science without humanity. Worship without sacrifice. Politics without principle.From a sermon given by Frederick Lewis Donaldson in Westminster Abbey, London, on March 20, 1925.",
	},
	{
		Author: "Marlene Dietrich",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215577909p2/459689.jpg",
		Content: "It's the friends you can call up at 4 a.m. that matter.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "If you don't have time to read, you don't have the time (or the tools) to write. Simple as that.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "A children's story that can only be enjoyed by children is not a good children's story in the slightest.",
	},
	{
		Author: "Charles BukowskiBarfly",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Do you hate people?",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "The world is not a wish-granting factory.",
	},
	{
		Author: "Anne FrankAnne Frank's Tales from the Secret Annex: A Collection of Her Short Stories, Fables, and Lesser-Known Writings, Revised Edition",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "How wonderful it is that nobody need wait a single moment before starting to improve the world.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "We're all human, aren't we? Every human life is worth the same, and worth saving.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Sometimes you have to be apart from people you love, but that doesn't make you love them any less. Sometimes you love them more.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "It's far better to be unhappy alone than unhappy with someone — so far.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Never forget what you are, for surely the world will not. Make it your strength. Then it can never be your weakness. Armour yourself in it, and it will never be used to hurt you.",
	},
	{
		Author: "Lisa KleypasBlue-Eyed Devil",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288899037p2/27847.jpg",
		Content: "I no longer believed in the idea of soul mates, or love at first sight. But I was beginning to believe that a very few times in your life, if you were lucky, you might meet someone who was exactly right for you. Not because he was perfect, or because you were, but because your combined flaws were arranged in a way that allowed two separate beings to hinge together.",
	},
	{
		Author: "Isaac Asimov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341965730p2/16667.jpg",
		Content: "The saddest aspect of life right now is that science gathers knowledge faster than society gathers wisdom.",
	},
	{
		Author: "Terry PratchettJingo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Give a man a fire and he's warm for a day, but set fire to him and he's warm for the rest of his life.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Becoming fearless isn't the point. That's impossible. It's learning how to control your fear, and how to be free from it.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "You will never be happy if you continue to search for what happiness consists of. You will never live if you are looking for the meaning of life.",
	},
	{
		Author: "George Harrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340857233p2/20108.jpg",
		Content: "If you don't know where you're going, any road'll take you there",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1429098233p2/170321.jpg",
		Content: "I am not pretty. I am not beautiful. I am as radiant as the sun.",
	},
	{
		Author: "John Green",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You can love someone so much...But you can never love people as much as you can miss them.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "I am no bird; and no net ensnares me: I am a free human being with an independent will.",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Love is needing someone. Love is putting up with someone's bad qualities because they somehow complete you.",
	},
	{
		Author: "Roald Dahl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1311554908p2/4273.jpg",
		Content: "And above all, watch with glittering eyes the whole world around you because the greatest secrets are always hidden in the most unlikely places. Those who don't believe in magic will never find it.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "There's nothing like deep breaths after laughing that hard. Nothing in the world like a sore stomach for the right reasons.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "The Road Not TakenTwo roads diverged in a yellow wood,And sorry I could not travel both\tAnd be one traveler, long I stood\tAnd looked down one as far as I could\tTo where it bent in the undergrowth;\t Then took the other, as just as fair,\tAnd having perhaps the better claim,\tBecause it was grassy and wanted wear;\tThough as for that the passing there\tHad worn them really about the same,\t And both that morning equally lay\tIn leaves no step had trodden black.\tOh, I kept the first for another day!\tYet knowing how way leads on to way,\tI doubted if I should ever come back.\t I shall be telling this with a sigh\tSomewhere ages and ages hence:\tTwo roads diverged in a wood, and I—\tI took the one less traveled by,\tAnd that has made all the difference.",
	},
	{
		Author: "H.N. TurteltaubThe Sacred Land",
		AuthorAvatar: "",
		Content: "If you were half as funny as you think you are, you'd be twice as funny as you really are.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "Nothing is impossible, the word itself says 'I'm possible'!",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "All God does is watch us and kill us when we get boring. We must never, ever be boring.",
	},
	{
		Author: "John GreenAn Abundance of Katherines",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What is the point of being alive if you don't at least try to do something remarkable?",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "One is loved because one is loved. No reason is needed for loving.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "When I despair, I remember that all through history the way of truth and love have always won. There have been tyrants and murderers, and for a time, they can seem invincible, but in the end, they always fall. Think of it--always.",
	},
	{
		Author: "Langston Hughes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206669966p2/36910.jpg",
		Content: "Hold fast to dreams,For if dreams dieLife is a broken-winged bird,That cannot fly.",
	},
	{
		Author: "Stephenie MeyerTwilight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "I like the night. Without the dark, we'd never see the stars.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Never memorize something that you can look up.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "And now I’m looking at you,",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Grief does not change you, Hazel. It reveals you.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "All we have to decide is what to do with the time that is given us.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "When you're in jail, a good friend will be trying to bail you out. A best friend will be in the cell next to you saying, 'Damn, that was fun'.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "In a good bookroom you feel in some mysterious way that you are absorbing the wisdom contained in all the books through your skin, without even opening them.",
	},
	{
		Author: "Kurt Vonnegut",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "Those who believe in telekinetics, raise my hand.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "When it comes down to it, I let them think what they want. If they care enough to bother with what I do, then I'm already better than them.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Never love anyone who treats you like you're ordinary.",
	},
	{
		Author: "Gilda Radner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212084245p2/145047.jpg",
		Content: "I wanted a perfect ending. Now I've learned, the hard way, that some poems don't rhyme, and some stories don't have a clear beginning, middle, and end. Life is about not knowing, having to change, taking the moment and making the best of it, without knowing what's going to happen next.Delicious Ambiguity.",
	},
	{
		Author: "Oscar Levant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1220982976p2/504377.jpg",
		Content: "There's a fine line between genius and insanity. I have erased this line.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I love you. I am who I am because of you. You are every reason, every hope, and every dream I've ever had, and no matter what happens to us in the future, everyday we are together is the greatest day of my life. I will always be yours. ",
	},
	{
		Author: "Terry PratchettThief of Time",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Some humans would do anything to see if it was possible to do it. If you put a large switch in some cave somewhere, with a sign on it saying 'End-of-the-World Switch. PLEASE DO NOT TOUCH', the paint wouldn't even have time to dry.",
	},
	{
		Author: "Richelle MeadFrostbite",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "You can't force love, I realized. It's there or it isn't. If it's not there, you've got to be able to admit it. If it is there, you've got to do whatever it takes to protect the ones you love.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "When you are courting a nice girl an hour seems like a second. When you sit on a red-hot cinder a second seems like an hour. That's relativity.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "Whatever you are, be a good one.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Death's got an Invisibility Cloak?\" Harry interrupted again.\"So he can sneak up on people,\" said Ron. \"Sometimes he gets bored of running at them, flapping his arms and shrieking...",
	},
	{
		Author: "Christian D. LarsonYour Forces and How to Use Them",
		AuthorAvatar: "",
		Content: "Promise YourselfTo be so strong that nothingcan disturb your peace of mind.To talk health, happiness, and prosperityto every person you meet.To make all your friends feelthat there is something in themTo look at the sunny side of everythingand make your optimism come true.To think only the best, to work only for the best,and to expect only the best.To be just as enthusiastic about the success of othersas you are about your own.To forget the mistakes of the pastand press on to the greater achievements of the future.To wear a cheerful countenance at all timesand give every living creature you meet a smile.To give so much time to the improvement of yourselfthat you have no time to criticize others.To be too large for worry, too noble for anger, too strong for fear,and too happy to permit the presence of trouble.To think well of yourself and to proclaim this fact to the world,not in loud words but great deeds.To live in faith that the whole world is on your sideso long as you are true to the best that is in you.",
	},
	{
		Author: "Lemony Snicket",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Wicked people never have time for reading. It's one of the reasons for their wickedness.",
	},
	{
		Author: "C.S. LewisThe Four Loves",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Friendship is unnecessary, like philosophy, like art.... It has no survival value; rather it is one of those things which give value to survival.",
	},
	{
		Author: "Thomas Szasz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357698047p2/11343.jpg",
		Content: "Two wrongs don't make a right, but they make a good excuse.",
	},
	{
		Author: "Nicholas Klein",
		AuthorAvatar: "",
		Content: "First they ignore you. Then they ridicule you. And then they attack you and want to burn you. And then they build monuments to you.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Where there is love there is life.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "I would always rather be happy than dignified.",
	},
	{
		Author: "Louise ErdrichThe Painted Drum",
		AuthorAvatar: "https://images.gr-assets.com/authors/1462224430p2/9388.jpg",
		Content: "Life will break you. Nobody can protect you from that, and living alone won't either, for solitude will also break you with its yearning. You have to love. You have to feel. It is the reason you are here on earth. You are here to risk your heart. You are here to be swallowed up. And when it happens that you are broken, or betrayed, or left, or hurt, or death brushes near, let yourself sit by an apple tree and listen to the apples falling all around you in heaps, wasting their sweetness. Tell yourself you tasted as many as you could.",
	},
	{
		Author: "Oscar WildeLord Arthur Savile's Crime and Other Stories",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Women are meant to be loved, not to be understood.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Well, don't expect us to be too impressed. We just saw Finnick Odair in his underwear.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I hope she'll be a fool -- that's the best thing a girl can be in this world, a beautiful little fool.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Peace begins with a smile..",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "May your coming year be filled with magic and dreams and good madness. I hope you read some fine books and kiss someone who thinks you're wonderful, and don't forget to make some art -- write or draw or build or sing or live as only you can. And I hope, somewhere in the next year, you surprise yourself.",
	},
	{
		Author: "Nicholas SparksThe Rescue",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You're going to come across people in your life who will say all the right words at all the right times. But in the end, it's always their actions you should judge them by. It's actions, not words, that matter.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "He can run faster than Severus Snape confronted with shampoo.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "Count your age by friends, not years. Count your life by smiles, not tears.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Mr. Moony presents his compliments to Professor Snape, and begs him to keep his abnormally large nose out of other people's business.Mr. Prongs agrees with Mr. Moony, and would like to add that Professor Snape is an ugly git.Mr. Padfoot would like to register his astonishment that an idiot like that ever became a professor.Mr. Wormtail bids Professor Snape good day, and advises him to wash his hair, the slimeball.",
	},
	{
		Author: "Mitch AlbomTuesdays with Morrie",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "Death ends a life, not a relationship.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "We are only as strong as we are united, as weak as we are divided.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Imagining the future is a kind of nostalgia. (...) You spend your whole life stuck in the labyrinth, thinking about how you'll escape it one day, and how awesome it will be, and imagining that future keeps you going, but you never do it. You just use the future to escape the present.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You are, and always have been, my dream.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "All that we see or seem is but a dream within a dream.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "I carry your heart with me (I carry it in my heart)I am never without it (anywhereI go you go,my dear; and whatever is done by only me is your doing,my darling)I fear no fate (for you are my fate,my sweet)I want no world (for beautiful you are my world,my true)and it's you are whatever a moon has always meant and whatever a sun will always sing is youhere is the deepest secret nobody knows(here is the root of the root and the bud of the bud and the sky of the sky of a tree called life; which growshigher than the soul can hope or mind can hide)and this is the wonder that's keeping the stars apartI carry your heart (I carry it in my heart)",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You are my best friend as well as my lover, and I do not know which side of you I enjoy the most. I treasure each side, just as I have treasured our life together.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Indifference and neglect often do much more damage than outright dislike.",
	},
	{
		Author: "Gustave Flaubert",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198541369p2/1461.jpg",
		Content: "Do not read, as children do, to amuse yourself, or like the ambitious, for the purpose of instruction. No, read in order to live.",
	},
	{
		Author: "Michael CunninghamThe Hours",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398891471p2/1432.jpg",
		Content: "You cannot find peace by avoiding life.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Is it true that you shouted at Professor Umbridge?\"\"Yes.\"\"You called her a liar?\"\"Yes.\"\"You told her He Who Must Not Be Named is back?\"\"Yes.\"\"Have a biscuit, Potter.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Don't think or judge, just listen.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "The Christian does not think God will love us because we are good, but that God will make us good because He loves us.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "So, I love you because the entire universe conspired to help me find you.",
	},
	{
		Author: "Sam LevensonIn One Era \u0026 Out the Other",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390755977p2/30932.jpg",
		Content: "For attractive lips, speak words of kindness.For lovely eyes, seek out the good in people.For a slim figure, share your food with the hungry.For beautiful hair, let a child run his fingers through it once a day.For poise, walk with the knowledge you’ll never walk alone....We leave you a tradition with a future.The tender loving care of human beings will never become obsolete.People even more than things have to be restored, renewed, revived, reclaimed and redeemed and redeemed and redeemed.Never throw out anybody.Remember, if you ever need a helping hand, you’ll find one at the end of your arm.As you grow older, you will discover that you have two hands: one for helping yourself, the other for helping others.Your “good old days",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "The only thing worse than a boy who hates you: a boy that loves you.",
	},
	{
		Author: "Margaret Mead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198589352p2/61107.jpg",
		Content: "Never doubt that a small group of thoughtful, committed, citizens can change the world. Indeed, it is the only thing that ever has.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "So, I guess we are who we are for alot of reasons. And maybe we'll never know most of them. But even if we don't have the power to choose where we come from, we can still choose where we go from there. We can still do things. And we can try to feel okay about them.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Don't touch any of my weapons without my permission.\"\"Well, there goes my plan for selling them all on eBay,\" Clary muttered.\"Selling them on what?\"Clary smiled blandly at him. \"A mythical place of great magical power.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "What lies behind us and what lies before us are tiny matters compared to what lies within us.",
	},
	{
		Author: "Carter Crocker",
		AuthorAvatar: "",
		Content: "Promise me you'll always remember: You're braver than you believe, and stronger than you seem, and smarter than you think.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Most people are other people. Their thoughts are someone else's opinions, their lives a mimicry, their passions a quotation.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Happy Hunger Games! And may the odds be ever in your favor.",
	},
	{
		Author: "Carter Crocker",
		AuthorAvatar: "",
		Content: "If ever there is tomorrow when we're not together... there is something you must always remember. You are braver than you believe, stronger than you seem, and smarter than you think. But the most important thing is, even if we're apart... I'll always be with you.",
	},
	{
		Author: "Rick RiordanThe Last Olympian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "With great power... comes great need to take a nap. Wake me up later.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "I don't know half of you half as well as I should like; and I like less than half of you half as well as you deserve.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "If we find ourselves with a desire that nothing in this world can satisfy, the most probable explanation is that we were made for another world.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What a treacherous thing to believe that a person is more than a person.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "When we love, we always strive to become better than we are. When we strive to become better than we are, everything around us becomes better too.",
	},
	{
		Author: "AnonymousHoly Bible: New International Version",
		AuthorAvatar: "",
		Content: "Love is patient, love is kind. It does not envy, it does not boast, it is not proud. It is not rude, it is not self-seeking, it is not easily angered, it keeps no record of wrongs. Love does not delight in evil but rejoices with the truth. It always protects, always trusts, always hopes, always perseveres.",
	},
	{
		Author: "Leo Tolstoy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "Everyone thinks of changing the world, but no one thinks of changing himself.",
	},
	{
		Author: "George V. HigginsThe Friends of Eddie Coyle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1413351408p2/19476.jpg",
		Content: "This life’s hard, but it’s harder if you’re stupid.",
	},
	{
		Author: "John Green",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Maybe our favorite quotations say more about us than about the stories and people we're quoting.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Wit beyond measure is man’s greatest treasure.",
	},
	{
		Author: "Jodi PicoultMy Sister's Keeper",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "Let me tell you this: if you meet a loner, no matter what they tell you, it's not because they enjoy solitude. It's because they have tried to blend into the world before, and people continue to disappoint them.",
	},
	{
		Author: "Herbert Bayard Swope",
		AuthorAvatar: "https://images.gr-assets.com/authors/1442529297p2/2333810.jpg",
		Content: "I can't give you a sure-fire formula for success, but I can give you a formula for failure: try to please everybody all the time.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "It takes ten times as long to put yourself back together as it does to fall apart.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "Do what you feel in your heart to be right – for you’ll be criticized anyway.",
	},
	{
		Author: "Douglas Adams",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "I refuse to answer that question on the grounds that I don't know the answer",
	},
	{
		Author: "Jack KerouacOn the Road",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "[...]the only people for me are the mad ones, the ones who are mad to live, mad to talk, mad to be saved, desirous of everything at the same time, the ones who never yawn or say a commonplace thing, but burn, burn, burn like fabulous yellow roman candles exploding like spiders across the stars and in the middle you see the blue centerlight pop and everybody goes “Awww!",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "We love the things we love for what they are.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "So we beat on, boats against the current, borne back ceaselessly into the past.",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "Sometimes I can hear my bones straining under the weight of all the lives I'm not living.",
	},
	{
		Author: "Ray Bradbury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "You don't have to burn books to destroy a culture. Just get people to stop reading them.",
	},
	{
		Author: "Henry Wadsworth Longfellow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211861766p2/2697.jpg",
		Content: "Every man has his secret sorrows which the world knows not; and often times we call a man cold when he is only sad.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I would die for you. But I won't live for you.",
	},
	{
		Author: "Virginia WoolfA Room of One's Own",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "Lock up your libraries if you like; but there is no gate, no lock, no bolt that you can set upon the freedom of my mind.",
	},
	{
		Author: "Mahatma GandhiAll Men are Brothers: Autobiographical Reflections",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "The weak can never forgive. Forgiveness is the attribute of the strong.",
	},
	{
		Author: "Jane AustenPride And Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "In vain have I struggled. It will not do. My feelings will not be repressed. You must allow me to tell you how ardently I admire and love you.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "There is only one thing that makes a dream impossible to achieve: the fear of failure.",
	},
	{
		Author: "Dalai Lama XIV",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241989386p2/570218.jpg",
		Content: "Happiness is not something ready made. It comes from your own actions.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "We came to see Jace. Is he alright?\"\"I don't know,\" Magnus said. \"Does he normally just lie on the floor like that without moving?",
	},
	{
		Author: "Rainer Maria RilkeLetters to a Young Poet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493785350p2/7906.jpg",
		Content: "Perhaps all the dragons in our lives are princesses who are only waiting to see us act, just once, with beauty and courage. Perhaps everything that frightens us is, in its deepest essence, something helpless that wants our love.",
	},
	{
		Author: "Neil GaimanFables \u0026 Reflections",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Sometimes you wake up. Sometimes the fall kills you. And sometimes, when you fall, you fly.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "For instance, on the planet Earth, man had always assumed that he was more intelligent than dolphins because he had achieved so much—the wheel, New York, wars and so on—whilst all the dolphins had ever done was muck about in the water having a good time. But conversely, the dolphins had always believed that they were far more intelligent than man—for precisely the same reasons.",
	},
	{
		Author: "Chris Rock",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206659889p2/29648.jpg",
		Content: "You know the world is going crazy when the best rapper is a white guy, the best golfer is a black guy, the tallest guy in the NBA is Chinese, the Swiss hold the America's Cup, France is accusing the U.S. of arrogance, Germany doesn't want to go to war, and the three most powerful men in America are named \"Bush\", \"Dick\", and \"Colin.\" Need I say more?",
	},
	{
		Author: "Aldous HuxleyComplete Essays 2, 1926-29",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "Facts do not cease to exist because they are ignored.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "A good friend will always stab you in the front.",
	},
	{
		Author: "Kahlil GibranThe Prophet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "You talk when you cease to be at peace with your thoughts.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "In the depth of winter, I finally learned that within me there lay an invincible summer.",
	},
	{
		Author: "George R.R. Martin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Sleep is good, he said, and books are better.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "Be kind, for everyone you meet is fighting a harder battle.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It means 'Shadowhunters: Looking Better in Black Than the Widows of our Enemies Since 1234'.",
	},
	{
		Author: "Théophile Gautier",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217372455p2/103384.jpg",
		Content: "Chance is perhaps the pseudonym of God when he does not want to sign.",
	},
	{
		Author: "Chuck PalahniukDiary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "We all die. The goal isn't to live forever, the goal is to create something that will.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I, with a deeper instinct, choose a man who compels my strength, who makes enormous demands on me, who does not doubt my courage or my toughness, who does not believe me naïve or innocent, who has the courage to treat me like a woman.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "When I look at my room, I see a girl who loves books.",
	},
	{
		Author: "Zelda Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1427041106p2/28243.jpg",
		Content: "Nobody has ever measured, not even poets, how much the heart can hold.",
	},
	{
		Author: "Nicole KraussThe History of Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285130428p2/2633.jpg",
		Content: "Once upon a time there was a boy who loved a girl, and her laughter was a question he wanted to spend his whole life answering.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Destroying things is much easier than making them.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "If more of us valued food and cheer and song above hoarded gold, it would be a merrier world.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I've never fooled anyone. I've let people fool themselves. They didn't bother to find out who and what I was. Instead they would invent a character for me. I wouldn't argue with them. They were obviously loving somebody I wasn't.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "A clever person solves a problem. A wise person avoids it.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Substitute 'damn' every time you're inclined to write 'very;' your editor will delete it and the writing will be just as it should be.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The best love is the kind that awakens the soul and makes us reach for more, that plants a fire in our hearts and brings peace to our minds. And that's what you've given me. That's what I'd hoped to give you forever",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "I love sleep. My life has the tendency to fall apart when I'm awake, you know?",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I don't want to go to heaven. None of my friends are there.",
	},
	{
		Author: "Lady Gaga",
		AuthorAvatar: "https://images.gr-assets.com/authors/1304697345p2/2945649.jpg",
		Content: "Some women choose to follow men, and some women choose to follow their dreams. If you're wondering which way to go, remember that your career will never wake up and tell you that it doesn't love you anymore.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You're a prefect? Oh Ronnie! That's everyone in the family!\"\"What are Fred and I? Next door neighbors?",
	},
	{
		Author: "Billy SundayBilly Sunday, the Man and His Message: With His Own Words Which Have Won Thousands for Christ",
		AuthorAvatar: "https://images.gr-assets.com/authors/1251485051p2/2319780.jpg",
		Content: "Going to church doesn’t make you a Christian any more than going to a garage makes you an automobile.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "There are two basic motivating forces: fear and love. When we are afraid, we pull back from life. When we are in love, we open to all that life has to offer with passion, excitement, and acceptance. We need to learn to love ourselves first, in all our glory and our imperfections. If we cannot love ourselves, we cannot fully open to our ability to love others or our potential to create. Evolution and all hopes for a better world rest in the fearlessness and open-hearted vision of people who embrace life.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "We fell in love, despite our differences, and once we did, something rare and beautiful was created. For me, love like that has only happened once, and that's why every minute we spent together has been seared in my memory. I'll never forget a single moment of it.",
	},
	{
		Author: "Neil GaimanAmerican Gods",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I can believe things that are true and things that aren't true and I can believe things where nobody knows if they're true or not. I can believe in Santa Claus and the Easter Bunny and the Beatles and Marilyn Monroe and Elvis and Mister Ed. Listen - I believe that people are perfectable, that knowledge is infinite, that the world is run by secret banking cartels and is visited by aliens on a regular basis, nice ones that look like wrinkled lemurs and bad ones who mutilate cattle and want our water and our women. I believe that the future sucks and I believe that the future rocks and I believe that one day White Buffalo Woman is going to come back and kick everyone's ass. I believe that all men are just overgrown boys with deep problems communicating and that the decline in good sex in America is coincident with the decline in drive-in movie theaters from state to state. I believe that all politicians are unprincipled crooks and I still believe that they are better than the alternative. I believe that California is going to sink into the sea when the big one comes, while Florida is going to dissolve into madness and alligators and toxic waste. I believe that antibacterial soap is destroying our resistance to dirt and disease so that one day we'll all be wiped out by the common cold like martians in War of the Worlds. I believe that the greatest poets of the last century were Edith Sitwell and Don Marquis, that jade is dried dragon sperm, and that thousands of years ago in a former life I was a one-armed Siberian shaman. I believe that mankind's destiny lies in the stars. I believe that candy really did taste better when I was a kid, that it's aerodynamically impossible for a bumble bee to fly, that light is a wave and a particle, that there's a cat in a box somewhere who's alive and dead at the same time (although if they don't ever open the box to feed it it'll eventually just be two different kinds of dead), and that there are stars in the universe billions of years older than the universe itself. I believe in a personal god who cares about me and worries and oversees everything I do. I believe in an impersonal god who set the universe in motion and went off to hang with her girlfriends and doesn't even know that I'm alive. I believe in an empty and godless universe of causal chaos, background noise, and sheer blind luck. I believe that anyone who says sex is overrated just hasn't done it properly. I believe that anyone who claims to know what's going on will lie about the little things too. I believe in absolute honesty and sensible social lies. I believe in a woman's right to choose, a baby's right to live, that while all human life is sacred there's nothing wrong with the death penalty if you can trust the legal system implicitly, and that no one but a moron would ever trust the legal system. I believe that life is a game, that life is a cruel joke, and that life is what happens when you're alive and that you might as well lie back and enjoy it.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "It's only after we've lost everything that we're free to do anything.",
	},
	{
		Author: "William ShakespeareHamlet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Doubt thou the stars are fire;Doubt that the sun doth move;Doubt truth to be a liar;But never doubt I love.",
	},
	{
		Author: "Stephenie MeyerTwilight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "And so the lion fell in love with the lamb…\" he murmured. I looked away, hiding my eyes as I thrilled to the word.\"What a stupid lamb,\" I sighed.\"What a sick, masochistic lion.",
	},
	{
		Author: "William ShakespeareJulius Caesar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "The fault, dear Brutus, is not in our stars, but in ourselves.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "You can't stay in your corner of the Forest waiting for others to come to you. You have to go to them sometimes.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "I DON'T CARE!\" Harry yelled at them, snatching up a lunascope and throwing it into the fireplace. \"I'VE HAD ENOUGH, I'VE SEEN ENOUGH, I WANT OUT, I WANT IT TO END, I DON'T CARE ANYMORE!\"\"You do care,\" said Dumbledore. He had not flinched or made a single move to stop Harry demolishing his office. His expression was calm, almost detached. \"You care so much you feel as though you will bleed to death with the pain of it.",
	},
	{
		Author: "Francis Bacon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397100712p2/50964.jpg",
		Content: "Some books should be tasted, some devoured, but only a few should be chewed and digested thoroughly.",
	},
	{
		Author: "Frank HerbertDune",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168661521p2/58.jpg",
		Content: "I must not fear. Fear is the mind-killer. Fear is the little-death that brings total obliteration. I will face my fear. I will permit it to pass over me and through me. And when it has gone past I will turn the inner eye to see its path. Where the fear has gone there will be nothing. Only I will remain.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "And those who were seen dancing were thought to be insane by those who could not hear the music.",
	},
	{
		Author: "May Sarton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1413742677p2/13166.jpg",
		Content: "We have to dare to be ourselves, however frightening or strange that self may prove to be.",
	},
	{
		Author: "Jane AustenJane Austen's Letters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I do not want people to be very agreeable, as it saves me the trouble of liking them a great deal.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Do you remember back at the hotel when you promised that if we lived, you’d get dressed up in a nurse’s outfit and give me a sponge bath?\" asked Jace.\"It was Simon who promised you the sponge bath.\"\"As soon as I’m back on my feet, handsome,\" said Simon.\"I knew we should have left you a rat.",
	},
	{
		Author: "Charlaine Harris",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399317093p2/17061.jpg",
		Content: "Here’s to books, the cheapest vacation you can buy.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "There are no good girls gone wrong - just bad girls found out.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "You will always be fond of me. I represent to you all the sins you never had the courage to commit.",
	},
	{
		Author: "Lemony Snicket",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Reading is one form of escape. Running for your life is another.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Anyone who lives within their means suffers from a lack of imagination.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It is the unknown we fear when we look upon death and darkness, nothing more.",
	},
	{
		Author: "Shana Abe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1219732484p2/15987.jpg",
		Content: "I heard what you said. I’m not the silly romantic you think. I don’t want the heavens or the shooting stars. I don’t want gemstones or gold. I have those things already. I want…a steady hand. A kind soul. I want to fall asleep, and wake, knowing my heart is safe. I want to love, and be loved.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1386666158p2/11996.jpg",
		Content: "Eating and reading are two pleasures that combine admirably.",
	},
	{
		Author: "Phyllis Diller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209492226p2/271137.jpg",
		Content: "Never go to bed mad. Stay up and fight.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Man is the only creature who refuses to be what he is.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Never laugh at live dragons.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "I believe in everything until it's disproved. So I believe in fairies, the myths, dragons. It all exists, even if it's in your mind. Who's to say that dreams and nightmares aren't as real as the here and now?",
	},
	{
		Author: "L.M. Montgomery",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188896723p2/5350.jpg",
		Content: "Isn't it nice to think that tomorrow is a new day with no mistakes in it yet?",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Books are the perfect entertainment: no commercials, no batteries, hours of enjoyment for each dollar spent. What I wonder is why everybody doesn't carry a book around for those inevitable dead spots in life.",
	},
	{
		Author: "William Shakespeare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "I would challenge you to a battle of wits, but I see you are unarmed!",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Reality is merely an illusion, albeit a very persistent one.",
	},
	{
		Author: "Voltaire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393683411p2/5754446.jpg",
		Content: "Let us read, and let us dance; these two amusements will never do any harm to the world.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "The planet is fine. The people are fucked.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I hope that in this year to come, you make mistakes. Because if you are making mistakes...you're Doing Something.",
	},
	{
		Author: "Robert FulghumAll I Really Need to Know I Learned in Kindergarten: Uncommon Thoughts On Common Things",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198520878p2/19630.jpg",
		Content: "I believe that imagination is stronger than knowledge. That myth is more potent than history. That dreams are more powerful than facts. That hope always triumphs over experience. That laughter is the only cure for grief. And I believe that love is stronger than death.",
	},
	{
		Author: "Jonathan Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183238507p2/1831.jpg",
		Content: "May you live every day of your life.",
	},
	{
		Author: "E.B. White",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198519412p2/988142.jpg",
		Content: "If the world were merely seductive, that would be easy. If it were merely challenging, that would be no problem. But I arise in the morning torn between a desire to improve the world and a desire to enjoy the world. This makes it hard to plan the day.",
	},
	{
		Author: "Anne FrankThe Diary of a Young Girl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "It's really a wonder that I haven't dropped all my ideals, because they seem so absurd and impossible to carry out. Yet I keep them, because in spite of everything, I still believe that people are really good at heart.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "We should all start to live before we get too old. Fear is stupid. So are regrets.",
	},
	{
		Author: "Stephen KingDifferent Seasons",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Get busy living or get busy dying.",
	},
	{
		Author: "Robert FrostStopping by Woods on a Snowy Evening",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "These woods are lovely, dark and deep,But I have promises to keep,And miles to go before I sleep,And miles to go before I sleep.",
	},
	{
		Author: "Haruki MurakamiSouth of the Border, West of the Sun",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "I think you still love me, but we can’t escape the fact that I’m not enough for you. I knew this was going to happen. So I’m not blaming you for falling in love with another woman. I’m not angry, either. I should be, but I’m not. I just feel pain. A lot of pain. I thought I could imagine how much this would hurt, but I was wrong.",
	},
	{
		Author: "Arundhati RoyThe Cost of Living",
		AuthorAvatar: "https://images.gr-assets.com/authors/1496705394p2/6134.jpg",
		Content: "To love. To be loved. To never forget your own insignificance. To never get used to the unspeakable violence and the vulgar disparity of life around you. To seek joy in the saddest places. To pursue beauty to its lair. To never simplify what is complicated or complicate what is simple. To respect strength, never power. Above all, to watch. To try and understand. To never look away. And never, never to forget.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Science without religion is lame, religion without science is blind.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "And I'm suppose to sit by while you date boys and fall in love with someone else, get married...?\" His voice tightened. \"And meanwhile, I'll die a little bit more every day, watching.",
	},
	{
		Author: "Brigham Young",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234971501p2/575321.jpg",
		Content: "You educate a man; you educate a man. You educate a woman; you educate a generation.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Reader, suppose you were an idiot. And suppose you were a member of Congress. But I repeat myself.",
	},
	{
		Author: "Evans G. ValensThe Other Side of the Mountain: The Story of Jill Kinmont",
		AuthorAvatar: "",
		Content: "How lucky I am to have known somebody and something that saying goodbye to is so damned awful.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "We write to taste life twice, in the moment and in retrospect.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "There will come a time when all of us are dead. All of us. There will come a time when there are no human beings remaining to remember that anyone ever existed or that our species ever did anything. There will be no one left to remember Aristotle or Cleopatra, let alone you. Everything that we did and built and wrote and thought and discovered will be forgotten and all of this will have been for naught. Maybe that time is coming soon and maybe it is millions of years away, but even if we survive the collapse of our sun, we will not survive forever. There was time before organisms experienced consciousness, and there will be time after. And if the inevitability of human oblivion worries you, I encourage you to ignore it. God knows that’s what everyone else does.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Deep in the meadow, hidden far awayA cloak of leaves, a moonbeam rayForget your woes and let your troubles layAnd when it's morning again, they'll wash awayHere it's safe, here it's warmHere the daisies guard you from every harmHere your dreams are sweet and tomorrow brings them trueHere is the place where I love you.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "All the world is made of faith, and trust, and pixie dust.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "When someone shows you who they are believe them; the first time.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "I have hated words and I have loved them, and I hope I have made them right.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Let us find the dam snack bar,\" Zoe said. \"We should eat while we can.\"Grover cracked a smile. \"The dam snack bar?\"Zoe blinked. \"Yes. What is funny?\"\"Nothing,\" Grover said, trying to keep a straight face. \"I could use some dam french fries.\"Even Thalia smiled at that. \"And I need to use the dam restroom.\"...I started cracking up, and Thalia and Grover joined in, while Zoe just looked at me. \"I do not understand.\"\"I want to use the dam water fountain,\" Grover said.\"And...\" Thalia tried to catch her breath. \"I want to buy a dam t-shirt.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I might be in love with you.\" He smiles a little. \"I'm waiting until I'm sure to tell you, though.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "We loved with a love that was more than love.",
	},
	{
		Author: "Dr. SeussThe Lorax",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Unless someone like you cares a whole awful lot,Nothing is going to get better. It's not.",
	},
	{
		Author: "C.G. Jung",
		AuthorAvatar: "https://images.gr-assets.com/authors/1471024439p2/38285.jpg",
		Content: "The meeting of two personalities is like the contact of two chemical substances: if there is any reaction, both are transformed.",
	},
	{
		Author: "Philip K. DickVALIS",
		AuthorAvatar: "https://images.gr-assets.com/authors/1264613853p2/4764.jpg",
		Content: "It is sometimes an appropriate response to reality to go insane.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I wanted so badly to lie down next to her on the couch, to wrap my arms around her and sleep. Not fuck, like in those movies. Not even have sex. Just sleep together in the most innocent sense of the phrase. But I lacked the courage and she had a boyfriend and I was gawky and she was gorgeous and I was hopelessly boring and she was endlessly fascinating. So I walked back to my room and collapsed on the bottom bunk, thinking that if people were rain, I was drizzle and she was hurricane.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "This is a good sign, having a broken heart. It means we have tried for something.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "Every heart sings a song, incomplete, until another heart whispers back. Those who wish to sing always find a song. At the touch of a lover, everyone becomes a poet.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Some people care too much. I think it's called love.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Out beyond ideas of wrongdoing and rightdoing there is a field.I'll meet you there.When the soul lies down in that grassthe world is too full to talk about.",
	},
	{
		Author: "Emily Dickinson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536260p2/7440.jpg",
		Content: "Hope is the thing with feathers That perches in the soul And sings the tune without the words And never stops at all.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "When I was 5 years old, my mother always told me that happiness was the key to life. When I went to school, they asked me what I wanted to be when I grew up. I wrote down ‘happy’. They told me I didn’t understand the assignment, and I told them they didn’t understand life.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I am not young enough to know everything.",
	},
	{
		Author: "Laurence J. Peter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1314950364p2/182617.jpg",
		Content: "If a cluttered desk is a sign of a cluttered mind, of what, then, is an empty desk a sign?",
	},
	{
		Author: "Augustine of Hippo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1404280110p2/6819578.jpg",
		Content: "The world is a book and those who do not travel read only one page.",
	},
	{
		Author: "Trisha Yearwood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1327280419p2/1332085.jpg",
		Content: "What's meant to be will always find a way",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "If you expect nothing from somebody you are never disappointed.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Always do what you are afraid to do.",
	},
	{
		Author: "Jack KerouacThe Dharma Bums",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "One day I will find the right words, and they will be simple.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "The secret of life, though, is to fall seven times and to get up eight times.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "The boy never cried again, and he never forgot what he'd learned: that to love is to destroy, and that to be loved is to be the one destroyed.",
	},
	{
		Author: "AristotleMetaphysics",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390143800p2/2192.jpg",
		Content: "It is the mark of an educated mind to be able to entertain a thought without accepting it.",
	},
	{
		Author: "John Greenleaf WhittierMaud Muller - Pamphlet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1220983305p2/267703.jpg",
		Content: "Of all sad words of tongue or pen, the saddest are these, 'It might have been.",
	},
	{
		Author: "Mary Oliver",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1389642386p2/1921.jpg",
		Content: "Tell me, what is it you plan to do with your one wild and precious life?",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "Make it a rule never to give a child a book you would not read yourself.",
	},
	{
		Author: "Martin Luther King Jr.I Have a Dream: Writings and Speeches That Changed the World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Our lives begin to end the day we become silent about things that matter.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Faith is taking the first step even when you can't see the whole staircase.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "She was a girl who knew how to be happy even when she was sad. And that’s important—you know ",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Books so special and rare and yours that advertising your affection feels like a betrayal.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Faithless is he that says farewell when the road darkens.",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "A life spent making mistakes is not only more honorable, but more useful than a life spent doing nothing.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I think that if I ever have kids, and they are upset, I won't tell them that people are starving in China or anything like that because it wouldn't change the fact that they were upset. And even if somebody else has it much worse, that doesn't really change the fact that you have what you have.",
	},
	{
		Author: "Alexandre Dumas",
		AuthorAvatar: "https://images.gr-assets.com/authors/1279049943p2/4785.jpg",
		Content: "There is neither happiness nor misery in the world; there is only the comparison of one state with another, nothing more. He who has felt the deepest grief is best able to experience supreme happiness. We must have felt what it is to die, Morrel, that we may appreciate the enjoyments of life. \" Live, then, and be happy, beloved children of my heart, and never forget, that until the day God will deign to reveal the future to man, all human wisdom is contained in these two words, 'Wait and Hope.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Well, I’m not kissing the mundane,\" said Jace. \"I’d rather stay down here and rot.\"\"Forever?\" said Simon. \"Forever’s an awfully long time.\"Jace raised his eyebrows. \"I knew it,\" he said. \"You want to kiss me, don’t you?",
	},
	{
		Author: "Cassandra ClareClockwork Princess",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Life is a book and there are a thousand pages I have not yet read.",
	},
	{
		Author: "Stephen Hawking",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197404653p2/1401.jpg",
		Content: "One, remember to look up at the stars and not down at your feet. Two, never give up work. Work gives you meaning and purpose and life is empty without it. Three, if you are lucky enough to find love, remember it is there and don't throw it away.",
	},
	{
		Author: "Stephen KingStorm of the Century: An Original Screenplay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "When his life was ruined, his family killed, his farm destroyed, Job knelt down on the ground and yelled up to the heavens, \"Why god? Why me?\" and the thundering voice of God answered, There's just something about you that pisses me off.",
	},
	{
		Author: "Frederick Douglass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1511594938p2/18943.jpg",
		Content: "Once you learn to read, you will be forever free.",
	},
	{
		Author: "Tom Bodett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1258767561p2/58627.jpg",
		Content: "They say a person needs just three things to be truly happy in this world: someone to love, something to do, and something to hope for.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "In the end, we will remember not the words of our enemies, but the silence of our friends.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "You have enemies? Good. That means you've stood up for something, sometime in your life.",
	},
	{
		Author: "Ayn RandAtlas Shrugged",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "Do not let your fire go out, spark by irreplaceable spark in the hopeless swamps of the not-quite, the not-yet, and the not-at-all. Do not let the hero in your soul perish in lonely frustration for the life you deserved and have never been able to reach. The world you desire can be won. It exists.. it is real.. it is possible.. it's yours.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I am a man\" he told her, \"and men do not consume pink beverages. Get thee gone woman, and bring me something brown.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Everyone, at some point in their lives, wakes up in the middle of the night with the feeling that they are all alone in the world, and that nobody loves them now and that nobody will ever love them, and that they will never have a decent night's sleep again and will spend their lives wandering blearily around a loveless landscape, hoping desperately that their circumstances will improve, but suspecting, in their heart of hearts, that they will remain unloved forever. The best thing to do in these circumstances is to wake somebody else up, so that they can feel this way, too.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It is so hard to leave—until you leave. And then it is the easiest goddamned thing in the world.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "I wanted you to see what real courage is, instead of getting the idea that courage is a man with a gun in his hand. It's when you know you're licked before you begin, but you begin anyway and see it through no matter what.- Atticus Finch",
	},
	{
		Author: "Henry Ford",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397426898p2/203714.jpg",
		Content: "Whether you think you can, or you think you can't--you're right.",
	},
	{
		Author: "Douglas AdamsThe Restaurant at the End of the Universe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "There is a theory which states that if ever anyone discovers exactly what the Universe is for and why it is here, it will instantly disappear and be replaced by something even more bizarre and inexplicable. There is another theory which states that this has already happened.",
	},
	{
		Author: "Walt Disney CompanyMulan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289112683p2/3510823.jpg",
		Content: "The flower that blooms in adversity is the rarest and most beautiful of all.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "And the day came when the risk to remain tight in a bud was more painful than the risk it took to blossom.",
	},
	{
		Author: "Arthur Schopenhauer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1404836700p2/11682.jpg",
		Content: "Talent hits a target no one else can hit. Genius hits a target no one else can see.",
	},
	{
		Author: "Kurt VonnegutPlayer Piano",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "I want to stand as close to the edge as I can without going over. Out on the edge you see all kinds of things you can't see from the center.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Not everything is about you,\" Clary said furiously.\"Possibly,\" Jace said, \"but you do have to admit that the majority of things are.",
	},
	{
		Author: "Milan Kundera",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465275207p2/6343.jpg",
		Content: "Two people in love, alone, isolated from the world, that's beautiful.",
	},
	{
		Author: "Anne Frank",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "Think of all the beauty still left around you and be happy.",
	},
	{
		Author: "Haruki MurakamiWhat I Talk About When I Talk About Running",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Pain is inevitable. Suffering is optional.",
	},
	{
		Author: "Ellen DeGeneresSeriously... I'm Kidding",
		AuthorAvatar: "https://images.gr-assets.com/authors/1328758901p2/40648.jpg",
		Content: "Accept who you are. Unless you're a serial killer.",
	},
	{
		Author: "E.B. WhiteCharlotte's Web",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198519412p2/988142.jpg",
		Content: "Why did you do all this for me?' he asked. 'I don't deserve it. I've never done anything for you.' 'You have been my friend,' replied Charlotte. 'That in itself is a tremendous thing.",
	},
	{
		Author: "Bertrand Russell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457021442p2/17854.jpg",
		Content: "There are two motives for reading a book; one, that you enjoy it; the other, that you can boast about it.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What the hell is that?\" I laughed.\"It's my fox hat.\"\"Your fox hat?\"\"Yeah, Pudge. My fox hat.\"\"Why are you wearing your fox hat?\" I asked.\"Because no one can catch the motherfucking fox.",
	},
	{
		Author: "Chad SuggMonsters Under Your Head",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419367972p2/1934928.jpg",
		Content: "If you're reading this...Congratulations, you're alive.If that's not something to smile about,then I don't know what is.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "I've never been lonely. I've been in a room -- I've felt suicidal. I've been depressed. I've felt awful -- awful beyond all -- but I never felt that one other person could enter that room and cure what was bothering me...or that any number of people could enter that room. In other words, loneliness is something I've never been bothered with because I've always had this terrible itch for solitude. It's being at a party, or at a stadium full of people cheering for something, that I might feel loneliness. I'll quote Ibsen, \"The strongest men are the most alone.\" I've never thought, \"Well, some beautiful blonde will come in here and give me a fuck-job, rub my balls, and I'll feel good.\" No, that won't help. You know the typical crowd, \"Wow, it's Friday night, what are you going to do? Just sit there?\" Well, yeah. Because there's nothing out there. It's stupidity. Stupid people mingling with stupid people. Let them stupidify themselves. I've never been bothered with the need to rush out into the night. I hid in bars, because I didn't want to hide in factories. That's all. Sorry for all the millions, but I've never been lonely. I like myself. I'm the best form of entertainment I have. Let's drink more wine!",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "The heart was made to be broken.",
	},
	{
		Author: "George CarlinWhen Will Jesus Bring The Pork Chops?",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "Here's all you have to know about men and women: women are crazy, men are stupid. And the main reason women are crazy is that men are stupid.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "We live and breathe words. .... It was books that made me feel that perhaps I was not completely alone. They could be honest with me, and I with them. Reading your words, what you wrote, how you were lonely sometimes and afraid, but always brave; the way you saw the world, its colors and textures and sounds, I felt--I felt the way you thought, hoped, felt, dreamt. I felt I was dreaming and thinking and feeling with you. I dreamed what you dreamed, wanted what you wanted--and then I realized that truly I just wanted you.",
	},
	{
		Author: "Henry Ward Beecher",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198680376p2/425221.jpg",
		Content: "Where is human nature so weak as in the bookstore?",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "No relationship is perfect, ever. There are always some ways you have to bend, to compromise, to give something up in order to gain something greater...The love we have for each other is bigger than these small differences. And that's the key. It's like a big pie chart, and the love in a relationship has to be the biggest piece. Love can make up for a lot.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Monsters are real, and ghosts are real too. They live inside us, and sometimes, they win.",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "And now here is my secret, a very simple secret: It is only with the heart that one can see rightly; what is essential is invisible to the eye.",
	},
	{
		Author: "Dave Matthews Band",
		AuthorAvatar: "",
		Content: "A guy and a girl can be just friends, but at one point or another, they will fall for each other...Maybe temporarily, maybe at the wrong time, maybe too late, or maybe forever",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "May I see you again?\" he asked. There was an endearing nervousness in his voice. I smiled. \"Sure.\"\"Tomorrow?\" he asked.\"Patience, grasshopper,\" I counseled. \"You don't want to seem overeager. \"Right, that's why I said tomorrow,\" he said. \"I want to see you again tonight. But I'm willing to wait all night and much of tomorrow.\" I rolled my eyes. \"I'm serious,\" he said. \"You don't even know me,\" I said. I grabbed the book from the center console. \"How about I call you when I finish this?\"\"But you don't even have my phone number,\" he said.\"I strongly suspect you wrote it in this book.\" He broke out into that goofy smile. \"And you say we don't know each other.",
	},
	{
		Author: "Erasmus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1505744161p2/706620.jpg",
		Content: "When I have a little money, I buy books; and if I have any left, I buy food and clothes.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "If my life is going to mean anything, I have to live it myself.",
	},
	{
		Author: "Emily BrontëWuthering Heights",
		AuthorAvatar: "https://images.gr-assets.com/authors/1473229007p2/4191.jpg",
		Content: "He's more myself than I am. Whatever our souls are made of, his and mine are the same.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "That's always seemed so ridiculous to me, that people want to be around someone because they're pretty. It's like picking your breakfeast cereals based on color instead of taste.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Don’t go around saying the world owes you a living. The world owes you nothing. It was here first.",
	},
	{
		Author: "Rainbow RowellEleanor \u0026 Park",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342324527p2/4208569.jpg",
		Content: "Eleanor was right. She never looked nice. She looked like art, and art wasn't supposed to look nice; it was supposed to make you feel something.",
	},
	{
		Author: "J.R.R. TolkienThe Two Towers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "There is some good in this world, and it's worth fighting for.",
	},
	{
		Author: "Pat ConroyThe Prince of Tides",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342711176p2/6942.jpg",
		Content: "You get a little moody sometimes but I think that's because you like to read. People that like to read are always a little fucked up.",
	},
	{
		Author: "Edgar Allan PoeEleonora",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "Those who dream by day are cognizant of many things which escape those who dream only by night.",
	},
	{
		Author: "Jerome K. Jerome",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335971515p2/3352.jpg",
		Content: "I like work: it fascinates me. I can sit and look at it for hours.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Malachi scowled. \"I don't remember the Clave inviting you into the Glass City, Magnus Bane.\"\"They didn't,\" Magnus said. \"Your wards are down.\"\"Really?\" the Consul's voice dripped sarcasm. \"I hadn't noticed.\"Magnus looked concerned. \"That's terrible. Someone should have told you.\" He glanced at Luke. \"Tell him the wards are down.",
	},
	{
		Author: "Jane AustenSense and Sensibility",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "The more I know of the world, the more I am convinced that I shall never see a man whom I can really love. I require so much!",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "For those who believe in God, most of the big questions are answered. But for those of us who can't readily accept the God formula, the big answers don't remain stone-written. We adjust to new conditions and discoveries. We are pliable. Love need not be a command nor faith a dictum. I am my own god. We are here to unlearn the teachings of the church, state, and our educational system. We are here to drink beer. We are here to kill war. We are here to laugh at the odds and live our lives so well that Death will tremble to take us.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Every saint has a past, and every sinner has a future.",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "Who controls the past controls the future. Who controls the present controls the past.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Sarcasm is the last refuge of the imaginatively bankrupt.",
	},
	{
		Author: "Bette Midler",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207797480p2/548179.jpg",
		Content: "Give a girl the right shoes, and she can conquer the world.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "You can talk with someone for years, everyday, and still, it won't mean as much as what you can have when you sit in front of someone, not saying a word, yet you feel that person with your heart, you feel like you have known the person for forever.... connections are made with the heart, not the tongue.",
	},
	{
		Author: "Laurie Halse AndersonSpeak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376224335p2/10003.jpg",
		Content: "THE FIRST TEN LIES THEY TELL YOU IN HIGH SCHOOL 1. We are here to help you. 2. You will have time to get to your class before the bell rings. 3. The dress code will be enforced. 4. No smoking is allowed on school grounds. 5. Our football team will win the championship this year. 6. We expect more of you here. 7. Guidance counselors are always available to listen. 8. Your schedule was created with you in mind. 9. Your locker combination is private. 10. These will be the years you look back on fondly. TEN MORE LIES THEY TELL YOU IN HIGH SCHOOL 1. You will use algebra in your adult lives. 2. Driving to school is a privilege that can be taken away. 3. Students must stay on campus during lunch. 4. The new text books will arrive any day now. 5. Colleges care more about you than your SAT scores. 6. We are enforcing the dress code. 7. We will figure out how to turn off the heat soon. 8. Our bus drivers are highly trained professionals. 9. There is nothing wrong with summer school. 10. We want to hear what you have to say.",
	},
	{
		Author: "Fyodor DostoyevskyThe Brothers Karamazov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "Above all, don't lie to yourself. The man who lies to himself and listens to his own lie comes to a point that he cannot distinguish the truth within him, or around him, and so loses all respect for himself and for others. And having no respect he ceases to love.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I'm coming back into focus when Caesar asks him if he has a girlfriend back home. Peeta hesitates, then gives an unconvincing shake of his head.Handsome lad like you. There must be some special girl. Come on, what’s her name?\" says Caesar.Peeta sighs. \"Well, there is this one girl. I’ve had a crush on her ever since I can remember. But I’m pretty sure she didn’t know I was alive until the reaping.\"Sounds of sympathy from the crowd. Unrequited love they can relate to.She have another fellow?\" asks Caesar.I don’t know, but a lot of boys like her,\" says Peeta.So, here’s what you do. You win, you go home. She can’t turn you down then, eh?\" says Caesar encouragingly.I don’t think it’s going to work out. Winning...won’t help in my case,\" says Peeta.Why ever not?\" says Caesar, mystified.Peeta blushes beet red and stammers out. \"Because...because...she came here with me.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Augustus Waters was a self-aggrandizing bastard. But we forgive him. We forgive him not because he had a heart as figuratively good as his literal one sucked, or because he knew more about how to hold a cigarette than any nonsmoker in history, or because he got eighteen years when he should've gotten more.''Seventeen,' Gus corrected.'I'm assuming you've got some time, you interupting bastard.'I'm telling you,' Isaac continued, 'Augustus Waters talked so much that he'd interupt you at his own funeral. And he was pretentious: Sweet Jesus Christ, that kid never took a piss without pondering the abundant metaphorical resonances of human waste production. And he was vain: I do not believe I have ever met a more physically attractive person who was more acutely aware of his own physical attractiveness.'But I will say this: When the scientists of the future show up at my house with robot eyes and they tell me to try them on, I will tell the scientists to screw off, because I do not want to see a world without him.'I was kind of crying by then.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Your task is not to seek for love, but merely to seek and find all the barriers within yourself that you have built against it.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "If we knew what it was we were doing, it would not be called research, would it?",
	},
	{
		Author: "Henry David ThoreauWalden: Or, Life in the Woods",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "I learned this, at least, by my experiment: that if one advances confidently in the direction of his dreams, and endeavors to live the life which he has imagined, he will meet with a success unexpected in common hours.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Sometimes you climb out of bed in the morning and you think, I'm not going to make it, but you laugh inside — remembering all the times you've felt that way.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "I have no special talents. I am only passionately curious.",
	},
	{
		Author: "Martin Luther King Jr.A Testament of Hope: The Essential Writings and Speeches",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "I have decided to stick to love...Hate is too great a burden to bear.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "I did not attend his funeral, but I sent a nice letter saying I approved of it.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "We delight in the beauty of the butterfly, but rarely admit the changes it has gone through to achieve that beauty.",
	},
	{
		Author: "Søren Kierkegaard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202410387p2/6172.jpg",
		Content: "Life can only be understood backwards; but it must be lived forwards.",
	},
	{
		Author: "AnonymousHoly Bible: New International Version",
		AuthorAvatar: "",
		Content: "Love is always patient and kind. It is never jealous. Love is never boastful or conceited. It is never rude or selfish. It does not take offense and is not resentful. Love takes no pleasure in other people’s sins, but delights in the truth. It is always ready to excuse, to trust, to hope, and to endure whatever comes.",
	},
	{
		Author: "John MiltonParadise Lost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1438861655p2/9876.jpg",
		Content: "The mind is its own place, and in itself can make a heaven of hell, a hell of heaven..",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "It is a curious thing, the death of a loved one. We all know that our time in this world is limited, and that eventually all of us will end up underneath some sheet, never to wake up. And yet it is always a surprise when it happens to someone we know. It is like walking up the stairs to your bedroom in the dark, and thinking there is one more stair than there is. Your foot falls down, through the air, and there is a sickly moment of dark surprise as you try and readjust the way you thought of things.",
	},
	{
		Author: "Abigail Van Buren",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351267432p2/812826.jpg",
		Content: "The best index to a person's character is how he treats people who can't do him any good, and how he treats people who can't fight back.",
	},
	{
		Author: "Thomas Jefferson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400443180p2/1673.jpg",
		Content: "I cannot live without books.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You know, you could live a thousand lifetimes and not deserve him.",
	},
	{
		Author: "W.C. Fields",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331712829p2/82951.jpg",
		Content: "If at first you don't succeed, try, try again. Then quit. No use being a damn fool about it.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Is this the part where you say if I hurt her, you'll kill me?\"\"No\" Simon said, \"If you hurt Clary she's quite capable of killing you herself. Possibly with a variety of weapons.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "The most important thing is to enjoy your life—to be happy—it's all that matters.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "I have a history of making decisions very quickly about men. I have always fallen in love fast and without measuring risks. I have a tendency not only to see the best in everyone, but to assume that everyone is emotionally capable of reaching his highest potential. I have fallen in love more times than I care to count with the highest potential of a man, rather than with the man himself, and I have hung on to the relationship for a long time (sometimes far too long) waiting for the man to ascend to his own greatness. Many times in romance I have been a victim of my own optimism.",
	},
	{
		Author: "Carl Sagan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475953320p2/10538.jpg",
		Content: "Somewhere, something incredible is waiting to be known.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Never be bullied into silence. Never allow yourself to be made a victim. Accept no one’s definition of your life; define yourself.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I can't imagine a man really enjoying a book and reading it only once.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "You don't love someone for their looks, or their clothes, or for their fancy car, but because they sing a song only you can hear.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "God created war so that Americans would learn geography.",
	},
	{
		Author: "Paulo CoelhoBy the River Piedra I Sat Down and Wept",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Waiting is painful. Forgetting is painful. But not knowing which to do is the worst kind of suffering.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "A thing is not necessarily true because a man dies for it.",
	},
	{
		Author: "William ShakespeareHamlet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "This above all: to thine own self be true, And it must follow, as the night the day, Thou canst not then be false to any man.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I took a deep breath and listened to the old brag of my heart. I am, I am, I am.",
	},
	{
		Author: "George Orwell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "In a time of deceit telling the truth is a revolutionary act.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "I generally avoid temptation unless I can't resist it.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I'm not saying that everything is survivable. Just that everything except the last thing is.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "It is hard enough to remember my opinions, without also remembering my reasons for them!",
	},
	{
		Author: "Aristotle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390143800p2/2192.jpg",
		Content: "What is a friend? A single soul dwelling in two bodies.",
	},
	{
		Author: "Stephanie PerkinsAnna and the French Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407443106p2/3095893.jpg",
		Content: "For the two of us, home isn't a place. It is a person. And we are finally home.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Any fool can know. The point is to understand.",
	},
	{
		Author: "Haruki Murakami",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "And once the storm is over, you won’t remember how you made it through, how you managed to survive. You won’t even be sure, whether the storm is really over. But one thing is certain. When you come out of the storm, you won’t be the same person who walked in. That’s what this storm’s all about.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "The most courageous act is still to think for yourself. Aloud.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "All little girls should be told they are pretty, even if they aren't.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Not my daughter, you bitch!",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "When you're struggling with something, look at all the people around you and realize that every single person you see is struggling with something, and to them, it's just as hard as what you're going through.",
	},
	{
		Author: "Lloyd Alexander",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353430382p2/8924.jpg",
		Content: "Fantasy is hardly an escape from reality. It's a way of understanding it.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Experience is merely the name men gave to their mistakes.",
	},
	{
		Author: "Lemony SnicketThe Grim Grotto",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "People aren't either wicked or noble. They're like chef's salads, with good things and bad things chopped and mixed together in a vinaigrette of confusion and conflict.",
	},
	{
		Author: "Nicholas SparksMessage in a Bottle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Nothing that’s worthwhile is ever easy. Remember that.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "And by the way, everything in life is writable about if you have the outgoing guts to do it, and the imagination to improvise. The worst enemy to creativity is self-doubt.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You said you were going for a walk!? What kind of walk takes six hours?\"\"A long one?",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "It's not true that I had nothing on. I had the radio on.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "Libraries were full of ideas—perhaps the most dangerous and powerful of all weapons.",
	},
	{
		Author: "Maya AngelouLetter to My Daughter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "I can be changed by what happens to me. But I refuse to be reduced by it.(Popular misquote of \"You may not control all the events that happen to you, but you can decide not to be reduced by them.\")",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You could have had anything else in the world, and you asked for me.\"She smiled up at him. Filthy as he was, covered in blood and dirt, he was the most beautiful thing she'd ever seen.\"But I don't want anything else in the world.",
	},
	{
		Author: "Alan Moore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1304944713p2/3961.jpg",
		Content: "My experience of life is that it is not divided up into genres; it’s a horrifying, romantic, tragic, comical, science-fiction cowboy detective novel. You know, with a bit of pornography if you're lucky.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "They didn’t agree on much. In fact, they didn’t agree on anything. They fought all the time and challenged each other ever day. But despite their differences, they had one important thing in common. They were crazy about each other.",
	},
	{
		Author: "J.D. Salinger",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1504717993p2/10624.jpg",
		Content: "Don't ever tell anybody anything. If you do, you start missing everybody.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Do not go where the path may lead, go instead where there is no path and leave a trail.",
	},
	{
		Author: "F. Scott Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "The loneliest moment in someone’s life is when they are watching their whole world fall apart, and all they can do is stare blankly.",
	},
	{
		Author: "Neil GaimanThe Kindly Ones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Have you ever been in love? Horrible isn't it? It makes you so vulnerable. It opens your chest and it opens up your heart and it means that someone can get inside you and mess you up. You build up all these defenses, you build up a whole suit of armor, so that nothing can hurt you, then one stupid person, no different from any other stupid person, wanders into your stupid life...You give them a piece of you. They didn't ask for it. They did something dumb one day, like kiss you or smile at you, and then your life isn't your own anymore. Love takes hostages. It gets inside you. It eats you out and leaves you crying in the darkness, so simple a phrase like 'maybe we should be just friends' turns into a glass splinter working its way into your heart. It hurts. Not just in the imagination. Not just in the mind. It's a soul-hurt, a real gets-inside-you-and-rips-you-apart pain. I hate love.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "what matters most is how well you walk through the fire",
	},
	{
		Author: "Charles M. Schulz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590750p2/209672.jpg",
		Content: "Happiness is a warm puppy.",
	},
	{
		Author: "Roald DahlCharlie and the Chocolate Factory",
		AuthorAvatar: "https://images.gr-assets.com/authors/1311554908p2/4273.jpg",
		Content: "So please, oh please, we beg, we pray,Go throw your TV set away,And in its place you can installA lovely bookshelf on the wall.Then fill the shelves with lots of books.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Without pain, how could we know joy?' This is an old argument in the field of thinking about suffering and its stupidity and lack of sophistication could be plumbed for centuries but suffice it to say that the existence of broccoli does not, in any way, affect the taste of chocolate.",
	},
	{
		Author: "Jon Katz",
		AuthorAvatar: "",
		Content: "I think if I've learned anything about friendship, it's to hang in, stay connected, fight for them, and let them fight for you. Don't walk away, don't be distracted, don't be too busy or tired, don't take them for granted. Friends are part of the glue that holds life and faith together. Powerful stuff.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "All happy families are alike; each unhappy family is unhappy in its own way.",
	},
	{
		Author: "Fran LebowitzThe Fran Lebowitz Reader",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414953906p2/8127311.jpg",
		Content: "Think before you speak. Read before you think.",
	},
	{
		Author: "Katie McGarryTake Me On",
		AuthorAvatar: "https://images.gr-assets.com/authors/1305138891p2/4575371.jpg",
		Content: "I don't go looking for trouble. Trouble usually finds me.",
	},
	{
		Author: "Benjamin Franklin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215314094p2/289513.jpg",
		Content: "Either write something worth reading or do something worth writing.",
	},
	{
		Author: "William Shakespeare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "It is not in the stars to hold our destiny but in ourselves.",
	},
	{
		Author: "Saul Bellow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470607356p2/4391.jpg",
		Content: "You never have to change anything you got up in the middle of the night to write.",
	},
	{
		Author: "Sarah   Williams",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1400180303p2/333029.jpg",
		Content: "I have loved the stars too fondly to be fearful of the night.",
	},
	{
		Author: "Pablo NerudaLove",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "Love is so short, forgetting is so long.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "It is a truth universally acknowledged, that a single man in possession of a good fortune, must be in want of a wife.",
	},
	{
		Author: "Alice Hoffman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477318484p2/3502.jpg",
		Content: "Books may well be the only true magic.",
	},
	{
		Author: "William ShakespeareHamlet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "There is nothing either good or bad, but thinking makes it so.",
	},
	{
		Author: "Victor HugoLes Misérables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "He never went out without a book under his arm, and he often came back with two.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Good books don't give up all their secrets at once.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "If you can't fly then run, if you can't run then walk, if you can't walk then crawl, but whatever you do you have to keep moving forward.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "The world is indeed full of peril, and in it there are many dark places; but still there is much that is fair, and though in all lands love is now mingled with grief, it grows perhaps the greater.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "Once on a yellow piece of paper with green lines\the wrote a poemAnd he called it \"Chops\"\tbecause that was the name of his dogAnd that's what it was all aboutAnd his teacher gave him an A\tand a gold starAnd his mother hung it on the kitchen door\tand read it to his auntsThat was the year Father Tracy\ttook all the kids to the zooAnd he let them sing on the busAnd his little sister was born\twith tiny toenails and no hairAnd his mother and father kissed a lotAnd the girl around the corner sent him aValentine signed with a row of X's\tand he had to ask his father what the X's meantAnd his father always tucked him in bed at nightAnd was always there to do itOnce on a piece of white paper with blue lines\the wrote a poemAnd he called it \"Autumn\"\tbecause that was the name of the seasonAnd that's what it was all aboutAnd his teacher gave him an A\tand asked him to write more clearlyAnd his mother never hung it on the kitchen door\tbecause of its new paintAnd the kids told him\tthat Father Tracy smoked cigarsAnd left butts on the pewsAnd sometimes they would burn holesThat was the year his sister got glasses\twith thick lenses and black framesAnd the girl around the corner laughed\twhen he asked her to go see Santa ClausAnd the kids told him why\this mother and father kissed a lotAnd his father never tucked him in bed at nightAnd his father got mad\twhen he cried for him to do it.Once on a paper torn from his notebook\the wrote a poemAnd he called it \"Innocence: A Question\"\tbecause that was the question about his girlAnd that's what it was all aboutAnd his professor gave him an A\tand a strange steady lookAnd his mother never hung it on the kitchen door\tbecause he never showed herThat was the year that Father Tracy diedAnd he forgot how the end\tof the Apostle's Creed wentAnd he caught his sister\tmaking out on the back porchAnd his mother and father never kissed\tor even talkedAnd the girl around the corner\twore too much makeupThat made him cough when he kissed her\tbut he kissed her anyway\tbecause that was the thing to doAnd at three a.m. he tucked himself into bed\this father snoring soundlyThat's why on the back of a brown paper bag\the tried another poemAnd he called it \"Absolutely Nothing\"Because that's what it was really all aboutAnd he gave himself an A\tand a slash on each damned wristAnd he hung it on the bathroom door\tbecause this time he didn't think\the could reach the kitchen.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "We're all going to die, all of us, what a circus! That alone should make us love each other but it doesn't. We are terrorized and flattened by trivialities, we are eaten up by nothing.",
	},
	{
		Author: "Woody Allen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501608198p2/10356.jpg",
		Content: "I don't know the question, but sex is definitely the answer.",
	},
	{
		Author: "Marie LuLegend",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285032806p2/4342215.jpg",
		Content: "Each day means a new twenty-four hours. Each day means everything's possible again. You live in the moment, you die in the moment, you take it all one day at a time.",
	},
	{
		Author: "Daphne RaeLove Until It Hurts",
		AuthorAvatar: "",
		Content: "I have found the paradox, that if you love until it hurts, there can be no more hurt, only more love.",
	},
	{
		Author: "Jamie McGuireBeautiful Disaster",
		AuthorAvatar: "https://images.gr-assets.com/authors/1479315727p2/4464118.jpg",
		Content: "To douchebags!\" he said, gesturing to Brad. \"And to girls that break your heart,\" he bowed his head to me. His eyes lost focus. \"And to the absolute fucking horror of losing your best friend because you were stupid enough to fall in love with her.",
	},
	{
		Author: "Charles BukowskiFactotum",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "If you're going to try, go all the way. Otherwise, don't even start. This could mean losing girlfriends, wives, relatives and maybe even your mind. It could mean not eating for three or four days. It could mean freezing on a park bench. It could mean jail. It could mean derision. It could mean mockery--isolation. Isolation is the gift. All the others are a test of your endurance, of how much you really want to do it. And, you'll do it, despite rejection and the worst odds. And it will be better than anything else you can imagine. If you're going to try, go all the way. There is no other feeling like that. You will be alone with the gods, and the nights will flame with fire. You will ride life straight to perfect laughter. It's the only good fight there is.",
	},
	{
		Author: "Eleanor RooseveltYou Learn by Living: Eleven Keys for a More Fulfilling Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "You gain strength, courage and confidence by every experience in which you really stop to look fear in the face. You are able to say to yourself, 'I have lived through this horror. I can take the next thing that comes along.' You must do the thing you think you cannot do.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "I love to see a young girl go out and grab the world by the lapels. Life's a bitch. You've got to go out and kick ass.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Life is a disease: sexually transmitted, and invariably fatal.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Everything in the world is about sex except sex. Sex is about power.",
	},
	{
		Author: "David SedarisBarrel Fever: Stories and Essays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213737698p2/2849.jpg",
		Content: "If you're looking for sympathy you'll find it between shit and syphilis in the dictionary.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "There comes a time when the world gets quiet and the only thing left is your own heart. So you'd better learn the sound of it. Otherwise you'll never understand what it's saying.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You're still trying to protect me. Real or not real,\" he whispers.\"Real,\" I answer. \"Because that's what you and I do, protect each other.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Why fit in when you were born to stand out?",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "That does it,\" said Jace. \"I'm going to get you a dictionary for Christmas this year.\"\"Why?\" Isabelle said.\"So you can look up 'fun.' I'm not sure you know what it means.",
	},
	{
		Author: "Jonathan Safran Foer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "You cannot protect yourself from sadness without protecting yourself from happiness.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Fantasy is a necessary ingredient in living, it's a way of looking at life through the wrong end of a telescope.",
	},
	{
		Author: "Jane SmileyThirteen Ways of Looking at the Novel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1245100250p2/1339.jpg",
		Content: "Many people, myself among them, feel better at the mere sight of a book.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "They love their hair because they're not smart enough to love something more interesting.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "I was gratified to be able to answer promptly, and I did. I said I didn’t know.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Never allow someone to be your priority while allowing yourself to be their option.",
	},
	{
		Author: "William Faulkner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437805575p2/3535.jpg",
		Content: "Read, read, read. Read everything -- trash, classics, good and bad, and see how they do it. Just like a carpenter who works as an apprentice and studies the master. Read! You'll absorb it.Then write. If it's good, you'll find out. If it's not, throw it out of the window.",
	},
	{
		Author: "Nora Ephron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1366180104p2/5691.jpg",
		Content: "Above all, be the heroine of your life, not the victim.",
	},
	{
		Author: "Pamela Redmond Satran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1327444633p2/1736.jpg",
		Content: "A WOMAN SHOULD HAVE .... enough money within her control to move out and rent a place of her own even if she never wants to or needs to... A WOMAN SHOULD HAVE .... something perfect to wear if the employer or date of her dreams wants to see her in an hour... A WOMAN SHOULD HAVE ... a youth she's content to leave behind.... A WOMAN SHOULD HAVE .... a past juicy enough that she's looking forward to retelling it in her old age.... A WOMAN SHOULD HAVE ..... a set of screwdrivers, a cordless drill, and a black lace bra... A WOMAN SHOULD HAVE .... one friend who always makes her laugh... and one who lets her cry... A WOMAN SHOULD HAVE .... a good piece of furniture not previously owned by anyone else in her family... A WOMAN SHOULD HAVE .... eight matching plates, wine glasses with stems, and a recipe for a meal that will make her guests feel honored... A WOMAN SHOULD HAVE .... a feeling of control over her destiny... EVERY WOMAN SHOULD KNOW... how to fall in love without losing herself.. EVERY WOMAN SHOULD KNOW... HOW TO QUIT A JOB, BREAK UP WITH A LOVER, AND CONFRONT A FRIEND WITHOUT RUINING THE FRIENDSHIP... EVERY WOMAN SHOULD KNOW... when to try harder... and WHEN TO WALK AWAY... EVERY WOMAN SHOULD KNOW... that she can't change the length of her calves, the width of her hips, or the nature of her parents.. EVERY WOMAN SHOULD KNOW... that her childhood may not have been perfect...but it's over... EVERY WOMAN SHOULD KNOW... what she would and wouldn't do for love or more... EVERY WOMAN SHOULD KNOW... how to live alone... even if she doesn't like it... EVERY WOMAN SHOULD KNOW... whom she can trust, whom she can't, and why she shouldn't take it personally... EVERY WOMAN SHOULD KNOW... where to go... be it to her best friend's kitchen table... or a charming inn in the woods... when her soul needs soothing... EVERY WOMAN SHOULD KNOW... what she can and can't accomplish in a day... a month...and a year...",
	},
	{
		Author: "Arthur Conan DoyleThe Case-Book of Sherlock Holmes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495008883p2/2448.jpg",
		Content: "When you have eliminated all which is impossible, then whatever remains, however improbable, must be the truth.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "No tears in the writer, no tears in the reader. No surprise in the writer, no surprise in the reader.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Everyone should be able to do one card trick, tell two jokes, and recite three poems, in case they are ever trapped in an elevator.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Try not to become a man of success. Rather become a man of value.",
	},
	{
		Author: "Daphne du Maurier",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422444467p2/2001717.jpg",
		Content: "But luxury has never appealed to me, I like simple things, books, being alone, or with somebody who understands.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "As usual, there is a great woman behind every idiot.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Beneath the makeup and behind the smile I am just a girl who wishes for the world.",
	},
	{
		Author: "Khaled HosseiniThe Kite Runner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "For you, a thousand times over",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "A human being is a part of the whole called by us universe, a part limited in time and space. He experiences himself, his thoughts and feeling as something separated from the rest, a kind of optical delusion of his consciousness. This delusion is a kind of prison for us, restricting us to our personal desires and to affection for a few persons nearest to us. Our task must be to free ourselves from this prison by widening our circle of compassion to embrace all living creatures and the whole of nature in its beauty.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Man is least himself when he talks in his own person. Give him a mask, and he will tell you the truth.",
	},
	{
		Author: "Walt Disney Company",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289112683p2/3510823.jpg",
		Content: "Laughter is timeless. Imagination has no age. And dreams are forever.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "Men occasionally stumble over the truth, but most of them pick themselves up and hurry off as if nothing ever happened.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Courage is the most important of all the virtues because without courage, you can't practice any other virtue consistently.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Bran thought about it. 'Can a man still be brave if he's afraid?''That is the only time a man can be brave,' his father told him.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Being crazy isn't enough.",
	},
	{
		Author: "Oprah Winfrey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1354837955p2/3518.jpg",
		Content: "Turn your wounds into wisdom.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Kiss me, and you will see how important I am.",
	},
	{
		Author: "John SteinbeckThe Winter of Our Discontent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182118389p2/585.jpg",
		Content: "I wonder how many people I've looked at all my life and never seen.",
	},
	{
		Author: "Joseph Brodsky",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243291789p2/11563.jpg",
		Content: "There are worse crimes than burning books. One of them is not reading them.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Remember when you tried to convince me to feed a poultry pie to the mallards in the park to see if you could breed a race of cannibal ducks?\" \"They ate it too,\" Will reminisced. \"Bloodthirsty little beasts. Never trust a duck.",
	},
	{
		Author: "Isaac AsimovFoundation",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341965730p2/16667.jpg",
		Content: "Never let your sense of morals prevent you from doing what is right.",
	},
	{
		Author: "Oscar WildeThe Importance of Being Earnest and Other Plays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "The very essence of romance is uncertainty.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "The wound is the place where the Light enters you.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "There are few people whom I really love, and still fewer of whom I think well. The more I see of the world, the more am I dissatisfied with it; and every day confirms my belief of the inconsistency of all human characters, and of the little dependence that can be placed on the appearance of merit or sense.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I am very interested and fascinated how everyone loves each other, but no one really likes each other.",
	},
	{
		Author: "William ShakespeareThe Tempest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Hell is empty and all the devils are here.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Some tourists think Amsterdam is a city of sin, but in truth it is a city of freedom. And in freedom, most people find sin.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "There is always some madness in love. But there is also always some reason in madness.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The important thing is not to stop questioning. Curiosity has its own reason for existence. One cannot help but be in awe when he contemplates the mysteries of eternity, of life, of the marvelous structure of reality. It is enough if one tries merely to comprehend a little of this mystery each day.—\"Old Man's Advice to Youth: 'Never Lose a Holy Curiosity.'\" LIFE Magazine (2 May 1955) p. 64",
	},
	{
		Author: "Douglas CouplandShampoo Planet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1489283217p2/1886.jpg",
		Content: "Remember: the time you feel lonely is the time you most need to be by yourself. Life's cruelest irony.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Percy wouldn't notice a joke if it danced naked in front of him wearing one of Dobby's hats.",
	},
	{
		Author: "Haruki Murakami",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Whatever it is you're seeking won't come in the form you're expecting.",
	},
	{
		Author: "Edgar Allan PoeMarginalia",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "I have great faith in fools - self-confidence my friends will call it.",
	},
	{
		Author: "Niccolò MachiavelliThe Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1508764329p2/16201.jpg",
		Content: "Everyone sees what you appear to be, few experience what you really are.",
	},
	{
		Author: "Emily BrontëWuthering Heights",
		AuthorAvatar: "https://images.gr-assets.com/authors/1473229007p2/4191.jpg",
		Content: "If all else perished, and he remained, I should still continue to be; and if all else remained, and he were annihilated, the universe would turn to a mighty stranger.",
	},
	{
		Author: "Socrates",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390145726p2/275648.jpg",
		Content: "The unexamined life is not worth living.",
	},
	{
		Author: "Jhumpa LahiriThe Namesake",
		AuthorAvatar: "https://images.gr-assets.com/authors/1378932972p2/3670.jpg",
		Content: "That's the thing about books. They let you travel without moving your feet.",
	},
	{
		Author: "Gabriel García Márquez",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502310670p2/13450.jpg",
		Content: "What matters in life is not what happens to you but what you remember and how you remember it.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Deserves it! I daresay he does. Many that live deserve death. And some that die deserve life. Can you give it to them? Then do not be too eager to deal out death in judgement. For even the very wise cannot see all ends.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "Religion has actually convinced people that there's an invisible man living in the sky who watches everything you do, every minute of every day. And the invisible man has a special list of ten things he does not want you to do. And if you do any of these ten things, he has a special place, full of fire and smoke and burning and torture and anguish, where he will send you to live and suffer and burn and choke and scream and cry forever and ever 'til the end of time!But He loves you. He loves you, and He needs money! He always needs money! He's all-powerful, all-perfect, all-knowing, and all-wise, somehow just can't handle money!",
	},
	{
		Author: "William GoldmanWilliam Goldman: Four Screenplays with Essays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198704782p2/12521.jpg",
		Content: "Life is pain, highness. Anyone who says differently is selling something.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I have a theory that selflessness and bravery aren't all that different.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "A girl should be two things: classy and fabulous.",
	},
	{
		Author: "Anthony G. Oettinger",
		AuthorAvatar: "",
		Content: "Time flies like an arrow; fruit flies like a banana.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "But who prays for Satan? Who, in eighteen centuries, has had the common humanity to pray for the one sinner that needed it most?",
	},
	{
		Author: "Tahereh MafiShatter Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "I spent my life folded between the pages of books.In the absence of human relationships I formed bonds with paper characters. I lived love and loss through stories threaded in history; I experienced adolescence by association. My world is one interwoven web of words, stringing limb to limb, bone to sinew, thoughts and images all together. I am a being comprised of letters, a character created by sentences, a figment of imagination formed through fiction.",
	},
	{
		Author: "Arthur C. Clarke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357191481p2/7779.jpg",
		Content: "Two possibilities exist: either we are alone in the Universe or we are not. Both are equally terrifying.",
	},
	{
		Author: "Chuck PalahniukChoke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "What I want is to be needed. What I need is to be indispensable to somebody. Who I need is somebody that will eat up all my free time, my ego, my attention. Somebody addicted to me. A mutual addiction.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Stupid people are dangerous.",
	},
	{
		Author: "Irwin Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1491654461p2/7735.jpg",
		Content: "There are too many books I haven’t read, too many places I haven’t seen, too many memories I haven’t kept long enough.",
	},
	{
		Author: "Candace BushnellSex and the City",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435063350p2/4415.jpg",
		Content: "Man may have discovered fire, but women discovered how to play with it.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "I love people who make me laugh. I honestly think it's the thing I like most, to laugh. It cures a multitude of ills. It's probably the most important thing in a person.",
	},
	{
		Author: "Jim Morrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429658456p2/7855.jpg",
		Content: "People are afraid of themselves, of their own reality; their feelings most of all. People talk about how great love is, but that’s bullshit. Love hurts. Feelings are disturbing. People are taught that pain is evil and dangerous. How can they deal with love if they’re afraid to feel? Pain is meant to wake us up. People try to hide their pain. But they’re wrong. Pain is something to carry, like a radio. You feel your strength in the experience of pain. It’s all in how you carry it. That’s what matters. Pain is a feeling. Your feelings are a part of you. Your own reality. If you feel ashamed of them, and hide them, you’re letting society destroy your reality. You should stand up for your right to feel your pain.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "My daddy said, that the first time you fall in love, it changes you forever and no matter how hard you try, that feeling just never goes away.",
	},
	{
		Author: "Marilyn MonroeMarilyn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I don't mind living in a man's world, as long as I can be a woman in it.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "I like your Christ, I do not like your Christians. Your Christians are so unlike your Christ.",
	},
	{
		Author: "Ray BradburyZen in the Art of Writing",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "You must stay drunk on writing so reality cannot destroy you.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "Whoever fights monsters should see to it that in the process he does not become a monster. And if you gaze long enough into an abyss, the abyss will gaze back into you.",
	},
	{
		Author: "Ned VizziniIt's Kind of a Funny Story",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341737392p2/11672.jpg",
		Content: "I didn't want to wake up. I was having a much better time asleep. And that's really sad. It was almost like a reverse nightmare, like when you wake up from a nightmare you're so relieved. I woke up into a nightmare.",
	},
	{
		Author: "Christopher PaoliniEragon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1412963950p2/8349.jpg",
		Content: "Books are my friends, my companions. They make me laugh and cry and find meaning in life.",
	},
	{
		Author: "James Baldwin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343346341p2/10427.jpg",
		Content: "You think your pain and your heartbreak are unprecedented in the history of the world, but then you read. It was books that taught me that the things that tormented me most were the very things that connected me with all the people who were alive, who had ever been alive.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Nobody realizes that some people expend tremendous energy merely to be normal.",
	},
	{
		Author: "Paul AusterThe Brooklyn Follies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455603276p2/296961.jpg",
		Content: "Reading was my escape and my comfort, my consolation, my stimulant of choice: reading for the pure pleasure of it, for the beautiful stillness that surrounds you when you hear an author's words reverberating in your head.",
	},
	{
		Author: "Khaled HosseiniThe Kite Runner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "It may be unfair, but what happens in a few days, sometimes even a single day, can change the course of a whole lifetime...",
	},
	{
		Author: "Mark TwainThe Wit and Wisdom of Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The difference between the almost right word and the right word is really a large matter. ’tis the difference between the lightning bug and the lightning.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Those who find ugly meanings in beautiful things are corrupt without being charming. This is a fault. Those who find beautiful meanings in beautiful things are the cultivated. For these there is hope. They are the elect to whom beautiful things mean only Beauty. There is no such thing as a moral or an immoral book. Books are well written, or badly written. That is all.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "Nothing of me is original. I am the combined effort of everyone I've ever known.",
	},
	{
		Author: "Terry PratchettThe Last Continent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "It is said that your life flashes before your eyes just before you die. That is true, it's called Life.",
	},
	{
		Author: "John Green",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "…because nerds like us are allowed to be unironically enthusiastic about stuff… Nerds are allowed to love stuff, like jump-up-and-down-in-the-chair-can’t-control-yourself love it. Hank, when people call people nerds, mostly what they’re saying is ‘you like stuff.’ Which is just not a good insult at all. Like, ‘you are too enthusiastic about the miracle of human consciousness’.",
	},
	{
		Author: "Ray BradburyZen in the Art of Writing",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "I have never listened to anyone who criticized my taste in space travel, sideshows or gorillas. When this occurs, I pack up my dinosaurs and leave the room.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "The nicest thing for me is sleep, then at least I can dream.",
	},
	{
		Author: "Daniel Pennac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1236788892p2/40737.jpg",
		Content: "Reader's Bill of Rights1. The right to not read 2. The right to skip pages 3. The right to not finish 4. The right to reread 5. The right to read anything 6. The right to escapism 7. The right to read anywhere 8. The right to browse 9. The right to read out loud 10. The right to not defend your tastes",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "Some people feel the rain. Others just get wet.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "If you are going through hell, keep going.",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I am a badass, and I recognize that you, too, are a badass.",
	},
	{
		Author: "C.S. LewisMere Christianity",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Imagine yourself as a living house. God comes in to rebuild that house. At first, perhaps, you can understand what He is doing. He is getting the drains right and stopping the leaks in the roof and so on; you knew that those jobs needed doing and so you are not surprised. But presently He starts knocking the house about in a way that hurts abominably and does not seem to make any sense. What on earth is He up to? The explanation is that He is building quite a different house from the one you thought of - throwing out a new wing here, putting on an extra floor there, running up towers, making courtyards. You thought you were being made into a decent little cottage: but He is building a palace. He intends to come and live in it Himself.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "What I need is the dandelion in the spring. The bright yellow that means rebirth instead of destruction. The promise that life can go on, no matter how bad our losses. That it can be good again.",
	},
	{
		Author: "Paul McCartneyThe Beatles Illustrated Lyrics",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677530p2/20127.jpg",
		Content: "And, in the endThe love you takeis equal to the love you make.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Prayer is not asking. It is a longing of the soul. It is daily admission of one's weakness. It is better in prayer to have a heart without words than words without a heart.",
	},
	{
		Author: "S.E. HintonThe Outsiders",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206505616p2/762707.jpg",
		Content: "I lie to myself all the time. But I never believe me.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Peter would probably throw a party if I stopped breathing.''Well,' he says, 'I would only go if there was cake.",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Books are like mirrors: if a fool looks in, you cannot expect a genius to look out.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I don't want to lose the boy with the bread.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "Wise men speak because they have something to say; fools because they have to say something.",
	},
	{
		Author: "Henry RollinsThe Portable Henry Rollins",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209771874p2/44736.jpg",
		Content: "It hurts to let go. Sometimes it seems the harder you try to hold on to something or someone the more it wants to get away. You feel like some kind of criminal for having felt, for having wanted. For having wanted to be wanted. It confuses you, because you think that your feelings were wrong and it makes you feel so small because it's so hard to keep it inside when you let it out and it doesn't coma back. You're left so alone that you can't explain. Damn, there's nothing like that, is there? I've been there and you have too. You're nodding your head.",
	},
	{
		Author: "Pablo Picasso",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536109p2/3253.jpg",
		Content: "Every child is an artist. The problem is how to remain an artist once he grows up.",
	},
	{
		Author: "Charles BukowskiFactotum",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "My ambition is handicapped by laziness",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "But I don’t want to go among mad people,\" Alice remarked.\"Oh, you can’t help that,\" said the Cat: \"we’re all mad here. I’m mad. You’re mad.\"\"How do you know I’m mad?\" said Alice.\"You must be,\" said the Cat, \"or you wouldn’t have come here.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What is an \"instant\" death anyway? How long is an instant? Is it one second? Ten? The pain of those seconds must have been awful as her heart burst and her lungs collapsed and there was no air and no blood to her brain and only raw panic. What the hell is instant? Nothing is instant. Instant rice takes five minutes, instant pudding an hour. I doubt that an instant of blinding pain feels particularly instantaneous.",
	},
	{
		Author: "Neil GaimanA Game of You",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Everybody has a secret world inside of them. I mean everybody. All of the people in the whole world, I mean everybody — no matter how dull and boring they are on the outside. Inside them they've all got unimaginable, magnificent, wonderful, stupid, amazing worlds... Not just one world. Hundreds of them. Thousands, maybe.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "The town was paper, but the memories were not.",
	},
	{
		Author: "Carlos Ruiz ZafónThe Shadow of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444199853p2/815.jpg",
		Content: "Books are mirrors: you only see in them what you already have inside you.",
	},
	{
		Author: "Nelson Mandela",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308928296p2/367338.jpg",
		Content: "Education is the most powerful weapon which you can use to change the world.",
	},
	{
		Author: "William Faulkner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437805575p2/3535.jpg",
		Content: "Never be afraid to raise your voice for honesty and truth and compassion against injustice and lying and greed. If people all over the world...would do this, it would change the earth.",
	},
	{
		Author: "Bette Davis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270599807p2/56410.jpg",
		Content: "When a man gives his opinion, he's a man. When a woman gives her opinion, she's a bitch.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You haven't got a letter on yours,\" George observed. \"I suppose she thinks you don't forget your name. But we're not stupid-we know we're called Gred and Forge.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Stop acting so small. You are the universe in ecstatic motion.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I think God, in creating man, somewhat overestimated his ability.",
	},
	{
		Author: "Antoine de Saint-ExupéryAirman's Odyssey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "Love does not consist of gazing at each other, but in looking outward together in the same direction.",
	},
	{
		Author: "Charles BukowskiThe People Look Like Flowers at Last",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "You have to die a few times before you can reallylive.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It's a metaphor, see: You put the killing thing right between your teeth, but you don't give it the power to do its killing.",
	},
	{
		Author: "Theodore Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198519048p2/44567.jpg",
		Content: "It is not the critic who counts; not the man who points out how the strong man stumbles, or where the doer of deeds could have done them better. The credit belongs to the man who is actually in the arena, whose face is marred by dust and sweat and blood; who strives valiantly; who errs, who comes short again and again, because there is no effort without error and shortcoming; but who does actually strive to do the deeds; who knows great enthusiasms, the great devotions; who spends himself in a worthy cause; who at the best knows in the end the triumph of high achievement, and who at the worst, if he fails, at least fails while daring greatly, so that his place shall never be with those cold and timid souls who neither know victory nor defeat.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "The future belongs to those who believe in the beauty of their dreams.",
	},
	{
		Author: "Greg Behrendt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331400965p2/6842.jpg",
		Content: "If he’s not calling you, it’s because you are not on his mind. If he creates expectations for you, and then doesn’t follow through on little things, he will do same for big things. Be aware of this and realize that he’s okay with disappointing you. Don’t be with someone who doesn’t do what they say they’re going to do. If he’s choosing not to make a simple effort that would put you at ease and bring harmony to a recurring fight, then he doesn’t respect your feelings and needs. “Busy",
	},
	{
		Author: "Leonardo da Vinci",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399080979p2/13560.jpg",
		Content: "Painting is poetry that is seen rather than felt, and poetry is painting that is felt rather than seen.",
	},
	{
		Author: "Dale CarnegieHow to Win Friends and Influence People",
		AuthorAvatar: "https://images.gr-assets.com/authors/1230226725p2/3317.jpg",
		Content: "It isn't what you have or who you are or where you are or what you are doing that makes you happy or unhappy. It is what you think about it.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "My nightmares are usually about losing you. I'm okay once I realize you're here.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "The simple things are also the most extraordinary things, and only the wise can see them.",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "War is peace. Freedom is slavery. Ignorance is strength.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "My tastes are simple: I am easily satisfied with the best.",
	},
	{
		Author: "Gabriel García Márquez",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502310670p2/13450.jpg",
		Content: "No medicine cures what happiness cannot.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I saw my life branching out before me like the green fig tree in the story. From the tip of every branch, like a fat purple fig, a wonderful future beckoned and winked. One fig was a husband and a happy home and children, and another fig was a famous poet and another fig was a brilliant professor, and another fig was Ee Gee, the amazing editor, and another fig was Europe and Africa and South America, and another fig was Constantin and Socrates and Attila and a pack of other lovers with queer names and offbeat professions, and another fig was an Olympic lady crew champion, and beyond and above these figs were many more figs I couldn't quite make out. I saw myself sitting in the crotch of this fig tree, starving to death, just because I couldn't make up my mind which of the figs I would choose. I wanted each and every one of them, but choosing one meant losing all the rest, and, as I sat there, unable to decide, the figs began to wrinkle and go black, and, one by one, they plopped to the ground at my feet.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "I am not sure exactly what heaven will be like, but I know that when we die and it comes time for God to judge us, he will not ask, 'How many good things have you done in your life?' rather he will ask, 'How much love did you put into what you did?",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "You have your way. I have my way. As for the right way, the correct way, and the only way, it does not exist.",
	},
	{
		Author: "Fyodor DostoyevskyCrime and Punishment",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "Pain and suffering are always inevitable for a large intelligence and a deep heart. The really great men must, I think, have great sadness on earth.",
	},
	{
		Author: "Mortimer J. Adler",
		AuthorAvatar: "https://images.gr-assets.com/authors/1250032408p2/22395.jpg",
		Content: "In the case of good books, the point is not to see how many of them you can get through, but rather how many can get through to you.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Freedom is not worth having if it does not include the freedom to make mistakes.",
	},
	{
		Author: "Marcus Tullius Cicero",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197881416p2/13755.jpg",
		Content: "If you have a garden and a library, you have everything you need.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "What happens when people open their hearts?\"\"They get better.",
	},
	{
		Author: "Vladimir NabokovLolita",
		AuthorAvatar: "https://images.gr-assets.com/authors/1482502806p2/5152.jpg",
		Content: "It was love at first sight, at last sight, at ever and ever sight.",
	},
	{
		Author: "Walt WhitmanLeaves of Grass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392303683p2/1438.jpg",
		Content: "Resist much, obey little.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "You must have chaos within you to give birth to a dancing star.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "No book is really worth reading at the age of ten which is not equally – and often far more – worth reading at the age of fifty and beyond.",
	},
	{
		Author: "Andrew  BoydDaily Afflictions: The Agony of Being Connected to Everything in the Universe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1300810770p2/4732744.jpg",
		Content: "We’re all seeking that special person who is right for us. But if you’ve been through enough relationships, you begin to suspect there’s no right person, just different flavors of wrong. Why is this? Because you yourself are wrong in some way, and you seek out partners who are wrong in some complementary way. But it takes a lot of living to grow fully into your own wrongness. And it isn’t until you finally run up against your deepest demons, your unsolvable problems—the ones that make you truly who you are—that we’re ready to find a lifelong mate. Only then do you finally know what you’re looking for. You’re looking for the wrong person. But not just any wrong person: it's got to be the right wrong person—someone you lovingly gaze upon and think, “This is the problem I want to have.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Deadlines just aren't real to me until I'm staring one in the face.",
	},
	{
		Author: "George OrwellAnimal Farm",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "All animals are equal, but some animals are more equal than others.",
	},
	{
		Author: "Stephen KingOn Writing: A Memoir of the Craft",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "The road to hell is paved with adverbs.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You are the answer to every prayer I've offered. You are a song, a dream, a whisper, and I don't know how I could have lived without you for as long as I have.",
	},
	{
		Author: "Raymond Carver",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525695952p2/7363.jpg",
		Content: "Woke up this morning with a terrific urge to lie in bed all day and read.",
	},
	{
		Author: "Terry Pratchett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Stories of imagination tend to upset those without one.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Wrinkles should merely indicate where the smiles have been.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Fiction is the truth inside the lie.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Nobody can hurt me without my permission.",
	},
	{
		Author: "George Burns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1248450358p2/43259.jpg",
		Content: "Happiness is having a large, loving, caring, close-knit family in another city.",
	},
	{
		Author: "AnonymousHoly Bible: King James Version",
		AuthorAvatar: "",
		Content: "And now these three remain: faith, hope and love. But the greatest of these is love.",
	},
	{
		Author: "Douglas AdamsThe Salmon of Doubt: Hitchhiking the Galaxy One Last Time",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "The fact that we live at the bottom of a deep gravity well, on the surface of a gas covered planet going around a nuclear fireball 90 million miles away and think this to be normal is obviously some indication of how skewed our perspective tends to be.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "One of the Silent Brothers is here to see you. Hodge sent me to wake you up. Actually he offered to wake you himself, but since it's 5 a.m., I figured you'd be less cranky if you had something nice to look at.\"\"Meaning you?\"\"What else?",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Who said nights were for sleep?",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I can feel Peeta press his forehead into my temple and he asks, 'So now that you've got me, what are you going to do with me?' I turn into him. 'Put you somewhere you can't get hurt.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I've got a stele we can use. Who wants to do me?\"\"A regrettable choice of words,\" muttered Magnus.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "People haven't always been there for me but music always has.",
	},
	{
		Author: "William Faulkner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437805575p2/3535.jpg",
		Content: "You cannot swim for new horizons until you have courage to lose sight of the shore.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Fire is catching! And if we burn, you burn with us!",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "To be a Christian means to forgive the inexcusable because God has forgiven the inexcusable in you.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "It's strange because sometimes, I read a book, and I think I am the people in the book.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Don't order any of the faerie food,\" said Jace, looking at her over the top of his menu. \"It tends to make humans a little crazy. One minute you're munching a faerie plum, the next minute you're running naked down Madison Avenue with antlers on your head. Not,\" he added hastily, \"that this has ever happened to me.",
	},
	{
		Author: "Jack Kerouac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "Live, travel, adventure, bless, and don't be sorry.",
	},
	{
		Author: "Sigmund Freud",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406688955p2/10017.jpg",
		Content: "One day, in retrospect, the years of struggle will strike you as the most beautiful.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Loyalty to country ALWAYS. Loyalty to government, when it deserves it.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "To be nobody but yourself in a world which is doing its best day and night to make you like everybody else means to fight the hardest battle which any human being can fight and never stop fighting.",
	},
	{
		Author: "Joan Crawford",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211859541p2/76857.jpg",
		Content: "Love is a fire. But whether it is going to warm your hearth or burn down your house, you can never tell.",
	},
	{
		Author: "F. Scott Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "That is part of the beauty of all literature. You discover that your longings are universal longings, that you're not lonely and isolated from anyone. You belong.",
	},
	{
		Author: "Haruki MurakamiKafka on the Shore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "If you remember me, then I don't care if everyone else forgets.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Will looked horrified. \"What kind of monster could possibly hate chocolate?",
	},
	{
		Author: "Agatha Christie",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321738793p2/123715.jpg",
		Content: "I like living. I have sometimes been wildly, despairingly, acutely miserable, racked with sorrow; but through it all I still know quite certainly that just to be alive is a grand thing.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Maybe 'okay' will be our 'always",
	},
	{
		Author: "William ShakespeareTwelfth Night",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "If music be the food of love, play on,Give me excess of it; that surfeiting,The appetite may sicken, and so die.",
	},
	{
		Author: "Kinky Friedman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282510465p2/32112.jpg",
		Content: "My dear,Find what you love and let it kill you.Let it drain you of your all. Let it cling onto your back and weigh you down into eventual nothingness.Let it kill you and let it devour your remains.For all things will kill you, both slowly and fastly, but it’s much better to be killed by a lover.~ Falsely yours",
	},
	{
		Author: "F. Scott Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "Show me a hero, and I'll write you a tragedy.",
	},
	{
		Author: "Bill Watterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "It's not denial. I'm just selective about the reality I accept.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "I don't want to die without any scars.",
	},
	{
		Author: "Oscar WildeThe Importance of Being Earnest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I never travel without my diary. One should always have something sensational to read in the train.",
	},
	{
		Author: "Ursula K. Le GuinThe Lathe of Heaven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244291425p2/874602.jpg",
		Content: "Love doesn't just sit there, like a stone, it has to be made, like bread; remade all the time, made new.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Sometimes crying or laughing are the only options left, and laughing feels better right now.",
	},
	{
		Author: "Mark TwainFollowing the Equator: A Journey Around the World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Truth is stranger than fiction, but it is because Fiction is obliged to stick to possibilities; Truth isn't.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "I don't trust people who don't love themselves and tell me, 'I love you.' ... There is an African saying which is: Be careful when a naked person offers you a shirt.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Should I kill myself, or have a cup of coffee?",
	},
	{
		Author: "Olin Miller",
		AuthorAvatar: "",
		Content: "You probably wouldn’t worry about what people think of you if you could know how seldom they do.",
	},
	{
		Author: "Gayle FormanIf I Stay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1472502468p2/295178.jpg",
		Content: "Sometimes you make choices in life and sometimes choices make you.",
	},
	{
		Author: "Louisa May AlcottWork: A Story of Experience",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200326665p2/1315.jpg",
		Content: "She is too fond of books, and it has turned her brain.",
	},
	{
		Author: "Lewis CarrollAlice's Adventures in Wonderland \u0026 Through the Looking-Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "She generally gave herself very good advice, (though she very seldom followed it).",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "It's just that I don't want to be somebody's crush. If somebody likes me, I want them to like the real me, not what they think I am. And I don't want them to carry it around inside. I want them to show me, so I can feel it too.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I have something I need to tell you,\" he says. I run my fingers along the tendons in his hands and look back at him. \"I might be in love with you.\" He smiles a little. \"I'm waiting until I'm sure to tell you, though.\"\"That's sensible of you,\" I say, smiling too. \"We should find some paper so you can make a list or a chart or something.\"I feel his laughter against my side, his nose sliding along my jaw, his lips pressing my ear.\"Maybe I'm already sure,\" he says, \"and I just don't want to frighten you.\"I laugh a little. \"Then you should know better.\"\"Fine,\" he says. \"Then I love you.",
	},
	{
		Author: "Anne RiceThe Vampire Lestat",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383250078p2/7577.jpg",
		Content: "None of us really changes over time. We only become more fully what we are.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Fear doesn't shut you down; it wakes you up",
	},
	{
		Author: "Margaret Mitchell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1185481392p2/11081.jpg",
		Content: "Life's under no obligation to give us what we expect.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I may die young, but at least I'll die smart.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Don't talk to me.\"\"Why not?\"\"Because I want to fix that in my memory for ever. Draco Malfoy, the amazing bouncing ferret...",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You know, Minister, I disagree with Dumbledore on many counts...but you cannot deny he's got style...",
	},
	{
		Author: "Helen Keller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1266096039p2/7275.jpg",
		Content: "The best and most beautiful things in the world cannot be seen or even touched. They must be felt with the heart",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "None but ourselves can free our minds.",
	},
	{
		Author: "Muhammad Ali",
		AuthorAvatar: "https://images.gr-assets.com/authors/1323184538p2/46261.jpg",
		Content: "Friendship is the hardest thing in the world to explain. It's not something you learn in school. But if you haven't learned the meaning of friendship, you really haven't learned anything.",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "We do not need magic to transform our world. We carry all the power we need inside ourselves already.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "The most important things are the hardest to say. They are the things you get ashamed of, because words diminish them -- words shrink things that seemed limitless when they were in your head to no more than living size when they're brought out. But it's more than that, isn't it? The most important things lie too close to wherever your secret heart is buried, like landmarks to a treasure your enemies would love to steal away. And you may make revelations that cost you dearly only to have people look at you in a funny way, not understanding what you've said at all, or why you thought it was so important that you almost cried while you were saying it. That's the worst, I think. When the secret stays locked within not for want of a teller but for want of an understanding ear.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Finnick?\" I say, \"Maybe some pants?\"He looks down at his legs as if noticing his outfit for the first time. Then he whips off his hospital gown leaving him in just his underwear. \"Why? Do you find this\" -- he strikes a ridiculously provocative pose -- \"distracting?\"I laugh. Boggs looks embarrassed and Finnick looks more like the guy I met at the Quarter Quell",
	},
	{
		Author: "Dan BrownDigital Fortress",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399396714p2/630.jpg",
		Content: "Everything is possible. The impossible just takes longer.",
	},
	{
		Author: "Robert Louis Stevenson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192746024p2/854076.jpg",
		Content: "Don't judge each day by the harvest you reap but by the seeds that you plant.",
	},
	{
		Author: "Confucius",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407613261p2/15321.jpg",
		Content: "By three methods we may learn wisdom: First, by reflection, which is noblest; Second, by imitation, which is easiest; and third by experience, which is the bitterest.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Like most misery, it started with apparent happiness.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Love conquers all,\" Aphrodite promised. \"Look at Helen and Paris. Did they let anything come between them?\"\"Didn't they start the Trojan War and get thousands of people killed?\"\"Pfft. That's not the point. Follow your heart.",
	},
	{
		Author: "Anne Tyler",
		AuthorAvatar: "https://images.gr-assets.com/authors/1332881249p2/457.jpg",
		Content: "I read so I can live more than one life in more than one place.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Wow,\" Thalia muttered. \"Apollo is hot.\" \"He's the sun god,\" I said.\"That's not what I meant.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Ever had a flying burrito hit you? Well, it's a deadly projectile, right up there with cannonballs and grenades.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Can I help you with something?\"Clary turned instant traitor against her gender. \"Those girls on the other side of the car are staring at you.\"Jace assumed an air of mellow gratification. \"Of course they are,\" he said, \"I am stunningly attractive.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "But I know, somehow, that only when it is dark enough can you see the stars.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "Never say goodbye because goodbye means going away and going away means forgetting.",
	},
	{
		Author: "Tupac Shakur",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208454439p2/87764.jpg",
		Content: "You can spend minutes, hours, days, weeks, or even months over-analyzing a situation; trying to put the pieces together, justifying what could've, would've happened... or you can just leave the pieces on the floor and move the fuck on.",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "If you love somebody, let them go, for if they return, they were always yours. If they don't, they never were.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Why are they all staring?\" demanded Albus as he and Rose craned around to look at the other students.\"Don’t let it worry you,\" said Ron. \"It’s me. I’m extremely famous.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You guessed? You must have been pretty sure, considering you could have killed me.\"\"I was ninety percent sure.\"\"I see,\" Clary said. There must have been something in her voice, because he turned to look at her. Her hand cracked across his face, a slap that rocked him back on his heels. He put his hands on his cheek, more in surprise than pain.\"What the hell was that for?\"\"The other ten percent.",
	},
	{
		Author: "Rudyard Kipling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183237590p2/6989.jpg",
		Content: "Words are, of course, the most powerful drug used by mankind.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "Angry people are not always wise.",
	},
	{
		Author: "Neil GaimanAmerican Gods",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "What I say is, a town isn’t a town without a bookstore. It may call itself a town, but unless it’s got a bookstore, it knows it’s not foolin’ a soul.",
	},
	{
		Author: "Alexander Pope",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206647570p2/25157.jpg",
		Content: "Blessed is he who expects nothing, for he shall never be disappointed.",
	},
	{
		Author: "Paulo CoelhoBy the River Piedra I Sat Down and Wept",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "If pain must come, may it come quickly. Because I have a life to live, and I need to live it in the best way possible. If he has to make a choice, may he make it now. Then I will either wait for him or forget him.",
	},
	{
		Author: "Cathy Guisewite",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213820999p2/19058.jpg",
		Content: "When life gives you lemons, squirt someone in the eye.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I don’t know if you’ve ever felt like that. That you wanted to sleep for a thousand years. Or just not exist. Or just not be aware that you do exist. Or something like that. I think wanting that is very morbid, but I want it when I get like this. That’s why I’m trying not to think. I just want it all to stop spinning.",
	},
	{
		Author: "Timothy Leary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429657521p2/47718.jpg",
		Content: "Women who seek to be equal with men lack ambition.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "I was never really insane except upon occasions when my heart was touched.",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "Why, sometimes I've believed as many as six impossible things before breakfast.",
	},
	{
		Author: "Anton Chekhov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1317162641p2/5031025.jpg",
		Content: "Don't tell me the moon is shining; show me the glint of light on broken glass.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "You could rattle the stars,\" she whispered. \"You could do anything, if only you dared. And deep down, you know it, too. That’s what scares you most.",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "I am coming to terms with the fact that loving someone requires a leap of faith, and that a soft landing is never guaranteed.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "A brave man acknowledges the strength of others.",
	},
	{
		Author: "Richelle MeadBlood Promise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "I’d said it before and meant it: Alive or undead, the love of my life was a badass.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Once you can accept the universe as matter expanding into nothing that is something, wearing stripes with plaid comes easy.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "He must have known I'd want to leave you.\"\"No, he must have known you would always want to come back.",
	},
	{
		Author: "William ShakespeareRomeo and Juliet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "When he shall die,Take him and cut him out in little stars,And he will make the face of heaven so fineThat all the world will be in love with nightAnd pay no worship to the garish sun.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Let us step into the night and pursue that flighty temptress, adventure.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Fear cuts deeper than swords.",
	},
	{
		Author: "Kurt VonnegutSlaughterhouse-Five",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "Everything was beautiful and nothing hurt.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "The problem with the world is that the intelligent people are full of doubts, while the stupid ones are full of confidence.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Music was my refuge. I could crawl into the space between the notes and curl my back to loneliness.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Be careful of love. It'll twist your brain around and leave you thinking up is down and right is wrong.",
	},
	{
		Author: "Shel Silverstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "How many slams in an old screen door? Depends how loud you shut it. How many slices in a bread? Depends how thin you cut it. How much good inside a day? Depends how good you live 'em. How much love inside a friend? Depends how much you give 'em.",
	},
	{
		Author: "Flannery O'Connor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1469878767p2/22694.jpg",
		Content: "The truth does not change according to our ability to stomach it.",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Some things don't last forever, but some things do. Like a good song, or a good book, or a good memory you can take out and unfold in your darkest times, pressing down on the corners and peering in close, hoping you still recognize the person you see there.",
	},
	{
		Author: "Nicholas SparksMessage in a Bottle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "If you like her, if she makes you happy, and if you feel like you know her---then don't let her go.",
	},
	{
		Author: "Albert Schweitzer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198534412p2/47146.jpg",
		Content: "Sometimes our light goes out, but is blown again into instant flame by an encounter with another human being.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It always shocked me when I realized that I wasn’t the only person in the world who thought and felt such strange and awful things.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "I do not fear death. I had been dead for billions and billions of years before I was born, and had not suffered the slightest inconvenience from it.",
	},
	{
		Author: "J.R.R. TolkienThe Lord of the Rings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "It's a dangerous business, Frodo, going out your door. You step onto the road, and if you don't keep your feet, there's no knowing where you might be swept off to.",
	},
	{
		Author: "F. Scott FitzgeraldThis Side of Paradise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I don't want to repeat my innocence. I want the pleasure of losing it again.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Mom says it's because she has PMS.Do you even know what that means?\"I'm not a little kid anymore. It means pissed-at- men syndrome",
	},
	{
		Author: "Margaret Atwood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "War is what happens when language fails.",
	},
	{
		Author: "Patrick RothfussThe Wise Man's Fear",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351307341p2/108424.jpg",
		Content: "There are three things all wise men fear: the sea in storm, a night with no moon, and the anger of a gentle man.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "God has no religion.",
	},
	{
		Author: "Umberto Eco",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455915753p2/1730.jpg",
		Content: "We live for books.",
	},
	{
		Author: "Virginia WoolfBetween the Acts",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "Books are the mirrors of the soul.",
	},
	{
		Author: "Terry PratchettA Hat Full of Sky",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Why do you go away? So that you can come back. So that you can see the place you came from with new eyes and extra colors. And the people there see you differently, too. Coming back to where you started is not the same as never leaving.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I cannot fix on the hour, or the spot, or the look or the words, which laid the foundation. It is too long ago. I was in the middle before I knew that I had begun.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Investigation?\" Isabelle laughed. \"Now we're detectives? Maybe we should all have code names.\"\"Good idea,\" said Jace. \"I shall be Baron Hotschaft Von Hugenstein.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Music is a total constant. That's why we have such a strong visceral connection to it, you know? Because a song can take you back instantly to a moment, or a place, or even a person. No matter what else has changed in your or the world, that one song says the same, just like that moment.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "You smell good, too,",
	},
	{
		Author: "Haruki MurakamiSputnik Sweetheart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Why do people have to be this lonely? What's the point of it all? Millions of people in this world, all of them yearning, looking to others to satisfy them, yet isolating themselves. Why? Was the earth put here just to nourish human loneliness?",
	},
	{
		Author: "John GreenWill Grayson, Will Grayson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Some people have lives; some people have music.",
	},
	{
		Author: "Joseph Campbell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114498p2/20105.jpg",
		Content: "Life has no meaning. Each of us has meaning and we bring it to life. It is a waste to be asking the question when you are the answer.",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "Begin at the beginning,\" the King said, very gravely, \"and go on till you come to the end: then stop.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "If you don't imagine, nothing ever happens at all.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "It kills me sometimes, how people die.",
	},
	{
		Author: "Herman Melville",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495029910p2/1624.jpg",
		Content: "It is better to fail in originality than to succeed in imitation.",
	},
	{
		Author: "Charles DickensA Christmas Carol",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "There is nothing in the world so irresistibly contagious as laughter and good humor.",
	},
	{
		Author: "Martin Luther King Jr.A Knock at Midnight: Inspiration from the Great Sermons of Reverend Martin Luther King, Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Let no man pull you so low as to hate him.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "People generally see what they look for, and hear what they listen for.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "I’m here. I love you. I don’t care if you need to stay up crying all night long, I will stay with you. If you need the medication again, go ahead and take it—I will love you through that, as well. If you don’t need the medication, I will love you, too. There’s nothing you can ever do to lose my love. I will protect you until you die, and after your death I will still protect you. I am stronger than Depression and I am braver than Loneliness and nothing will ever exhaust me.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Everyone seems to have a clear idea of how other people should lead their lives, but none about his or her own.",
	},
	{
		Author: "Pedro Calderón de la Barca",
		AuthorAvatar: "https://images.gr-assets.com/authors/1227042355p2/227771.jpg",
		Content: "When love is not madness it is not love.",
	},
	{
		Author: "Benjamin Spock",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206592906p2/35038.jpg",
		Content: "Trust yourself. You know more than you think you do.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "Unbeing dead isn't being alive.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Fear of a name increases fear of the thing itself.",
	},
	{
		Author: "Richelle MeadVampire Academy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Only a true best friend can protect you from your immortal enemies.",
	},
	{
		Author: "Anonymous",
		AuthorAvatar: "",
		Content: "A lie gets halfway around the world before the truth has a chance to get its pants on.",
	},
	{
		Author: "Lois LowryThe Giver",
		AuthorAvatar: "https://images.gr-assets.com/authors/1348162077p2/2493.jpg",
		Content: "The worst part of holding the memories is not the pain. It's the loneliness of it. Memories need to be shared.",
	},
	{
		Author: "Anna QuindlenHow Reading Changed My Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387414577p2/3500.jpg",
		Content: "Books are the plane, and the train, and the road. They are the destination, and the journey. They are home.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "Angry, and half in love with her, and tremendously sorry, I turned away.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "It is more fun to talk with someone who doesn't use long, difficult words but rather short, easy words like \"What about lunch?",
	},
	{
		Author: "James Baldwin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343346341p2/10427.jpg",
		Content: "Love does not begin and end the way we seem to think it does. Love is a battle, love is a war; love is a growing up.",
	},
	{
		Author: "Stephanie KleinStraight Up and Dirty",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204924878p2/80145.jpg",
		Content: "Tell the truth, or someone will tell it for you.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I love you more than there are stars in the sky and fish in the sea.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "If you look for perfection, you'll never be content.",
	},
	{
		Author: "Richelle MeadVampire Academy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Do you think I'm pretty?I think you're beautifulBeautiful?You are so beautiful, it hurts sometimes.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Declarations of love amuse me. Especially when unrequited.",
	},
	{
		Author: "Charles BukowskiWomen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "That's the problem with drinking, I thought, as I poured myself a drink. If something bad happens you drink in an attempt to forget; if something good happens you drink in order to celebrate; and if nothing happens you drink to make something happen.",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "I love you. Remember. They cannot take it",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Weeds are flowers, too, once you get to know them.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "What you seek is seeking you.",
	},
	{
		Author: "Marcel Proust",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392271688p2/233619.jpg",
		Content: "Let us be grateful to the people who make us happy; they are the charming gardeners who make our souls blossom.",
	},
	{
		Author: "Oscar WildeThe Canterville Ghost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Death must be so beautiful. To lie in the soft brown earth, with the grasses waving above one's head, and listen to silence. To have no yesterday, and no tomorrow. To forget time, to forgive life, to be at peace.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What you must understand about me is that I’m a deeply unhappy person.",
	},
	{
		Author: "Khaled HosseiniAnd the Mountains Echoed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "I suspect the truth is that we are waiting, all of us, against insurmountable odds, for something extraordinary to happen to us.",
	},
	{
		Author: "Max Lucado",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493056226p2/2737.jpg",
		Content: "A woman's heart should be so hidden in God that a man has to seek Him just to find her.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "What's this?\" he demanded, looking from Clary to his companions, as if they might know what she was doing there.\"It's a girl,\" Jace said,recovering his composure. \"Surely you've seen girls before, Alec. Your sister Isabelle is one.",
	},
	{
		Author: "Patrick NessA Monster Calls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244216486p2/370361.jpg",
		Content: "You do not write your life with words...You write it with actions. What you think is not important. It is only important what you do.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "The only way to find true happiness is to risk being completely cut open.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The world as we have created it is a process of our thinking. It cannot be changed without changing our thinking.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "People think dreams aren't real just because they aren't made of matter, of particles. Dreams are real. But they are made of viewpoints, of images, of memories and puns and lost hopes.",
	},
	{
		Author: "Walt Disney Company",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289112683p2/3510823.jpg",
		Content: "It's kind of fun to do the impossible.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It's all right to love someone who doesn't love you back, as long as they're worth you loving them. As long as they deserve it.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Think left and think right and think low and think high. Oh, the thinks you can think up if only you try!",
	},
	{
		Author: "Lemony SnicketThe Penultimate Peril",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "I suppose I'll have to add the force of gravity to my list of enemies.",
	},
	{
		Author: "Nicholas SparksMessage in a Bottle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "True love is rare, and it's the only thing that gives life real meaning.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Cinderella? Snow White? What's that? An illness?",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "We are not necessarily doubting that God will do the best for us; we are wondering how painful the best will turn out to be.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "If I were not a physicist, I would probably be a musician. I often think in music. I live my daydreams in music. I see my life in terms of music.",
	},
	{
		Author: "Gillian Anderson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1486469860p2/193300.jpg",
		Content: "Well, it seems to me that the best relationships - the ones that last - are frequently the ones that are rooted in friendship. You know, one day you look at the person and you see something more than you did the night before. Like a switch has been flicked somewhere. And the person who was just a friend is... suddenly the only person you can ever imagine yourself with.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "This moment will just be another story someday.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "People aren't born good or bad. Maybe they're born with tendencies either way, but its the way you live your life that matters.",
	},
	{
		Author: "Anaïs NinThe Diary of Anaïs Nin, Vol. 1: 1931-1934",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Each friend represents a world in us, a world possibly not born until they arrive, and it is only by this meeting that a new world is born.",
	},
	{
		Author: "Maya AngelouLetter to My Daughter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "You may not control all the events that happen to you, but you can decide not to be reduced by them.",
	},
	{
		Author: "Alice Walker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406752585p2/7380.jpg",
		Content: "No person is your friend who demands your silence, or denies your right to grow.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "I am haunted by humans.",
	},
	{
		Author: "Leo TolstoyThe Kreutzer Sonata",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "It is amazing how complete is the delusion that beauty is goodness.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "I did then what I knew how to do. Now that I know better, I do better.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You know,\" Gabriel said, \"there was a time I thought we could be friends, Will.\"\"There was a time I thought I was a ferret,\" Will said, \"but that turned out to be the opium haze. Did you know it had that effect? Because I didn't.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "When I was a little girl I used to read fairy tales. In fairy tales you meet Prince Charming and he's everything you ever wanted. In fairy tales the bad guy is very easy to spot. The bad guy is always wearing a black cape so you always know who he is. Then you grow up and you realize that Prince Charming is not as easy to find as you thought. You realize the bad guy is not wearing a black cape and he's not easy to spot; he's really funny, and he makes you laugh, and he has perfect hair.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Nobody likes being alone that much. I don't go out of my way to make friends, that's all. It just leads to disappointment. ",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "Enjoy it. Because it's happening.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I believe the universe wants to be noticed. I think the universe is inprobably biased toward the consciousness, that it rewards intelligence in part because the universe enjoys its elegance being observed. And who am I, living in the middle of history, to tell the universe that it-or my observation of it-is temporary?",
	},
	{
		Author: "Charles BukowskiPost Office",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "I wanted the whole world or nothing.",
	},
	{
		Author: "Sarah J. MaasA Court of Mist and Fury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "To the people who look at the stars and wish, Rhys.\"Rhys clinked his glass against mine. “To the stars who listen— and the dreams that are answered.",
	},
	{
		Author: "Annie Proulx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1219720509p2/1262010.jpg",
		Content: "You know, one of the tragedies of real life is that there is no background music.",
	},
	{
		Author: "Pablo Picasso",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536109p2/3253.jpg",
		Content: "Art is the lie that enables us to realize the truth.",
	},
	{
		Author: "E.L. JamesFifty Shades of Grey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308409727p2/4725841.jpg",
		Content: "Laters, baby.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "I have heard there are troubles of more than one kind. Some come from ahead and some come from behind. But I've bought a big bat. I'm all ready you see. Now my troubles are going to have troubles with me!",
	},
	{
		Author: "Shel SilversteinEvery Thing on It",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "There are no happy endings.Endings are the saddest part, So just give me a happy middleAnd a very happy start.",
	},
	{
		Author: "Lao TzuTao Te Ching",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "Simplicity, patience, compassion.These three are your greatest treasures.Simple in actions and thoughts, you return to the source of being.Patient with both friends and enemies,you accord with the way things are.Compassionate toward yourself,you reconcile all beings in the world.",
	},
	{
		Author: "Langston HughesThe Collected Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206669966p2/36910.jpg",
		Content: "Life is for the living.Death is for the dead.Let life be like music. And death a note unsaid.",
	},
	{
		Author: "Rabindranath TagoreStray Birds",
		AuthorAvatar: "https://images.gr-assets.com/authors/1453892068p2/36913.jpg",
		Content: "Clouds come floating into my life, no longer to carry rain or usher storm, but to add color to my sunset sky.",
	},
	{
		Author: "Edward Gorey",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1418090334p2/204145.jpg",
		Content: "Books. Cats. Life is Good.",
	},
	{
		Author: "Elbert Hubbard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216826209p2/114059.jpg",
		Content: "God will not look you over for medals, degrees or diplomas but for scars.",
	},
	{
		Author: "Frank Zappa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1315160559p2/22302.jpg",
		Content: "Without deviation from the norm, progress is not possible.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Some cause happiness wherever they go; others whenever they go.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I mean, if the relationship can't survive the long term, why on earth would it be worth my time and energy for the short term?",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Albus Severus,\" Harry said quietly, so that nobody but Ginny could hear, and she was tactful enough to pretend to be waving to Rose, who was now on the train, \"you were named for two headmasters of Hogwarts. One of them was a Slytherin and he was probably the bravest man I ever knew.",
	},
	{
		Author: "Douglas AdamsLife, the Universe and Everything",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "The Guide says there is an art to flying\", said Ford, \"or rather a knack. The knack lies in learning how to throw yourself at the ground and miss.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "That’s part of what I like about the book in some ways. It portrays death truthfully. You die in the middle of your life, in the middle of a sentence",
	},
	{
		Author: "Leonardo da Vinci",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399080979p2/13560.jpg",
		Content: "A painter should begin every canvas with a wash of black, because all things in nature are dark except where exposed by the light.",
	},
	{
		Author: "Lao Tzu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "The journey of a thousand miles begins with a single step.",
	},
	{
		Author: "Diane DuaneSo You Want to Be a Wizard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285664395p2/11761.jpg",
		Content: "Reading one book is like eating one potato chip.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "With freedom, books, flowers, and the moon, who could not be happy?",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Dumbledore says people find it far easier to forgive others for being wrong than being right.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "It's just that...I just think that some things are meant to be broken. Imperfect. Chaotic. It's the universe's way of providing contrast, you know? There have to be a few holes in the road. It's how life is.",
	},
	{
		Author: "P.G. Wodehouse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198684105p2/7963.jpg",
		Content: "There is no surer foundation for a beautiful friendship than a mutual taste in literature.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "You believe lies so you eventually learn to trust no one but yourself.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Don't.\" Clary raised a warning hand. \"I'm not really in the mood right now.\"\"That's got to be the first time a girl's ever said that to me,\" Jace mused.",
	},
	{
		Author: "Lance ArmstrongEvery Second Counts",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213137491p2/1544.jpg",
		Content: "Pain is temporary. Quitting lasts forever.",
	},
	{
		Author: "Steve MaraboliUnapologetically You: Reflections on Life and the Human Experience",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "The truth is, unless you let go, unless you forgive yourself, unless you forgive the situation, unless you realize that the situation is over, you cannot move forward.",
	},
	{
		Author: "James PattersonSaving the World and Other Extreme Sports",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468347205p2/3780.jpg",
		Content: "Basically, I have two speeds.... Hostile or smart-aleck. Your choice.",
	},
	{
		Author: "John Waters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263397316p2/17366.jpg",
		Content: "If you go home with somebody, and they don't have books, don't fuck 'em!",
	},
	{
		Author: "Laurell K. HamiltonMistral's Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399387919p2/9550.jpg",
		Content: "There are wounds that never show on the body that are deeper and more hurtful than anything that bleeds.",
	},
	{
		Author: "Elisabeth Kübler-Ross",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282171501p2/1506.jpg",
		Content: "The most beautiful people we have known are those who have known defeat, known suffering, known struggle, known loss, and have found their way out of the depths. These persons have an appreciation, a sensitivity, and an understanding of life that fills them with compassion, gentleness, and a deep loving concern. Beautiful people do not just happen.",
	},
	{
		Author: "Kahlil GibranThe Prophet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "Let there be spaces in your togetherness, And let the winds of the heavens dance between you. Love one another but make not a bond of love: Let it rather be a moving sea between the shores of your souls. Fill each other's cup but drink not from one cup. Give one another of your bread but eat not from the same loaf. Sing and dance together and be joyous, but let each one of you be alone, Even as the strings of a lute are alone though they quiver with the same music. Give your hearts, but not into each other's keeping. For only the hand of Life can contain your hearts. And stand together, yet not too near together: For the pillars of the temple stand apart, And the oak tree and the cypress grow not in each other's shadow.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "I know not with what weapons World War III will be fought, but World War IV will be fought with sticks and stones.",
	},
	{
		Author: "Hunter S. ThompsonThe Proud Highway: Saga of a Desperate Southern Gentleman, 1955-1967",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206560814p2/5237.jpg",
		Content: "Life should not be a journey to the grave with the intention of arriving safely in a pretty and well preserved body, but rather to skid in broadside in a cloud of smoke, thoroughly used up, totally worn out, and loudly proclaiming \"Wow! What a Ride!",
	},
	{
		Author: "Corrie ten BoomClippings from My Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217173231p2/102203.jpg",
		Content: "Worry does not empty tomorrow of its sorrow, it empties today of its strength.",
	},
	{
		Author: "Martin Luther King Jr.A Testament of Hope: The Essential Writings and Speeches",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "There comes a time when one must take a position that is neither safe, nor politic, nor popular, but he must take it because conscience tells him it is right.",
	},
	{
		Author: "J.R.R. TolkienThe Return of the King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "I will not say: do not weep; for not all tears are an evil.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "People say nothing is impossible, but I do nothing every day.",
	},
	{
		Author: "Aldous HuxleyBrave New World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "Words can be like X-rays if you use them properly -- they’ll go through anything. You read and you’re pierced.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "You’re wishin’ too much, baby. You gotta stop wearing your wishbone where your backbone oughtta be.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "I think we dream so we don’t have to be apart for so long. If we’re in each other’s dreams, we can be together all the time.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You’ve got about as much charm as a dead slug.",
	},
	{
		Author: "Elie Wiesel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1255518412p2/1049.jpg",
		Content: "There may be times when we are powerless to prevent injustice, but there must never be a time when we fail to protest.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "Sometimes people don't want to hear the truth because they don't want their illusions destroyed.",
	},
	{
		Author: "Jeanette Winterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406177070p2/9399.jpg",
		Content: "Book collecting is an obsession, an occupation, a disease, an addiction, a fascination, an absurdity, a fate. It is not a hobby. Those who do it must do it.",
	},
	{
		Author: "Vladimir NabokovLolita",
		AuthorAvatar: "https://images.gr-assets.com/authors/1482502806p2/5152.jpg",
		Content: "And the rest is rust and stardust.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "If writers wrote as carelessly as some people talk, then adhasdh asdglaseuyt[bn[ pasdlgkhasdfasdf.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Books are for people who wish they were somewhere else.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Francois Rabelais. He was a poet. And his last words were \"I go to seek a Great Perhaps.\" That's why I'm going. So I don't have to wait until I die to start seeking a Great Perhaps.",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "Would you tell me, please, which way I ought to go from here?\"\"That depends a good deal on where you want to get to.\"\"I don't much care where –\"\"Then it doesn't matter which way you go.",
	},
	{
		Author: "Orhan PamukThe New Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1423469681p2/1728.jpg",
		Content: "I read a book one day and my whole life was changed.",
	},
	{
		Author: "Terry PratchettLords and Ladies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "If cats looked like frogs we'd realize what nasty, cruel little bastards they are. Style. That's what people remember.",
	},
	{
		Author: "Allen Ginsberg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1421583811p2/4261.jpg",
		Content: "Follow your inner moonlight; don't hide the madness.",
	},
	{
		Author: "Jack KerouacOn the Road: the Original Scroll",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "There was nowhere to go but everywhere, so just keep on rolling under the stars.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Love is an irresistible desire to be irresistibly desired.",
	},
	{
		Author: "David Foster Wallace",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466019433p2/4339.jpg",
		Content: "The so-called ‘psychotically depressed’ person who tries to kill herself doesn’t do so out of quote ‘hopelessness’ or any abstract conviction that life’s assets and debits do not square. And surely not because death seems suddenly appealing. The person in whom Its invisible agony reaches a certain unendurable level will kill herself the same way a trapped person will eventually jump from the window of a burning high-rise. Make no mistake about people who leap from burning windows. Their terror of falling from a great height is still just as great as it would be for you or me standing speculatively at the same window just checking out the view; i.e. the fear of falling remains a constant. The variable here is the other terror, the fire’s flames: when the flames get close enough, falling to death becomes the slightly less terrible of two terrors. It’s not desiring the fall; it’s terror of the flames. And yet nobody down on the sidewalk, looking up and yelling ‘Don’t!’ and ‘Hang on!’, can understand the jump. Not really. You’d have to have personally been trapped and felt flames to really understand a terror way beyond falling.",
	},
	{
		Author: "William Shakespeare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "We know what we are, but not what we may be.",
	},
	{
		Author: "Lou Holtz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1310021362p2/85179.jpg",
		Content: "It's not the load that breaks you down, it's the way you carry it.",
	},
	{
		Author: "Erica Jong",
		AuthorAvatar: "https://images.gr-assets.com/authors/1286537368p2/6085.jpg",
		Content: "Everyone has talent. What's rare is the courage to follow it to the dark places where it leads.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "How do you feel, Georgie?\" whispered Mrs. Weasley.George's fingers groped for the side of his head.\"Saintlike,\" he murmured.\"What's wrong with him?\" croaked Fred, looking terrified. \"Is his mind affected?\"\"Saintlike,\" repeated George, opening his eyes and looking up at his brother. \"You see...I'm HOLEY, Fred, geddit?",
	},
	{
		Author: "Kalu Ndukwe Kalu",
		AuthorAvatar: "",
		Content: "The things you do for yourself are gone when you are gone, but the things you do for others remain as your legacy.",
	},
	{
		Author: "E.L. Doctorow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387414707p2/12584.jpg",
		Content: "Writing is a socially acceptable form of schizophrenia.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "Do you think I am an automaton? — a machine without feelings? and can bear to have my morsel of bread snatched from my lips, and my drop of living water dashed from my cup? Do you think, because I am poor, obscure, plain, and little, I am soulless and heartless? You think wrong! — I have as much soul as you — and full as much heart! And if God had gifted me with some beauty and much wealth, I should have made it as hard for you to leave me, as it is now for me to leave you. I am not talking to you now through the medium of custom, conventionalities, nor even of mortal flesh: it is my spirit that addresses your spirit; just as if both had passed through the grave, and we stood at God's feet, equal — as we are!",
	},
	{
		Author: "أحلام مستغانمي",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351458215p2/1109606.jpg",
		Content: "أجمل حب هو الذي نعثر عليه أثناء بحثنا عن شيء آخر",
	},
	{
		Author: "Paulo CoelhoEleven Minutes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Everything tells me that I am about to make a wrong decision, but making mistakes is just part of life. What does the world want of me? Does it want me to take no risks, to go back to where I came from because I didn't have the courage to say \"yes\" to life?",
	},
	{
		Author: "Vincent van Gogh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1329405187p2/34583.jpg",
		Content: "It is good to love many things, for therein lies the true strength, and whosoever loves much performs much, and can accomplish much, and what is done in love is well done.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You realize that trying to keep your distance from me will not lessen my affection for you. All efforts to save me from you will fail.",
	},
	{
		Author: "Henry David ThoreauWalden",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "Rather than love, than money, than fame, give me truth.",
	},
	{
		Author: "Jim Morrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429658456p2/7855.jpg",
		Content: "The most important kind of freedom is to be what you really are. You trade in your reality for a role. You trade in your sense for an act. You give up your ability to feel, and in exchange, put on a mask. There can't be any large-scale revolution until there's a personal revolution, on an individual level. It's got to happen inside first.",
	},
	{
		Author: "Santosh KalwarQuote Me Everyday",
		AuthorAvatar: "https://images.gr-assets.com/authors/1275309637p2/2894169.jpg",
		Content: "I was smiling yesterday,I am smiling today and I will smile tomorrow.Simply because life is too short to cry for anything.",
	},
	{
		Author: "Steven Wright",
		AuthorAvatar: "https://images.gr-assets.com/authors/1347482080p2/181771.jpg",
		Content: "Right now I’m having amnesia and déjà vu at the same time. I think I’ve forgotten this before.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Just because you call an electric eel a rubber duck doesn't make it a rubber duck, does it? And God help the poor bastard who decides they want to take a bath with the duckie. (Jace Wayland)",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You are so busy being YOU that you have no idea how utterly unprecedented you are.",
	},
	{
		Author: "D.H. LawrenceLady Chatterley's Lover",
		AuthorAvatar: "https://images.gr-assets.com/authors/1375811585p2/17623.jpg",
		Content: "A woman has to live her life, or live to repent not having lived it.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "She wasn't bitter. She was sad, though. But it was a hopeful kind of sad. The kind of sad that just takes time. ",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Mom. I have something to tell you. I’m undead. Now, I know you may have some preconceived notions about the undead. I know you may not be comfortable with the idea of me being undead. But I’m here to tell you that undead are just like you and me … well, okay. Possibly more like me than you.",
	},
	{
		Author: "David Foster WallaceInfinite Jest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466019433p2/4339.jpg",
		Content: "Everybody is identical in their secret unspoken belief that way deep down they are different from everyone else.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Holey? You have the the whole world of ear-related humor before you, you go for holey?",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "The best way to find out if you can trust somebody is to trust them.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You see, cuckoos are parasites. They lay their eggs in other birds' nests. When the egg hatches, the baby cuckoo pushes the other baby birds out of the nest. The poor parent birds work themselves to death trying to find enough food to feed the enormous cuckoo child who has murdered their babies and taken their places.\"\"Enormous?\" said Jace. \"Did you just call me fat?\"\"It was an analogy.\"\"I am not fat.",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "I don't believe in the kind of magic in my books. But I do believe something very magical can happen when you read a good book.",
	},
	{
		Author: "Cornelia FunkeInkspell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437000100p2/15873.jpg",
		Content: "Stories never really end...even if the books like to pretend they do. Stories always go on. They don't end on the last page, any more than they begin on the first page.",
	},
	{
		Author: "مصطفى محمودالقرآن: محاولة لفهم عصري",
		AuthorAvatar: "https://images.gr-assets.com/authors/1458565195p2/871358.jpg",
		Content: "لن تكون متدينا إلا بالعلم ...فالله لا يعبد بالجهل",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "The more I see, the less I know for sure.",
	},
	{
		Author: "Bob MooreheadWords Aptly Spoken",
		AuthorAvatar: "",
		Content: "The paradox of our time in history is that we have taller buildings but shorter tempers, wider Freeways, but narrower viewpoints. We spend more, but have less, we buy more, but enjoy less. We have bigger houses and smaller families, more conveniences, but less time. We have more degrees but less sense, more knowledge, but less judgment, more experts, yet more problems, more medicine, but less wellness.We drink too much, smoke too much, spend too recklessly, laugh too little, drive too fast, get too angry, stay up too late, get up too tired, read too little, watch TV too much, and pray too seldom. We have multiplied our possessions, but reduced our values. We talk too much, love too seldom, and hate too often.We've learned how to make a living, but not a life. We've added years to life not life to years. We've been all the way to the moon and back, but have trouble crossing the street to meet a new neighbor. We conquered outer space but not inner space. We've done larger things, but not better things.We've cleaned up the air, but polluted the soul. We've conquered the atom, but not our prejudice. We write more, but learn less. We plan more, but accomplish less. We've learned to rush, but not to wait. We build more computers to hold more information, to produce more copies than ever, but we communicate less and less.These are the times of fast foods and slow digestion, big men and small character, steep profits and shallow relationships.These are the days of two incomes but more divorce, fancier houses, but broken homes. These are days of quick trips, disposable diapers, throwaway morality, one night stands, overweight bodies, and pills that do everything from cheer, to quiet, to kill. It is a time when there is much in the showroom window and nothing in the stockroom. A time when technology can bring this letter to you, and a time when you can choose either to share this insight, or to just hit delete...Remember, to spend some time with your loved ones, because they are not going to be around forever. Remember, say a kind word to someone who looks up to you in awe, because that little person soon will grow up and leave your side.Remember, to give a warm hug to the one next to you, because that is the only treasure you can give with your heart and it doesn't cost a cent.Remember, to say, \"I love you\" to your partner and your loved ones, but most of all mean it. A kiss and an embrace will mend hurt when it comes from deep inside of you.Remember to hold hands and cherish the moment for someday that person might not be there again. Give time to love, give time to speak! And give time to share the precious thoughts in your mind.",
	},
	{
		Author: "John F. Kennedy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1498639086p2/3047.jpg",
		Content: "The Chinese use two brush strokes to write the word 'crisis.' One brush stroke stands for danger; the other for opportunity. In a crisis, be aware of the danger--but recognize the opportunity.",
	},
	{
		Author: "David MitchellCloud Atlas",
		AuthorAvatar: "https://images.gr-assets.com/authors/1409248688p2/6538289.jpg",
		Content: "My life amounts to no more than one drop in a limitless ocean. Yet what is any ocean, but a multitude of drops?",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Some old wounds never truly heal, and bleed again at the slightest word.",
	},
	{
		Author: "Charles BukowskiLove Is a Dog from Hell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "there is a loneliness in this world so greatthat you can see it in the slow movement ofthe hands of a clock.people so tiredmutilatedeither by love or no love.people just are not good to each otherone on one.the rich are not good to the richthe poor are not good to the poor.we are afraid.our educational system tells usthat we can all bebig-ass winners.it hasn't told usabout the guttersor the suicides.or the terror of one personaching in one placealoneuntouchedunspoken towatering a plant.",
	},
	{
		Author: "Patrick RothfussThe Name of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351307341p2/108424.jpg",
		Content: "It's like everyone tells a story about themselves inside their own head. Always. All the time. That story makes you what you are. We build ourselves out of that story.",
	},
	{
		Author: "Katharine Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1332828172p2/84099.jpg",
		Content: "If you obey all of the rules, you miss all of the fun.",
	},
	{
		Author: "Francis of AssisiThe Little Flowers of St. Francis of Assisi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1347231992p2/149151.jpg",
		Content: "All the darkness in the world cannot extinguish the light of a single candle.",
	},
	{
		Author: "L. Frank BaumThe Lost Princess of Oz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383720421p2/3242.jpg",
		Content: "No thief, however skillful, can rob one of knowledge, and that is why knowledge is the best and safest treasure to acquire.",
	},
	{
		Author: "Gordon B. Hinckley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202224343p2/313356.jpg",
		Content: "Life is to be enjoyed, not endured",
	},
	{
		Author: "RumiThe Illuminated Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "The minute I heard my first love story,I started looking for you, not knowinghow blind that was.Lovers don't finally meet somewhere.They're in each other all along.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Now, you two – this year, you behave yourselves. If I get one more owl telling me you've – you've blown up a toilet or –\"\"Blown up a toilet? We've never blown up a toilet.\"\"Great idea though, thanks, Mum.",
	},
	{
		Author: "Steve MaraboliLife, the Truth, and Being Free",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "Dare to BeWhen a new day begins, dare to smile gratefully.When there is darkness, dare to be the first to shine a light.When there is injustice, dare to be the first to condemn it.When something seems difficult, dare to do it anyway.When life seems to beat you down, dare to fight back.When there seems to be no hope, dare to find some.When you’re feeling tired, dare to keep going.When times are tough, dare to be tougher.When love hurts you, dare to love again.When someone is hurting, dare to help them heal.When another is lost, dare to help them find the way.When a friend falls, dare to be the first to extend a hand.When you cross paths with another, dare to make them smile.When you feel great, dare to help someone else feel great too.When the day has ended, dare to feel as you’ve done your best.Dare to be the best you can –At all times, Dare to be!",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "You can't be happy unless you're unhappy sometimes\".",
	},
	{
		Author: "Jodi PicoultMy Sister's Keeper",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "Maybe who we are isn't so much about what we do, but rather what we're capable of when we least expect it.",
	},
	{
		Author: "H.G. Wells",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515507862p2/880695.jpg",
		Content: "We all have our time machines, don't we. Those that take us back are memories...And those that carry us forward, are dreams.",
	},
	{
		Author: "Henry James",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468309415p2/159.jpg",
		Content: "Three things in human life are important: the first is to be kind; the second is to be kind; and the third is to be kind.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "I think... if it is true that there are as many minds as there are heads, then there are as many kinds of love as there are hearts.",
	},
	{
		Author: "Nicolas Chamfort",
		AuthorAvatar: "https://images.gr-assets.com/authors/1420504765p2/308646.jpg",
		Content: "The most wasted of all days is one without laughter.",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "All grown-ups were once children... but only few of them remember it.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "A poem begins as a lump in the throat, a sense of wrong, a homesickness, a lovesickness.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "In heaven, all the interesting people are missing.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "You never fail until you stop trying.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "What would men be without women? Scarce, sir...mighty scarce.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "Vanity and pride are different things, though the words are often used synonymously. A person may be proud without being vain. Pride relates more to our opinion of ourselves, vanity to what we would have others think of us.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Success is liking yourself, liking what you do, and liking how you do it.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Out of the corner of her eye she thought she saw Jace shoot her a look of white rage - but when she glanced at him, he looked as he always did: easy, confident, slightly bored.\"In future, Clarissa,\" he said, \"it might be wise to mention that you already have a man in your bed, to avoid such tedious situations.\"\"You invited him into bed?\" Simon demanded, looking shaken.\"Ridiculous, isn't it?\" said Jace. \"We would never have all fit.\"\"I didn't invite him into bed,\" Clary snapped. \"We were just kissing.\"\"Just kissing?\" Jace's tone mocked her with its false hurt. \"How swiftly you dismiss our love.",
	},
	{
		Author: "Mitch AlbomThe Five People You Meet In Heaven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "All endings are also beginnings. We just don't know it at the time.",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Usually I'm remarkably good natured. Try me on a day that doesn't end in y.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "A good library will never be too neat, or too dusty, because somebody will always be in it, taking books off the shelves and staying up late reading them.",
	},
	{
		Author: "Fyodor DostoyevskyCrime and Punishment",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "To go wrong in one's own way is better than to go right in someone else's.",
	},
	{
		Author: "Gregory MaguireWicked: The Life and Times of the Wicked Witch of the West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1319068553p2/7025.jpg",
		Content: "People who claim that they're evil are usually no worse than the rest of us... It's people who claim that they're good, or any way better than the rest of us, that you have to be wary of.",
	},
	{
		Author: "Bertrand Russell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457021442p2/17854.jpg",
		Content: "Do not fear to be eccentric in opinion, for every opinion now accepted was once eccentric.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Both Rowling and Meyer, they’re speaking directly to young people. … The real difference is that Jo Rowling is a terrific writer and Stephenie Meyer can’t write worth a darn. She’s not very good.",
	},
	{
		Author: "Albert EinsteinThe World as I See It",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The most beautiful experience we can have is the mysterious. It is the fundamental emotion that stands at the cradle of true art and true science.",
	},
	{
		Author: "Jonathan Safran Foer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "I love you also means I love you more than anyone loves you, or has loved you, or will love you, and also, I love you in a way that no one loves you, or has loved you, or will love you, and also, I love you in a way that I love no one else, and never have loved anyone else, and never will love anyone else.",
	},
	{
		Author: "Emily Dickinson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536260p2/7440.jpg",
		Content: "If I can stop one heart from breaking, I shall not live in vain.",
	},
	{
		Author: "G.K. Chesterton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1365860649p2/7014283.jpg",
		Content: "The true soldier fights not because he hates what is in front of him, but because he loves what is behind him.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "If you want to forget something or someone, never hate it, or never hate him/her. Everything and everyone that you hate is engraved upon your heart; if you want to let go of something, if you want to forget, you cannot hate.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Don’t grieve. Anything you lose comes round in another form.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "Do I not destroy my enemies when I make them my friends?",
	},
	{
		Author: "Louis L'Amour",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343675199p2/858.jpg",
		Content: "Start writing, no matter what. The water does not flow until the faucet is turned on.",
	},
	{
		Author: "Alfred Tennyson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454788521p2/13638502.jpg",
		Content: "HopeSmiles from the threshold of the year to come, Whispering 'it will be happier'...",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "There was a clatter as the basilisk fangs cascaded out of Hermione's arms. Running at Ron, she flung them around his neck and kissed him full on the mouth. Ron threw away the fangs and broomstick he was holding and responded with such enthusiasm that he lifted Hermione off her feet. \"Is this the moment?\" Harry asked weakly, and when nothing happened except that Ron and Hermione gripped each other still more firmly and swayed on the spot, he raised his voice. \"OI! There's a war going on here!\" Ron and Hermione broke apart, their arms still around each other. \"I know, mate,\" said Ron, who looked as though he had recently been hit on the back of the head with a Bludger, \"so it's now or never, isn't it?\" \"Never mind that, what about the Horcrux?\" Harry shouted. \"D'you think you could just --- just hold it in, until we've got the diadem?\" \"Yeah --- right --- sorry ---\" said Ron, and he and Hermione set about gathering up fangs, both pink in the face.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "Don't be afraid of your fears. They're not there to scare you. They're there to let you know that something is worth it.",
	},
	{
		Author: "Jimi Hendrix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207679011p2/7268.jpg",
		Content: "When the power of love overcomes the love of power, the world will know peace.",
	},
	{
		Author: "Donald MillerA Million Miles in a Thousand Years: What I Learned While Editing My Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198545955p2/4829.jpg",
		Content: "When you stop expecting people to be perfect, you can like them for who they are.",
	},
	{
		Author: "Stephen KingOn Writing: A Memoir of the Craft",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "The scariest moment is always just before you start.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The reason it hurts so much to separate is because our souls are connected.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "No, I'm just a very naughty boy. I do all sorts of bad things. I kick kittens. I make rude gestures at nuns.",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "Why didn't I learn to treat everything like it was the last time. My greatest regret was how much I believed in the future.",
	},
	{
		Author: "Dr. SeussOh, The Places You'll Go!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "You're off to Great Places!Today is your day!Your mountain is waiting,So... get on your way!",
	},
	{
		Author: "Ray BradburyFahrenheit 451",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "Why is it,\" he said, one time, at the subway entrance, \"I feel I've known you so many years?\"\"Because I like you,\" she said, \"and I don't want anything from you.",
	},
	{
		Author: "Agatha ChristieAn Autobiography",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321738793p2/123715.jpg",
		Content: "It is a curious thought, but it is only when you see people looking ridiculous that you realize just how much you love them. ",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I'm a grenade and at some point I'm going to blow up and I would like to minimize the casualties, okay?",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "There's plenty of sense in nonsense sometimes, if you wish to look for it.",
	},
	{
		Author: "Jay AsherThirteen Reasons Why",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243931536p2/569269.jpg",
		Content: "You don’t know what goes on in anyone’s life but your own. And when you mess with one part of a person’s life, you’re not messing with just that part. Unfortunately, you can’t be that precise and selective. When you mess with one part of a person’s life, you’re messing with their entire life. Everything. . . affects everything.",
	},
	{
		Author: "Pearl S. Buck",
		AuthorAvatar: "https://images.gr-assets.com/authors/1344969427p2/704.jpg",
		Content: "Many people lose the small joys in the hope for the big happiness.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You'll stay with me?'Until the very end,' said James.",
	},
	{
		Author: "Jane AustenPersuasion",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I hate to hear you talk about all women as if they were fine ladies instead of rational creatures. None of us want to be in calm waters all our lives.",
	},
	{
		Author: "Jane Austen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "but for my own part, if a book is well written, I always find it too short.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Some say the world will end in fire,Some say in ice.From what I've tasted of desire,I hold with those who favor fire. But if it had to perish twiceI think I know enough of hateTo say that for destruction iceIs also greatAnd would suffice.",
	},
	{
		Author: "Veronica RothInsurgent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "People, I have discovered, are layers and layers of secrets. You believe you know them, that you understand them, but their motives are always hidden from you, buried in their own hearts. You will never know them, but sometimes you decide to trust them.",
	},
	{
		Author: "John GreenWill Grayson, Will Grayson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Maybe there's something you're afraid to say, or someone you're afraid to love, or somewhere you're afraid to go. It's gonna hurt. It's gonna hurt because it matters.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Great spirits have always encountered violent opposition from mediocre minds.",
	},
	{
		Author: "Charles DickensGreat Expectations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "I loved her against reason, against promise, against peace, against hope, against happiness, against all discouragement that could be.",
	},
	{
		Author: "Bukkyo Dendo KyokaiThe Teaching of Buddha",
		AuthorAvatar: "",
		Content: "The secret of health for both mind and body is not to mourn for the past, nor to worry about the future, but to live the present moment wisely and earnestly.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I have the choice of being constantly active and happy or introspectively passive and sad. Or I can go mad by ricocheting in between.",
	},
	{
		Author: "H. Jackson Brown Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288560924p2/33394.jpg",
		Content: "Don't say you don't have enough time. You have exactly the same number of hours per day that were given to Helen Keller, Pasteur, Michelangelo, Mother Teresa, Leonardo da Vinci, Thomas Jefferson, and Albert Einstein.",
	},
	{
		Author: "Lao TzuTao Te Ching",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "Knowing others is intelligence;knowing yourself is true wisdom.Mastering others is strength; mastering yourself is true power.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Hate the sin, love the sinner.",
	},
	{
		Author: "G.K. ChestertonAlarms and Discursions",
		AuthorAvatar: "https://images.gr-assets.com/authors/1365860649p2/7014283.jpg",
		Content: "Poets have been mysteriously silent on the subject of cheese.",
	},
	{
		Author: "Philip Pullman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1396622492p2/3618.jpg",
		Content: "After nourishment, shelter and companionship, stories are the thing we need most in the world.",
	},
	{
		Author: "Amy Bloom",
		AuthorAvatar: "https://images.gr-assets.com/authors/1205260949p2/115220.jpg",
		Content: "You are imperfect, permanently and inevitably flawed. And you are beautiful.",
	},
	{
		Author: "Madeleine L'Engle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1305256804p2/106.jpg",
		Content: "A book, too, can be a star, a living fire to lighten the darkness, leading out into the expanding universe.",
	},
	{
		Author: "Frank Herbert",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168661521p2/58.jpg",
		Content: "There is no real ending. It’s just the place where you stop the story.",
	},
	{
		Author: "Lemony SnicketThe Wide Window",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "If you are allergic to a thing, it is best not to put that thing in your mouth, particularly if the thing is cats.",
	},
	{
		Author: "Hermann Hesse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499981916p2/1113469.jpg",
		Content: "Learn what is to be taken seriously and laugh at the rest.",
	},
	{
		Author: "Gillian FlynnGone Girl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232123231p2/2383.jpg",
		Content: "Men always say that as the defining compliment, don’t they? She’s a cool girl. Being the Cool Girl means I am a hot, brilliant, funny woman who adores football, poker, dirty jokes, and burping, who plays video games, drinks cheap beer, loves threesomes and anal sex, and jams hot dogs and hamburgers into her mouth like she’s hosting the world’s biggest culinary gang bang while somehow maintaining a size 2, because Cool Girls are above all hot. Hot and understanding. Cool Girls never get angry; they only smile in a chagrined, loving manner and let their men do whatever they want. Go ahead, shit on me, I don’t mind, I’m the Cool Girl.Men actually think this girl exists. Maybe they’re fooled because so many women are willing to pretend to be this girl. For a long time Cool Girl offended me. I used to see men – friends, coworkers, strangers – giddy over these awful pretender women, and I’d want to sit these men down and calmly say: You are not dating a woman, you are dating a woman who has watched too many movies written by socially awkward men who’d like to believe that this kind of woman exists and might kiss them. I’d want to grab the poor guy by his lapels or messenger bag and say: The bitch doesn’t really love chili dogs that much – no one loves chili dogs that much! And the Cool Girls are even more pathetic: They’re not even pretending to be the woman they want to be, they’re pretending to be the woman a man wants them to be. Oh, and if you’re not a Cool Girl, I beg you not to believe that your man doesn’t want the Cool Girl. It may be a slightly different version – maybe he’s a vegetarian, so Cool Girl loves seitan and is great with dogs; or maybe he’s a hipster artist, so Cool Girl is a tattooed, bespectacled nerd who loves comics. There are variations to the window dressing, but believe me, he wants Cool Girl, who is basically the girl who likes every fucking thing he likes and doesn’t ever complain. (How do you know you’re not Cool Girl? Because he says things like: “I like strong women.",
	},
	{
		Author: "Bertrand RussellThe Conquest of Happiness",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457021442p2/17854.jpg",
		Content: "Of all forms of caution, caution in love is perhaps the most fatal to true happiness.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Don't Panic.",
	},
	{
		Author: "Nicholas SparksNights in Rodanthe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The greater the love, the greater the tragedy when it's over.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Because you are beautiful. I enjoy looking at beautiful people, and I decided a while ago not to deny myself the simpler pleasures of existence",
	},
	{
		Author: "François Rabelais",
		AuthorAvatar: "https://images.gr-assets.com/authors/1223897787p2/11029.jpg",
		Content: "I go to seek a Great Perhaps.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "My mission in life is not merely to survive, but to thrive; and to do so with some passion, some compassion, some humor, and some style",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "An intellectual says a simple thing in a hard way. An artist says a hard thing in a simple way.",
	},
	{
		Author: "Jane AustenPersuasion",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "You pierce my soul. I am half agony, half hope...I have loved none but you.",
	},
	{
		Author: "F. Scott Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I fell in love with her courage, her sincerity, and her flaming self respect. And it's these things I'd believe in, even if the whole world indulged in wild suspicions that she wasn't all she should be. I love her and it is the beginning of everything.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Sanity and happiness are an impossible combination.",
	},
	{
		Author: "Scott WesterfeldUglies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1442207392p2/13957.jpg",
		Content: "What you do, the way you think, makes you beautiful.",
	},
	{
		Author: "Rainer Maria Rilke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493785350p2/7906.jpg",
		Content: "Be patient toward all that is unsolved in your heart and try to love the questions themselves, like locked rooms and like books that are now written in a very foreign tongue. Do not now seek the answers, which cannot be given you because you would not be able to live them. And the point is, to live everything. Live the questions now. Perhaps you will then gradually, without noticing it, live along some distant day into the answer.",
	},
	{
		Author: "Rose Fitzgerald Kennedy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1462387083p2/86892.jpg",
		Content: "It has been said, 'time heals all wounds.' I do not agree. The wounds remain. In time, the mind, protecting its sanity, covers them with scar tissue and the pain lessens. But it is never gone.",
	},
	{
		Author: "Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Perhaps when we find ourselves wanting everything, it is because we are dangerously close to wanting nothing.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Fantasy is escapist, and that is its glory. If a soldier is imprisioned by the enemy, don't we consider it his duty to escape?. . .If we value the freedom of mind and soul, if we're partisans of liberty, then it's our plain duty to escape, and to take as many people with us as we can!",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "He's a wallflower. You see things. You keep quiet about them. And you understand.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Is standing by the window muttering about blood something he does all the time?\" asked Simon.\"No,\" Jace said. \"Sometimes he sits on the couch and does it.",
	},
	{
		Author: "S.E. Hinton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206505616p2/762707.jpg",
		Content: "If you have two friends in your lifetime, you're lucky. If you have one good friend, you're more than lucky.",
	},
	{
		Author: "Yann MartelLife of Pi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454012903p2/811.jpg",
		Content: "It is true that those we meet can change us, sometimes so profoundly that we are not the same afterwards, even unto our names.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It's not because I want to make out with her.\"Hold on.\" He grabbed a pencil and scrawled excitedly at the paper as if he'd just made a mathematical breakthrough and then looked back up at me. \"I just did some calculations, and I've been able to determine that you're full of shit",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I was thinking about the first time I ever saw you,\" he said, \"and how after that I couldn't forget you. I wanted to, but I couldn't stop myself. I forced Hodge to let me be the one who came to find you and bring you back to the Institue. And even back then, in that stupid coffee shop, when I saw you sitting on that couch with Simon, even then that felt wrong to me-- I should have been the one sitting with you. The one who made you laugh like that. I couldn't get rid of that feeling. That it should have been me. And the more I knew you, the more I felt it--it had never been like that for me before. I'd always wanted a girl and then gotten to know her and not wanted her anymore, but with you the feeling just got stronger and stronger until that night when you showed up at Renwick's and I knew.",
	},
	{
		Author: "Sarah Dessen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Sometimes it seems safer to hold it all in, where the only person who can judge is yourself.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1504718077p2/53259.jpg",
		Content: "Among other things, you'll find that you're not the first person who was ever confused and frightened and even sickened by human behavior. You're by no means alone on that score, you'll be excited and stimulated to know. Many, many men have been just as troubled morally and spiritually as you are right now. Happily, some of them kept records of their troubles. You'll learn from them—if you want to. Just as someday, if you have something to offer, someone will learn something from you. It's a beautiful reciprocal arrangement. And it isn't education. It's history. It's poetry.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Ah,",
	},
	{
		Author: "Thich Nhat HanhPeace Is Every Step: The Path of Mindfulness in Everyday Life",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1421193042p2/299262.jpg",
		Content: "Walk as if you are kissing the Earth with your feet.",
	},
	{
		Author: "Paulo CoelhoThe Zahir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "When someone leaves, it's because someone else is about to arrive.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "History will be kind to me for I intend to write it.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "One love, one heart, one destiny.",
	},
	{
		Author: "William BlakeAuguries of Innocence",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199069675p2/13453.jpg",
		Content: "To see a World in a Grain of Sand And a Heaven in a Wild Flower, Hold Infinity in the palm of your hand And Eternity in an hour.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "All we demanded was our right to twinkle.",
	},
	{
		Author: "William ShakespeareAs You Like It",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "All the world's a stage, And all the men and women merely players; They have their exits and their entrances; And one man in his time plays many parts, His acts being seven ages.",
	},
	{
		Author: "Antoine de Saint-Exupéry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "Goodbye,\" said the fox. \"And now here is my secret, a very simple secret: It is only with the heart that one can see rightly; what is essential is invisible to the eye.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "This is your life and its ending one moment at a time.",
	},
	{
		Author: "Colleen HooverSlammed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1464032240p2/5430144.jpg",
		Content: "Don't take life too seriously. Punch it in the face when it needs a good hit. Laugh at it.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "How did you die?\"\"We er....drowned in a bathtub.\"\"All three of you?\"\"It was a big bathtub.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "If no one in the entire world cared about you, did you really exist at all?",
	},
	{
		Author: "Haruki MurakamiKafka on the Shore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Sometimes fate is like a small sandstorm that keeps changing directions. You change direction but the sandstorm chases you. You turn again, but the storm adjusts. Over and over you play this out, like some ominous dance with death just before dawn. Why? Because this storm isn't something that blew in from far away, something that has nothing to do with you. This storm is you. Something inside of you. So all you can do is give in to it, step right inside the storm, closing your eyes and plugging up your ears so the sand doesn't get in, and walk through it, step by step. There's no sun there, no moon, no direction, no sense of time. Just fine white sand swirling up into the sky like pulverized bones. That's the kind of sandstorm you need to imagine.An you really will have to make it through that violent, metaphysical, symbolic storm. No matter how metaphysical or symbolic it might be, make no mistake about it: it will cut through flesh like a thousand razor blades. People will bleed there, and you will bleed too. Hot, red blood. You'll catch that blood in your hands, your own blood and the blood of others.And once the storm is over you won't remember how you made it through, how you managed to survive. You won't even be sure, in fact, whether the storm is really over. But one thing is certain. When you come out of the storm you won't be the same person who walked in. That's what this storm's all about.",
	},
	{
		Author: "Margaret Fuller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244049471p2/113907.jpg",
		Content: "Today a reader, tomorrow a leader.",
	},
	{
		Author: "Hans Christian AndersenThe Little Mermaid",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625240p2/6378.jpg",
		Content: "But a mermaid has no tears, and therefore she suffers so much more.",
	},
	{
		Author: "Tina FeyBossypants",
		AuthorAvatar: "https://images.gr-assets.com/authors/1286658273p2/4385839.jpg",
		Content: "Do your thing and don't care if they like it.",
	},
	{
		Author: "Gabriel García MárquezOne Hundred Years of Solitude",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502310670p2/13450.jpg",
		Content: "It's enough for me to be sure that you and I exist at this moment.",
	},
	{
		Author: "Victor HugoLes Misérables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "Even the darkest night will end and the sun will rise.",
	},
	{
		Author: "Pablo NerudaSelected Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "Well, nowIf little by little you stop loving meI shall stop loving youLittle by littleIf suddenly you forget meDo not look for meFor I shall already have forgotten youIf you think it long and mad the wind of banners that passes through my lifeAnd you decide to leave me at the shore of the heart where I have rootsRememberThat on that day, at that hour, I shall lift my armsAnd my roots will set off to seek another land",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "There is nothing noble in being superior to your fellow man; true nobility is being superior to your former self.",
	},
	{
		Author: "Robert Bloch",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208225228p2/12540.jpg",
		Content: "Despite my ghoulish reputation, I really have the heart of a small boy. I keep it in a jar on my desk.",
	},
	{
		Author: "Jodi PicoultNineteen Minutes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "If you gave someone your heart and they died, did they take it with them? Did you spend the rest of forever with a hole inside you that couldn't be filled?",
	},
	{
		Author: "Leonardo da Vinci",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399080979p2/13560.jpg",
		Content: "Once you have tasted flight, you will forever walk the earth with your eyes turned skyward, for there you have been, and there you will always long to return.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "That was the thing. You never got used to it, the idea of someone being gone. Just when you think it's reconciled, accepted, someone points it out to you, and it just hits you all over again, that shocking.",
	},
	{
		Author: "Mary Wollstonecraft ShelleyFrankenstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1386351586p2/11139.jpg",
		Content: "Nothing is so painful to the human mind as a great and sudden change.",
	},
	{
		Author: "Victor HugoLes Misérables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "What Is Love? I have met in the streets a very poor young man who was in love. His hat was old, his coat worn, the water passed through his shoes and the stars through his soul",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Is 'fat' really the worst thing a human being can be? Is 'fat' worse than 'vindictive', 'jealous', 'shallow', 'vain', 'boring' or 'cruel'? Not to me.",
	},
	{
		Author: "Irving StoneClarence Darrow for the Defense",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208472912p2/12833.jpg",
		Content: "There are no faster or firmer friendships than those formed between people who love the same books.",
	},
	{
		Author: "Stephenie MeyerThe Host",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "It's not the face, but the expressions on it. It's not the voice, but what you say. It's not how you look in that body, but the thing you do with it. You are beautiful.",
	},
	{
		Author: "Henry David Thoreau",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "How vain it is to sit down to write when you have not stood up to live.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Oh well... I'd just been thinking, if you had died, you'd have been welcome to share my toilet.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "there are worse thingsthan being alonebut it often takesdecades to realize thisand most often when you doit's too lateand there's nothing worsethan too late",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "The one thing that doesn't abide by majority rule is a person's conscience.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "The thing about growing up with Fred and George,\" said Ginny thoughtfully, \"is that you sort of start thinking anything's possible if you've got enough nerve.",
	},
	{
		Author: "Augusten Burroughs",
		AuthorAvatar: "https://images.gr-assets.com/authors/1459810857p2/3058.jpg",
		Content: "I, myself, am made entirely of flaws, stitched together with good intentions.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Live to the point of tears.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "At some point, you just pull off the Band-Aid, and it hurts, but then it's over and you're relieved.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Don't stop there. I suppose there are also, what, vampires and werewolves and zombies?\"\"Of course there are. Although you mostly find zombies farther south, where the voudun priests are.\"\"What about mummies? Do they only hang around Egypt?\"\"Don't be ridiculous. No one believes in mummies.",
	},
	{
		Author: "Bob MarleyGuitar Chord Songbook - Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "If she's amazing, she won't be easy. If she's easy, she won't be amazing. If she's worth it, you wont give up. If you give up, you're not worthy. ... Truth is, everybody is going to hurt you; you just gotta find the ones worth suffering for.",
	},
	{
		Author: "August Wilson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321642333p2/13944.jpg",
		Content: "Confront the dark parts of yourself, and work to banish them with illumination and forgiveness. Your willingness to wrestle with your demons will cause your angels to sing.",
	},
	{
		Author: "David Foster WallaceInfinite Jest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466019433p2/4339.jpg",
		Content: "I do things like get in a taxi and say, \"The library, and step on it.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "If you're losing your soul and you know it, then you've still got a soul left to lose",
	},
	{
		Author: "Miguel de Cervantes SaavedraDon Quixote",
		AuthorAvatar: "https://images.gr-assets.com/authors/1276109217p2/4037220.jpg",
		Content: "Finally, from so little sleeping and so much reading, his brain dried up and he went completely out of his mind.",
	},
	{
		Author: "Jerry Seinfeld",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206678597p2/19838.jpg",
		Content: "If a book about failures doesn't sell, is it a success?",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Sometimes, you read a book and it fills you with this weird evangelical zeal, and you become convinced that the shattered world will never be put back together unless and until all living humans read the book. And then there are books like An Imperial Affliction, which you can't tell people about, books so special and rare and yours that advertising your affection feels like betrayal",
	},
	{
		Author: "Nelson Mandela",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308928296p2/367338.jpg",
		Content: "When a man is denied the right to live the life he believes in, he has no choice but to become an outlaw.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Yesterday I was clever, so I wanted to change the world. Today I am wise, so I am changing myself.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Yeah, Quirrell was a great teacher. There was just that minor drawback of him having Lord Voldemort sticking out of the back of his head!",
	},
	{
		Author: "Stephen KingRita Hayworth and Shawshank Redemption: A Story from Different Seasons",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Some birds are not meant to be caged, that's all. Their feathers are too bright, their songs too sweet and wild. So you let them go, or when you open the cage to feed them they somehow fly out past you. And the part of you that knows it was wrong to imprison them in the first place rejoices, but still, the place where you live is that much more drab and empty for their departure.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "To me, Fearless is not the absense of fear. It's not being completely unafraid. To me, Fearless is having fears. Fearless is having doubts. Lots of them. To me, Fearless is living in spite of those things that scare you to death.",
	},
	{
		Author: "Nora Ephron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1366180104p2/5691.jpg",
		Content: "Reading is escape, and the opposite of escape; it's a way to make contact with reality after a day of making things up, and it's a way of making contact with someone else's imagination after a day that's all too real.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "Say 'provoking' again. Your mouth looks provocative when you do.",
	},
	{
		Author: "Scott Adams",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200403987p2/5282.jpg",
		Content: "I love you like a fat kid loves cake!",
	},
	{
		Author: "Fyodor DostoyevskyThe Brothers Karamazov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "What is hell? I maintain that it is the suffering of being unable to love.",
	},
	{
		Author: "Anita Desai",
		AuthorAvatar: "https://images.gr-assets.com/authors/1296816073p2/8841.jpg",
		Content: "Wherever you go becomes a part of you somehow.",
	},
	{
		Author: "Richard DawkinsThe God Delusion",
		AuthorAvatar: "https://images.gr-assets.com/authors/1377030297p2/1194.jpg",
		Content: "We are all atheists about most of the gods that humanity has ever believed in. Some of us just go one god further.",
	},
	{
		Author: "Richard Feynman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1498916887p2/1429989.jpg",
		Content: "Nobody ever figures out what life is all about, and it doesn't matter. Explore the world. Nearly everything is really interesting if you go into it deeply enough.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "People are more difficult to work with than machines. And when you break a person, he can't be fixed.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Gravitation is not responsible for people falling in love.",
	},
	{
		Author: "Kathryn StockettThe Help",
		AuthorAvatar: "https://images.gr-assets.com/authors/1233458107p2/1943477.jpg",
		Content: "You is kind. You is smart. You is important.",
	},
	{
		Author: "Rob Reiner",
		AuthorAvatar: "",
		Content: "Everybody talks about wanting to change things and help and fix, but ultimately all you can do is fix yourself. And that's a lot. Because if you can fix yourself, it has a ripple effect.",
	},
	{
		Author: "Clare Boothe Luce",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206894543p2/332721.jpg",
		Content: "Money can't buy happiness, but it can make you awfully comfortable while you're being miserable.",
	},
	{
		Author: "Charles Dickens",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "Never close your lips to those whom you have already opened your heart.",
	},
	{
		Author: "Jonathan Safran FoerEverything Is Illuminated",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "It was not the feeling of completeness I so needed, but the feeling of not being empty.",
	},
	{
		Author: "Kurt Vonnegut",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "If I should ever die, God forbid, let this be my epitaph:THE ONLY PROOF HE NEEDEDFOR THE EXISTENCE OF GODWAS MUSIC",
	},
	{
		Author: "Nicholson BakerThe Anthologist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1247028956p2/15882.jpg",
		Content: "I woke up thinking a very pleasant thought. There is lots left in the world to read.",
	},
	{
		Author: "Tom Robbins",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430660478p2/197.jpg",
		Content: "We waste time looking for the perfect lover, instead of creating the perfect love.",
	},
	{
		Author: "George Bernard ShawBack to Methuselah",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "You see things; you say, 'Why?' But I dream things that never were; and I say 'Why not?",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "There are far, far better things ahead than any we leave behind.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Remember that wherever your heart is, there you will find your treasure.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Isabelle drifted over, Jace a pace behind her. She was wearing a long black dress with boots and an even longer cutaway coat of soft green velvet, the color of moss. \"I can't believe you did it!\" she exclaimed. \"How did you get Magnus to let Jace leave?\"\"Traded him for Alec,\" Clary said.Isabelle looked mildly alarmed. \"Not permanently?\"\"No,\" said Jace. \"Just for a few hours. Unless I don't come back,\" he added thoughtfully. \"In which case, maybe he does get to keep Alec. Think of it as a lease with an option to buy.\"Isabelle looked dubious. \"Mom and Dad won't be pleased if they find out.\"\"That you freed a possible criminal by trading away your brother to a warlock who looks like a gay Sonic the Hedgehog and dresses like the Child Catcher from Chitty Chitty Bang Bang?\" Simon inquired. \"No, probably not.",
	},
	{
		Author: "Joyce Carol OatesSolstice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454307466p2/3524.jpg",
		Content: "I never change, I simply become more myself.",
	},
	{
		Author: "Tahereh MafiShatter Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "All I ever wanted was to reach out and touch another human being not just with my hands but with my heart.",
	},
	{
		Author: "Jorge Luis Borges",
		AuthorAvatar: "https://images.gr-assets.com/authors/1496948506p2/500.jpg",
		Content: "I am not sure that I exist, actually. I am all the writers that I have read, all the people that I have met, all the women that I have loved; all the cities I have visited.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Education is an admirable thing, but it is well to remember from time to time that nothing that is worth knowing can be taught.",
	},
	{
		Author: "Federico García LorcaBlood Wedding and Yerma",
		AuthorAvatar: "https://images.gr-assets.com/authors/1287259251p2/44150.jpg",
		Content: "To burn with desire and keep quiet about it is the greatest punishment we can bring on ourselves.",
	},
	{
		Author: "John GreenAn Abundance of Katherines",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You don't remember what happened. What you remember becomes what happened.",
	},
	{
		Author: "Dave Eggers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1432065054p2/3371.jpg",
		Content: "Books have a unique way of stopping time in a particular moment and saying: Let’s not forget this.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "I will not let anyone walk through my mind with their dirty feet.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Stories may well be lies, but they are good lies that say true things, and which can sometimes pay the rent.",
	},
	{
		Author: "C.S. LewisThe Problem of Pain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "A man can no more diminish God's glory by refusing to worship Him than a lunatic can put out the sun by scribbling the word 'darkness' on the walls of his cell.",
	},
	{
		Author: "Dr. SeussOh, The Places You'll Go!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "So be sure when you step, Step with care and great tact. And remember that life's A Great Balancing Act. And will you succeed? Yes! You will, indeed! (98 and ¾ percent guaranteed) Kid, you'll move mountains.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "Always do sober what you said you'd do drunk. That will teach you to keep your mouth shut.",
	},
	{
		Author: "Voltaire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393683411p2/5754446.jpg",
		Content: "Judge a man by his questions rather than by his answers.",
	},
	{
		Author: "Arthur Conan Doyle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495008883p2/2448.jpg",
		Content: "It is a great thing to start life with a small number of really good books which are your very own.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "We have to allow ourselves to be loved by the people who really love us, the people who really matter. Too much of the time, we are blinded by our own pursuits of people to love us, people that don't even matter, while all that time we waste and the people who do love us have to stand on the sidewalk and watch us beg in the streets! It's time to put an end to this. It's time for us to let ourselves be loved.",
	},
	{
		Author: "Ernest HemingwayMen Without Women",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "The most painful thing is losing yourself in the process of loving someone too much, and forgetting that you are special too.",
	},
	{
		Author: "Emily Dickinson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536260p2/7440.jpg",
		Content: "Forever is composed of nows.",
	},
	{
		Author: "Italo CalvinoThe Uses of Literature",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501975461p2/155517.jpg",
		Content: "A classic is a book that has never finished saying what it has to say.",
	},
	{
		Author: "Stephenie MeyerEclipse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "Fall down again, Bella?'No, Emmett, I punched a werewolf in the face.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "There are no facts, only interpretations.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "When someone is crying, of course, the noble thing to do is to comfort them. But if someone is trying to hide their tears, it may also be noble to pretend you do not notice them.",
	},
	{
		Author: "Lauren OliverBefore I Fall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "Maybe you can afford to wait. Maybe for you there's a tomorrow. Maybe for you there's one thousand tomorrows, or three thousand, or ten, so much time you can bathe in it, roll around it, let it slide like coins through you fingers. So much time you can waste it.But for some of us there's only today. And the truth is, you never really know.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "It happens to everyone as they grow up. You find out who you are and what you want, and then you realize that people you've known forever don't see things the way you do. So you keep the wonderful memories, but find yourself moving on.",
	},
	{
		Author: "Kurt VonnegutA Man Without a Country",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "And on the subject of burning books: I want to congratulate librarians, not famous for their physical strength or their powerful political connections or their great wealth, who, all over this country, have staunchly resisted anti-democratic bullies who have tried to remove certain books from their shelves, and have refused to reveal to thought police the names of persons who have checked out those titles.So the America I loved still exists, if not in the White House or the Supreme Court or the Senate or the House of Representatives or the media. The America I love still exists at the front desks of our public libraries.",
	},
	{
		Author: "E.M. ForsterA Room with a View",
		AuthorAvatar: "https://images.gr-assets.com/authors/1402057803p2/86404.jpg",
		Content: "It isn't possible to love and part. You will wish that it was. You can transmute love, ignore it, muddle it, but you can never pull it out of you. I know by experience that the poets are right: love is eternal.",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "You know, when it works, love is pretty amazing. It's not overrated. There's a reason for all those songs.",
	},
	{
		Author: "Orson Scott CardEnder's Game",
		AuthorAvatar: "https://images.gr-assets.com/authors/1294099952p2/589.jpg",
		Content: "Perhaps it's impossible to wear an identity without becoming what you pretend to be.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "Fighting for peace is like screwing for virginity.",
	},
	{
		Author: "Edna St. Vincent Millay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183231710p2/33998.jpg",
		Content: "They say when you are missing someone that they are probably feeling the same, but I don't think it's possible for you to miss me as much as I'm missing you right now",
	},
	{
		Author: "Clive BarkerDays of Magic, Nights of War",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430330407p2/10366.jpg",
		Content: "Any fool can be happy. It takes a man with real heart to make beauty out of the stuff that makes us weep.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "How wrong is it for a woman to expect the man to build the world she wants, rather than to create it herself?",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "So it's you and a syringe against the Capitol? See, this is why no one lets you make the plans.",
	},
	{
		Author: "Logan Pearsall Smith",
		AuthorAvatar: "https://images.gr-assets.com/authors/1252194829p2/181547.jpg",
		Content: "People say that life is the thing, but I prefer reading.",
	},
	{
		Author: "Ian Fleming",
		AuthorAvatar: "https://images.gr-assets.com/authors/1364532740p2/2565.jpg",
		Content: "Never say 'no' to adventures. Always say 'yes,' otherwise you'll lead a very dull life.",
	},
	{
		Author: "Kurt VonnegutIf This Isn't Nice, What Is?: Advice for the Young",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "We have to continually be jumping off cliffs and developing our wings on the way down.",
	},
	{
		Author: "Aldous HuxleyMusic at Night and Other Essays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "After silence, that which comes nearest to expressing the inexpressible is music.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It's just that most really good-looking people are stupid, so I exceed expectations.''Right, it's primarily his hotness,' I said.'It can be sort of blinding,' he said.'It actually did blind our friend Isaac,' I said.'Terrible tragedy, that. But can I help my own deadly beauty?''You cannot.''It is my burden, this beautiful face.''Not to mention your body.''Seriously, don't even get me started on my hot bod. You don't want to see me naked, Dave. Seeing me naked actually took Hazel Grace's breath away,' he said, nodding toward the oxygen tank.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "The Road goes ever on and onDown from the door where it began.Now far ahead the Road has gone,And I must follow, if I can,Pursuing it with eager feet,Until it joins some larger wayWhere many paths and errands meet.And whither then? I cannot say",
	},
	{
		Author: "Ralph EllisonInvisible Man",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200336431p2/7508.jpg",
		Content: "What and how much had I lost by trying to do only what was expected of me instead of what I myself had wished to do?",
	},
	{
		Author: "Audrey NiffeneggerThe Time Traveler's Wife",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367342548p2/498072.jpg",
		Content: "Don't you think it's better to be extremely happy for a short while, even if you lose it, than to be just okay for your whole life?",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "Happiness is the consequence of personal effort. You fight for it, strive for it, insist upon it, and sometimes even travel around the world looking for it. You have to participate relentlessly in the manifestations of your own blessings. And once you have achieved a state of happiness, you must never become lax about maintaining it. You must make a mighty effort to keep swimming upward into that happiness forever, to stay afloat on top of it.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "What are men to rocks and mountains?",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "And when her lips met mine, I knew that I could live to be a hundred and visit every country in the world, but nothing would ever compare to that single moment when I first kissed the girl of my dreams and knew that my love would last forever.",
	},
	{
		Author: "أحمد خالد توفيق",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316682283p2/1479015.jpg",
		Content: "أحيانًا يساعدنا الآخرون بأن يكونوا فى حياتنا فحسب",
	},
	{
		Author: "T.H. WhiteThe Once and Future King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246071532p2/426944.jpg",
		Content: "The bravest people are the ones who don’t mind looking like cowards.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "The silence depressed me. It wasn't the silence of silence. It was my own silence.",
	},
	{
		Author: "Johnny Cash",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207242655p2/20062.jpg",
		Content: "All your life, you will be faced with a choice. You can choose love or hate…I choose love.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Dance, when you're broken open. Dance, if you've torn the bandage off. Dance in the middle of the fighting. Dance in your blood. Dance when you're perfectly free.",
	},
	{
		Author: "Gayle FormanIf I Stay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1472502468p2/295178.jpg",
		Content: "I realize now that dying is easy. Living is hard.",
	},
	{
		Author: "Barbara KingsolverAnimal Dreams",
		AuthorAvatar: "https://images.gr-assets.com/authors/1350499031p2/3541.jpg",
		Content: "The very least you can do in your life is figure out what you hope for. And the most you can do is live inside that hope. Not admire it from a distance but live right in it, under its roof.",
	},
	{
		Author: "Cornelia FunkeInkspell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437000100p2/15873.jpg",
		Content: "Isn't it odd how much fatter a book gets when you've read it several times?\" Mo had said...\"As if something were left between the pages every time you read it. Feelings, thoughts, sounds, smells...and then, when you look at the book again many years later, you find yourself there, too, a slightly younger self, slightly different, as if the book had preserved you like a pressed flower...both strange and familiar.",
	},
	{
		Author: "Lois LowryGathering Blue",
		AuthorAvatar: "https://images.gr-assets.com/authors/1348162077p2/2493.jpg",
		Content: "Take pride in your pain; you are stronger than those who have none",
	},
	{
		Author: "J.R. WardDark Lover",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437768069p2/20248.jpg",
		Content: "Welcome to the wonderful world of jealousy, he thought. For the price of admission, you get a splitting headache, a nearly irresistable urge to commit murder, and an inferiority complex. Yippee.",
	},
	{
		Author: "Christopher PaoliniEldest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1412963950p2/8349.jpg",
		Content: "Live in the present, remember the past, and fear not the future, for it doesn't exist and never shall. There is only now.",
	},
	{
		Author: "Philip K. DickI Hope I Shall Arrive Soon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1264613853p2/4764.jpg",
		Content: "Reality is that which, when you stop believing in it, doesn't go away.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Find what you love and let it kill you.",
	},
	{
		Author: "Patrick NessA Monster Calls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244216486p2/370361.jpg",
		Content: "There is not always a good guy. Nor is there always a bad one. Most people are somewhere in between.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "So you're always honest,\" I said.\"Aren't you?\"\"No,\" I told him. \"I'm not.\"\"Well, that's good to know, I guess.\"\"I'm not saying I'm a liar,\" I told him. He raised his eyebrows. \"That's not how I meant it, anyways.\"\"How'd you mean it, then?\"\"I just...I don't always say what I feel.\"\"Why not?\"\"Because the truth sometimes hurts,\" I said.\"Yeah,\" he said. \"So do lies, though.",
	},
	{
		Author: "Ivan Turgenev",
		AuthorAvatar: "https://images.gr-assets.com/authors/1239589274p2/410680.jpg",
		Content: "If we wait for the moment when everything, absolutely everything is ready, we shall never begin.",
	},
	{
		Author: "Bill  Nye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419933903p2/7373044.jpg",
		Content: "Everyone you will ever meet knows something you don't.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Things we lose have a way of coming back to us in the end, if not always in the way we expect.",
	},
	{
		Author: "Sarah DessenLock and Key",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "It's a lot easier to be lost than found. It's the reason we're always searching and rarely discovered--so many locks not enough keys.",
	},
	{
		Author: "T.S. EliotFour Quartets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1507887310p2/18540.jpg",
		Content: "For last year's words belong to last year's language And next year's words await another voice.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Love is not affectionate feeling, but a steady wish for the loved person's ultimate good as far as it can be obtained.",
	},
	{
		Author: "Benjamin Disraeli",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212684660p2/47030.jpg",
		Content: "There are three types of lies -- lies, damn lies, and statistics.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "Atticus told me to delete the adjectives and I'd have the facts.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You here to finish me off, Sweetheart?",
	},
	{
		Author: "Terry Pratchett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "In ancient times cats were worshipped as gods; they have not forgotten this.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "The man of knowledge must be able not only to love his enemies but also to hate his friends.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "The world breaks everyone, and afterward, many are strong at the broken places.",
	},
	{
		Author: "Stewart O'NanThe Odds: A Love Story",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206592605p2/18341.jpg",
		Content: "You couldn't relive your life, skipping the awful parts, without losing what made it worthwhile. You had to accept it as a whole--like the world, or the person you loved.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "And I like large parties. They’re so intimate. At small parties there isn’t any privacy.",
	},
	{
		Author: "Kurt VonnegutSlaughterhouse-Five",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "And so it goes...",
	},
	{
		Author: "Nathaniel HawthorneFanshawe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1291476587p2/7799.jpg",
		Content: "A single dream is more powerful than a thousand realities.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "When I was a boy of 14, my father was so ignorant I could hardly stand to have the old man around. But when I got to be 21, I was astonished at how much the old man had learned in seven years.",
	},
	{
		Author: "Chinua Achebe",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1436918703p2/29190.jpg",
		Content: "If you don't like someone's story, write your own.",
	},
	{
		Author: "Natalie BabbittTuck Everlasting",
		AuthorAvatar: "https://images.gr-assets.com/authors/1226361832p2/1954.jpg",
		Content: "Don't be afraid of death; be afraid of an unlived life. You don't have to live forever, you just have to live.",
	},
	{
		Author: "Khaled HosseiniThe Kite Runner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "And that's the thing about people who mean everything they say. They think everyone else does too.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "To define is to limit.",
	},
	{
		Author: "Henry Van DykeMusic and Other Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1317333338p2/7568.jpg",
		Content: "Time is Too Slow for those who Wait, Too Swift for those who Fear, Too Long for those who Grieve, Too Short for those who Rejoice; But for those who Love, Time is not.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Tell your heart that the fear of suffering is worse than the suffering itself. And that no heart has ever suffered when it goes in search of its dreams, because every second of the search is a second's encounter with God and with eternity.",
	},
	{
		Author: "Christopher Paolini",
		AuthorAvatar: "https://images.gr-assets.com/authors/1412963950p2/8349.jpg",
		Content: "Books should go where they will be most appreciated, and not sit unread, gathering dust on a forgotten shelf, don't you agree?",
	},
	{
		Author: "Oscar WildeLady Windermere's Fan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I can resist anything except temptation.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "The meek may inherit the earth, but at the moment it belongs to the conceited. Like me.",
	},
	{
		Author: "William ShakespeareThe Passionate Pilgrim",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Words are easy, like the wind; Faithful friends are hard to find.",
	},
	{
		Author: "Honoré de Balzac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206567834p2/228089.jpg",
		Content: "Solitude is fine but you need someone to tell that solitude is fine.",
	},
	{
		Author: "Colette",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206665516p2/51575.jpg",
		Content: "You will do foolish things, but do them with enthusiasm.",
	},
	{
		Author: "Maurice Sendak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201028880p2/4489.jpg",
		Content: "A book is really like a lover. It arranges itself in your life in a way that is beautiful.",
	},
	{
		Author: "George Orwell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "Every generation imagines itself to be more intelligent than the one that went before it, and wiser than the one that comes after it.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Grover was sniffing the wind, looking nervous. He fished out his acorns and threw them into the sand, then played his pipes. They rearranged themselves in a pattern that made no sense to me, but Grover looked concerned. \"That's us,\" he said. \"Those five nuts right there.\" \"Which one is me?\" I asked. \"The little deformed one,\" Zoe suggested. \"Oh, shut up.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "A cynic is a man who knows the price of everything, and the value of nothing.",
	},
	{
		Author: "Joss Whedon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1302721520p2/18015.jpg",
		Content: "Remember to always be yourself. Unless you suck.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It was books that made me feel that perhaps I was not completely alone. They could be honest with me, and I with them.",
	},
	{
		Author: "Norton JusterThe Phantom Tollbooth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201117378p2/214.jpg",
		Content: "So many things are possible just as long as you don't know they're impossible.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "Why are women... so much more interesting to men than men are to women?",
	},
	{
		Author: "J.R.R. TolkienThe Lord of the Rings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Three Rings for the Elven-kings under the sky, Seven for the Dwarf-lords in halls of stone,Nine for Mortal Men, doomed to die, One for the Dark Lord on his dark throneIn the Land of Mordor where the Shadows lie. One Ring to rule them all, One Ring to find them, One Ring to bring them all and in the darkness bind them.In the Land of Mordor where the Shadows lie.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Don't part with your illusions. When they are gone you may still exist, but you have ceased to live.",
	},
	{
		Author: "Carl SaganPale Blue Dot: A Vision of the Human Future in Space",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475953320p2/10538.jpg",
		Content: "Look again at that dot. That's here. That's home. That's us. On it everyone you love, everyone you know, everyone you ever heard of, every human being who ever was, lived out their lives. The aggregate of our joy and suffering, thousands of confident religions, ideologies, and economic doctrines, every hunter and forager, every hero and coward, every creator and destroyer of civilization, every king and peasant, every young couple in love, every mother and father, hopeful child, inventor and explorer, every teacher of morals, every corrupt politician, every \"superstar,\" every \"supreme leader,\" every saint and sinner in the history of our species lived there-on a mote of dust suspended in a sunbeam.The Earth is a very small stage in a vast cosmic arena. Think of the endless cruelties visited by the inhabitants of one corner of this pixel on the scarcely distinguishable inhabitants of some other corner, how frequent their misunderstandings, how eager they are to kill one another, how fervent their hatreds. Think of the rivers of blood spilled by all those generals and emperors so that, in glory and triumph, they could become the momentary masters of a fraction of a dot.Our posturings, our imagined self-importance, the delusion that we have some privileged position in the Universe, are challenged by this point of pale light. Our planet is a lonely speck in the great enveloping cosmic dark. In our obscurity, in all this vastness, there is no hint that help will come from elsewhere to save us from ourselves.The Earth is the only world known so far to harbor life. There is nowhere else, at least in the near future, to which our species could migrate. Visit, yes. Settle, not yet. Like it or not, for the moment the Earth is where we make our stand.It has been said that astronomy is a humbling and character-building experience. There is perhaps no better demonstration of the folly of human conceits than this distant image of our tiny world. To me, it underscores our responsibility to deal more kindly with one another, and to preserve and cherish the pale blue dot, the only home we've ever known.",
	},
	{
		Author: "Raymond ChandlerThe Long Goodbye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206535318p2/1377.jpg",
		Content: "To say goodbye is to die a little.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "If you are irritated by every rub, how will your mirror be polished?",
	},
	{
		Author: "Elizabeth Peters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232144920p2/16549.jpg",
		Content: "No woman really wants a man to carry her off; she only wants him to want to do it.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "It’s probably not just by chance that I’m alone. It would be very hard for a man to live with me, unless he’s terribly strong. And if he’s stronger than I, I’m the one who can’t live with him. … I’m neither smart nor stupid, but I don’t think I’m a run-of-the-mill person. I’ve been in business without being a businesswoman, I’ve loved without being a woman made only for love. The two men I’ve loved, I think, will remember me, on earth or in heaven, because men always remember a woman who caused them concern and uneasiness. I’ve done my best, in regard to people and to life, without precepts, but with a taste for justice.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Life, he realize, was much like a song. In the beginning there is mystery, in the end there is confirmation, but it's in the middle where all the emotion resides to make the whole thing worthwhile.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I don't want to be at the mercy of my emotions. I want to use them, to enjoy them, and to dominate them.",
	},
	{
		Author: "Joel OsteenYour Best Life Now: 7 Steps to Living at Your Full Potential",
		AuthorAvatar: "https://images.gr-assets.com/authors/1452791213p2/55044.jpg",
		Content: "You must make a decision that you are going to move on. It wont happen automatically. You will have to rise up and say, ‘I don’t care how hard this is, I don’t care how disappointed I am, I’m not going to let this get the best of me. I’m moving on with my life.",
	},
	{
		Author: "Pablo NerudaTwenty Love Poems and a Song of Despair",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "I wantTo do with you what spring does with the cherry trees.",
	},
	{
		Author: "Betty  SmithA Tree Grows in Brooklyn",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1470760765p2/29673.jpg",
		Content: "The world was hers for the reading.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I am selfish. I am brave.",
	},
	{
		Author: "Tennessee WilliamsConversations with Tennessee Williams",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206504901p2/7751.jpg",
		Content: "If I got rid of my demons, I’d lose my angels.",
	},
	{
		Author: "Margaret Thatcher",
		AuthorAvatar: "https://images.gr-assets.com/authors/1344397927p2/198468.jpg",
		Content: "In politics, If you want anything said, ask a man. If you want anything done, ask a woman.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "The first draft of anything is shit.",
	},
	{
		Author: "Christopher MorleyPipefuls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199111805p2/30802.jpg",
		Content: "There is no mistaking a real book when one meets it. It is like falling in love.",
	},
	{
		Author: "Virginia WoolfA Room of One's Own",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "One cannot think well, love well, sleep well, if one has not dined well.",
	},
	{
		Author: "Louis de BernièresCaptain Corelli's Mandolin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1438215719p2/2313.jpg",
		Content: "Love is a temporary madness, it erupts like volcanoes and then subsides. And when it subsides, you have to make a decision. You have to work out whether your roots have so entwined together that it is inconceivable that you should ever part. Because this is what love is. Love is not breathlessness, it is not excitement, it is not the promulgation of promises of eternal passion, it is not the desire to mate every second minute of the day, it is not lying awake at night imagining that he is kissing every cranny of your body. No, don't blush, I am telling you some truths. That is just being \"in love\", which any fool can do. Love itself is what is left over when being in love has burned away, and this is both an art and a fortunate accident.",
	},
	{
		Author: "J.R.R. TolkienThe Hobbit",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Do you wish me a good morning, or mean that it is a good morning whether I want it or not; or that you feel good this morning; or that it is a morning to be good on?",
	},
	{
		Author: "Marjorie Pay Hinckley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1250494999p2/226482.jpg",
		Content: "The only way to get through life is to laugh your way through it. You either have to laugh or cry. I prefer to laugh. Crying gives me a headache.",
	},
	{
		Author: "Oriah Mountain Dreamer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1375409812p2/101574.jpg",
		Content: "It doesn't interest me what you do for a living. I want to know what you ache for, and if you dare to dream of meeting your heart's longing.It doesn't interest me how old you are. I want to know if you will risk looking like a fool for love, for your dream, for the adventure of being alive. It doesn't interest me what planets are squaring your moon. I want to know if you have touched the center of your own sorrow, if you have been opened by life's betrayals or have become shriveled and closed from fear of further pain!I want to know if you can sit with pain, mine or your own, without moving to hide it or fade it, or fix it. I want to know if you can be with joy, mine or your own, if you can dance with wildness and let the ecstasy fill you to the tips of your fingers and toes without cautioning us to be careful, to be realistic, to remember the limitations of being human. It doesn't interest me if the story you are telling me is true. I want to know if you can disappoint another to be true to yourself; if you can bear the accusation of betrayal and not betray your own soul; if you can be faithlessand therefore trustworthy. I want to know if you can see beauty even when it's not pretty, every day,and if you can source your own life from its presence. I want to know if you can live with failure, yours and mine, and still stand on the edge of the lake and shout to the silver of the full moon, “Yes!",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "I cannot remember the books I've read any more than the meals I have eaten; even so, they have made me.",
	},
	{
		Author: "Tina FeyBossypants",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1337186108p2/414466.jpg",
		Content: "Some people say, “Never let them see you cry.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I felt my lungs inflate with the onrush of scenery—air, mountains, trees, people. I thought, \"This is what it is to be happy.",
	},
	{
		Author: "Vincent van Gogh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1329405187p2/34583.jpg",
		Content: "I dream my painting and I paint my dream.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Imagine smiling after a slap in the face. Then think of doing it twenty-four hours a day.",
	},
	{
		Author: "Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Mad Girl's Love SongI shut my eyes and all the world drops dead;I lift my lids and all is born again.(I think I made you up inside my head.)The stars go waltzing out in blue and red,And arbitrary blackness gallops in:I shut my eyes and all the world drops dead.I dreamed that you bewitched me into bedAnd sung me moon-struck, kissed me quite insane.(I think I made you up inside my head.)God topples from the sky, hell's fires fade:Exit seraphim and Satan's men:I shut my eyes and all the world drops dead.I fancied you'd return the way you said,But I grow old and I forget your name.(I think I made you up inside my head.)I should have loved a thunderbird instead;At least when spring comes they roar back again.I shut my eyes and all the world drops dead.(I think I made you up inside my head.)",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Love is an untamed force. When we try to control it, it destroys us. When we try to imprison it, it enslaves us. When we try to understand it, it leaves us feeling lost and confused.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "I wanted to tell the book thief many things, about beauty and brutality. But what could I tell her about those things that she didn't already know? I wanted to explain that I am constantly overestimating and underestimating the human race-that rarely do I ever simply estimate it. I wanted to ask her how the same thing could be so ugly and so glorious, and its words and stories so damning and brilliant.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Live in the sunshine, swim the sea, drink the wild air.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "How did it get so late so soon?",
	},
	{
		Author: "Hafez",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466549548p2/15291385.jpg",
		Content: "And still, after all this time,The sun never says to the earth,\"You owe Me.\"Look what happens withA love like that,It lights the Whole Sky.",
	},
	{
		Author: "Rodney Dangerfield",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208829061p2/78470.jpg",
		Content: "I came from a real tough neighborhood. Once a guy pulled a knife on me. I knew he wasn't a professional, the knife had butter on it.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "I knew when I met you an adventure was going to happen.",
	},
	{
		Author: "John Cheever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208899860p2/7464.jpg",
		Content: "I've been homesick for countries I've never been, and longed to be where I couldn't be.",
	},
	{
		Author: "Jeffrey EugenidesMiddlesex",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374107943p2/1467.jpg",
		Content: "Biology gives you a brain. Life turns it into a mind.",
	},
	{
		Author: "John SteinbeckOf Mice and Men",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182118389p2/585.jpg",
		Content: "Maybe ever’body in the whole damn world is scared of each other.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I am an excitable person who only understands life lyrically, musically, in whom feelings are much stronger as reason. I am so thirsty for the marvelous that only the marvelous has power over me. Anything I can not transform into something marvelous, I let go. Reality doesn't impress me. I only believe in intoxication, in ecstasy, and when ordinary life shackles me, I escape, one way or another. No more walls.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "I'm quite illiterate, but I read a lot. ",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "Someday, somewhere - anywhere, unfailingly, you'll find yourself, and that, and only that, can be the happiest or bitterest hour of your life.",
	},
	{
		Author: "Stephenie MeyerEclipse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "Look after my heart - I've left it with you.",
	},
	{
		Author: "Haruki Murakami",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "No matter how much suffering you went through, you never wanted to let go of those memories.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Every great love starts with a great story...",
	},
	{
		Author: "Frank McCourtAngela's Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263585664p2/3347.jpg",
		Content: "You might be poor, your shoes might be broken, but your mind is a palace.",
	},
	{
		Author: "William ShakespeareThe Merchant of Venice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "You speak an infinite deal of nothing.",
	},
	{
		Author: "Margaret Atwood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "I would like to be the air that inhabits you for a moment only. I would like to be that unnoticed and that necessary.",
	},
	{
		Author: "Chico Xavier",
		AuthorAvatar: "",
		Content: "‎Though nobody can go back and make a new beginning... Anyone can start over and make a new ending.",
	},
	{
		Author: "J.R.R. TolkienThe Two Towers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "War must be, while we defend our lives against a destroyer who would devour all; but I do not love the bright sword for its sharpness, nor the arrow for its swiftness, nor the warrior for his glory. I love only that which they defend.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "It's often just enough to be with someone. I don't need to touch them. Not even talk. A feeling passes between you both. You're not alone.",
	},
	{
		Author: "William Arthur Ward",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455493296p2/416931.jpg",
		Content: "The pessimist complains about the wind; the optimist expects it to change; the realist adjusts the sails.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "He does something to me, that boy. Every time. It’s his only detriment. He steps on my heart. He makes me cry.",
	},
	{
		Author: "Anne Frank",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "I can shake off everything as I write; my sorrows disappear, my courage is reborn.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "You were born with wings, why prefer to crawl through life?",
	},
	{
		Author: "Washington Irving",
		AuthorAvatar: "https://images.gr-assets.com/authors/1218187394p2/28525.jpg",
		Content: "There is a sacredness in tears....They are the messengers of overwhelming grief, of deep contrition and of unspeakable love.",
	},
	{
		Author: "W.P. Kinsella",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212093640p2/32549.jpg",
		Content: "Success is getting what you want, happiness is wanting what you get",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "What makes the desert beautiful,' said the little prince, 'is that somewhere it hides a well...",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Have enough courage to trust love one more time and always one more time.",
	},
	{
		Author: "Kristin CashoreGraceling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1273894652p2/1373880.jpg",
		Content: "I'm not going to wear a red dress,\" she said.\"It would look stunning, My Lady,\" she called.She spoke to the bubbles gathered on the surface of the water. \"If there's anyone I wish to stun at dinner, I'll hit him in the face.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "There is a stubbornness about me that never can bear to be frightened at the will of others. My courage always rises at every attempt to intimidate me.",
	},
	{
		Author: "Malcolm XBy Any Means Necessary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1185201379p2/17435.jpg",
		Content: "You're not to be so blind with patriotism that you can't face reality. Wrong is wrong, no matter who does it or says it.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Clary, Despite everything, I can't bear the thought of this ring being lost forever, any more then I can bear the thought of leaving you forever. And though I have no choice about the one, at least I can choose about the other. I'm leaving you our family ring because you have as much right to it as I do.I'm writing this watching the sun come up. You're asleep, dreams moving behind your restless eyelids. I wish I knew what you were thinking. I wish I could slip into your head and see the world the way you do. I wish I could see myself the way you do. But maybe I dont want to see that. Maybe it would make me feel even more than I already do that I'm perpetuating some kind of Great Lie on you, and I couldn't stand that. I belong to you. You could do anything you wanted with me and I would let you. You could ask anything of me and I'd break myself trying to make you happy. My heart tells me this is the best and greatest feeling I have ever had. But my mind knows the difference between wanting what you can't have and wanting what you shouldn't want. And I shouldn't want you.All night I've watched you sleeping, watched the moonlight come and go, casting its shadows across your face in black and white. I've never seen anything more beautiful. I think of the life we could have had if things were different, a life where this night is not a singular event, separate from everything else that's real, but every night. But things aren't different, and I can't look at you without feeling like I've tricked you into loving me.The truth no one is willing to say out loud is that no one has a shot against Valentine but me. I can get close to him like no one else can. I can pretend I want to join him and he'll believe me, up until that last moment where I end it all, one way or another. I have something of Sebastian's; I can track him to where my father's hiding, and that's what I'm going to do. So I lied to you last night. I said I just wanted one night with you. But I want every night with you. And that's why I have to slip out of your window now, like a coward. Because if I had to tell you this to your face, I couldn't make myself go. I don't blame you if you hate me, I wish you would. As long as I can still dream, I will dream of you. _Jace",
	},
	{
		Author: "Paul Simon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207839259p2/89923.jpg",
		Content: "I've got nothing to do today but smile.",
	},
	{
		Author: "Ralph EllisonInvisible Man",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200336431p2/7508.jpg",
		Content: "When I discover who I am, I’ll be free.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "The individual has always had to struggle to keep from being overwhelmed by the tribe. If you try it, you will be lonely often, and sometimes frightened. But no price is too high to pay for the privilege of owning yourself.",
	},
	{
		Author: "Ally CondieMatched",
		AuthorAvatar: "https://images.gr-assets.com/authors/1361564387p2/1304470.jpg",
		Content: "Growing apart doesn't change the fact that for a long time we grew side by side; our roots will always be tangled. I'm glad for that.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Understanding is the first step to acceptance, and only with acceptance can there be recovery.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "And so with the sunshine and the great bursts of leaves growing on the trees, just as things grow in fast movies, I had that familiar conviction that life was beginning over again with the summer.",
	},
	{
		Author: "Terry PratchettSmall Gods",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Time is a drug. Too much of it kills you.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Clothes make the man. Naked people have little or no influence on society.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "We should consider every day lost on which we have not danced at least once.",
	},
	{
		Author: "J.M. BarrieThe Admirable Crichton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "I'm not young enough to know everything.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Don't tell me,\" Jace said, \"Simon's turned himself into an ocelot and you want me to do something about it before Isabelle makes him into a stole. Well, you'll have have to wait till tomorrow. I'm out of commission.\" He pointed at himself - he was wearing blue pajamas with a hole in the sleeve. \"Look. Jammies.\"\"Jace,\" Clary said, \"this is important.\"\"Don't tell me,\" he said. \"You've got a drawing emergency. You need a nude model. Well, I'm not in the mood. You could always ask Hodge,\" he said as an afterthought. \"I hear he'll do anything for a -\"\"JACE!\" she interrupted him, her voice rising to a scream. \"JUST SHUT UP FOR A SECOND AND LISTEN, WILL YOU?",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "If you want to be a writer, you must do two things above all others: read a lot and write a lot.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "Have you ever noticed that anybody driving slower than you is an idiot, and anyone going faster than you is a maniac?",
	},
	{
		Author: "W. Somerset Maugham",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414096390p2/4176632.jpg",
		Content: "There are three rules for writing a novel. Unfortunately, no one knows what they are.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "You need to learn how to select your thoughts just the same way you select your clothes every day. This is a power you can cultivate. If you want to control things in your life so bad, work on the mind. That's the only thing you should be trying to control.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "And then he gives me a smile that just seems so genuinely sweet with just the right touch of shyness that unexpected warmth rushes through me.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Even death has a heart.",
	},
	{
		Author: "Veronica RothInsurgent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "We both have war inside us. Sometimes it keeps us alive. Sometimes it threatens to destroy us.",
	},
	{
		Author: "Stephen Hawking",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197404653p2/1401.jpg",
		Content: "We are just an advanced breed of monkeys on a minor planet of a very average star. But we can understand the Universe. That makes us something very special.",
	},
	{
		Author: "Lilith SaintcrowStrange Angels",
		AuthorAvatar: "https://images.gr-assets.com/authors/1358029579p2/131208.jpg",
		Content: "Better to be strong than pretty and useless.",
	},
	{
		Author: "Veronica RothInsurgent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Cruelty does not make a person dishonest, the same way bravery does not make a person kind.",
	},
	{
		Author: "Hermann HesseBäume. Betrachtungen und Gedichte",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499981916p2/1113469.jpg",
		Content: "For me, trees have always been the most penetrating preachers. I revere them when they live in tribes and families, in forests and groves. And even more I revere them when they stand alone. They are like lonely persons. Not like hermits who have stolen away out of some weakness, but like great, solitary men, like Beethoven and Nietzsche. In their highest boughs the world rustles, their roots rest in infinity; but they do not lose themselves there, they struggle with all the force of their lives for one thing only: to fulfil themselves according to their own laws, to build up their own form, to represent themselves. Nothing is holier, nothing is more exemplary than a beautiful, strong tree. When a tree is cut down and reveals its naked death-wound to the sun, one can read its whole history in the luminous, inscribed disk of its trunk: in the rings of its years, its scars, all the struggle, all the suffering, all the sickness, all the happiness and prosperity stand truly written, the narrow years and the luxurious years, the attacks withstood, the storms endured. And every young farmboy knows that the hardest and noblest wood has the narrowest rings, that high on the mountains and in continuing danger the most indestructible, the strongest, the ideal trees grow. Trees are sanctuaries. Whoever knows how to speak to them, whoever knows how to listen to them, can learn the truth. They do not preach learning and precepts, they preach, undeterred by particulars, the ancient law of life. A tree says: A kernel is hidden in me, a spark, a thought, I am life from eternal life. The attempt and the risk that the eternal mother took with me is unique, unique the form and veins of my skin, unique the smallest play of leaves in my branches and the smallest scar on my bark. I was made to form and reveal the eternal in my smallest special detail. A tree says: My strength is trust. I know nothing about my fathers, I know nothing about the thousand children that every year spring out of me. I live out the secret of my seed to the very end, and I care for nothing else. I trust that God is in me. I trust that my labor is holy. Out of this trust I live. When we are stricken and cannot bear our lives any longer, then a tree has something to say to us: Be still! Be still! Look at me! Life is not easy, life is not difficult. Those are childish thoughts. Let God speak within you, and your thoughts will grow silent. You are anxious because your path leads away from mother and home. But every step and every day lead you back again to the mother. Home is neither here nor there. Home is within you, or home is nowhere at all. A longing to wander tears my heart when I hear trees rustling in the wind at evening. If one listens to them silently for a long time, this longing reveals its kernel, its meaning. It is not so much a matter of escaping from one's suffering, though it may seem to be so. It is a longing for home, for a memory of the mother, for new metaphors for life. It leads home. Every path leads homeward, every step is birth, every step is death, every grave is mother. So the tree rustles in the evening, when we stand uneasy before our own childish thoughts: Trees have long thoughts, long-breathing and restful, just as they have longer lives than ours. They are wiser than we are, as long as we do not listen to them. But when we have learned how to listen to trees, then the brevity and the quickness and the childlike hastiness of our thoughts achieve an incomparable joy. Whoever has learned how to listen to trees no longer wants to be a tree. He wants to be nothing except what he is. That is home. That is happiness.",
	},
	{
		Author: "Julian BarnesFlaubert's Parrot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387175450p2/1462.jpg",
		Content: "Books say: She did this because. Life says: She did this. Books are where things are explained to you; life is where things aren't. I'm not surprised some people prefer books.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Why were you lurking under our window?\"\"Yes - yes, good point, Petunia! What were you doing under our windows, boy?\"\"Listening to the news,\" said Harry in a resigned voice.His aunt and uncle exchanged looks of outrage.\"Listening to the news! Again?\"\"Well, it changes every day, you see,\" said Harry.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Goblet of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It is my belief... that the truth is generally preferable to lies.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Dogs never bite me. Just humans.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Sometimes,' said Pooh, 'the smallest things take up the most room in your heart.",
	},
	{
		Author: "Milan Kundera",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465275207p2/6343.jpg",
		Content: "You can't measure the mutual affection of two human beings by the number of words they exchange.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "They’re not hideous,",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "I regret that it takes a life to learn how to live.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Our story has three parts: a beginning, a middle, and an end. And although this is the way all stories unfold, I still can't believe that ours didn't go on forever.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "All you have to do is write one true sentence. Write the truest sentence that you know.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "Life appears to me too short to be spent in nursing animosity or registering wrongs.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Forget safety.Live where you fear to live.Destroy your reputation.Be notorious.",
	},
	{
		Author: "Walter Cronkite",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206849391p2/30807.jpg",
		Content: "Whatever the cost of our libraries, the price is cheap compared to that of an ignorant nation.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "And I thought about how many people have loved those songs. And how many people got through a lot of bad times because of those songs. And how many people enjoyed good times with those songs. And how much those songs really mean. I think it would be great to have written one of those songs. I bet if I wrote one of them, I would be very proud. I hope the people who wrote those songs are happy. I hope they feel it's enough. I really do because they've made me happy. And I'm only one person.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "Every man I meet wants to protect me. I can't figure out what from.",
	},
	{
		Author: "Haruki MurakamiSputnik Sweetheart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "I dream. Sometimes I think that's the only right thing to do.",
	},
	{
		Author: "Kurt VonnegutA Man Without a Country",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "If you want to really hurt you parents, and you don't have the nerve to be gay, the least you can do is go into the arts. I'm not kidding. The arts are not a way to make a living. They are a very human way of making life more bearable. Practicing an art, no matter how well or badly, is a way to make your soul grow, for heaven's sake. Sing in the shower. Dance to the radio. Tell stories. Write a poem to a friend, even a lousy poem. Do it as well as you possible can. You will get an enormous reward. You will have created something.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "I care for myself. The more solitary, the more friendless, the more unsustained I am, the more I will respect myself.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It was, he thought, the difference between being dragged into the arena to face a battle to the death and walking into the arena with your head held high. Some people, perhaps, would say that there was little to choose between the two ways, but Dumbledore knew - and so do I, thought Harry, with a rush of fierce pride, and so did my parents - that there was all the difference in the world.",
	},
	{
		Author: "Mark TwainA Connecticut Yankee in King Arthur's Court",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "You can't depend on your eyes when your imagination is out of focus.",
	},
	{
		Author: "Ray Bradbury",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1444700042p2/1004.jpg",
		Content: "We are cups, constantly and quietly being filled. The trick is, knowing how to tip ourselves over and let the beautiful stuff out.",
	},
	{
		Author: "Emily Dickinson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536260p2/7440.jpg",
		Content: "That it will never come again is what makes life so sweet.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "How do you spell 'love'?\" - Piglet\"You don't spell it...you feel it.\" - Pooh",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Much of my life had been devoted to trying not to cry in front of people who loved me, so I knew what Augustus was doing. You clench your teeth. You look up. You tell yourself that if they see you cry, it will hurt them, and you will be nothing but a Sadness in their lives, and you must not become a mere sadness, so you will not cry, and you say all of this to yourself while looking up at the ceiling, and then you swallow even though your throat does not want to close and you look at the person who loves you and smile.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Life shrinks or expands in proportion to one's courage.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I realize only one person will be damaged beyond repair if Peeta dies. Me.",
	},
	{
		Author: "Pablo NerudaTwenty Love Poems and a Song of Despair",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "Tonight I can write the saddest linesI loved her, and sometimes she loved me too.",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "If you want to tell people the truth, make them laugh, otherwise they'll kill you.",
	},
	{
		Author: "Voltaire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393683411p2/5754446.jpg",
		Content: "‎Life is a shipwreck, but we must not forget to sing in the lifeboats.",
	},
	{
		Author: "Jamie McGuireBeautiful Disaster",
		AuthorAvatar: "https://images.gr-assets.com/authors/1479315727p2/4464118.jpg",
		Content: "I knew the second I met youthat there was something about you I needed. Turns out it wasn’t something about you at all. It was just you.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "What you do speaks so loudly that I cannot hear what you say.",
	},
	{
		Author: "Isaac Asimov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341965730p2/16667.jpg",
		Content: "In life, unlike chess, the game continues after checkmate.",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "Trees are poems the earth writes upon the sky, We fell them down and turn them into paper,That we may record our emptiness.",
	},
	{
		Author: "RumiThe Essential Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Don't be satisfied with stories, how things have gone with others. Unfold your own myth.",
	},
	{
		Author: "Corrie ten Boom",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217173231p2/102203.jpg",
		Content: "Never be afraid to trust an unknown future to a known God.",
	},
	{
		Author: "Stephenie MeyerNew Moon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "Before you, Bella, my life was like a moonless night. Very dark, but there were stars, points of light and reason. ...And then you shot across my sky like a meteor. Suddenly everything was on fire; there was brilliancy, there was beauty. When you were gone, when the meteor had fallen over the horizon, everything went black. Nothing had changed, but my eyes were blinded by the light. I couldn’t see the stars anymore. And there was no more reason, for anything.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "It's possible to go on, no matter how impossible it seems, and that in time, the grief . . . lessens. It may not go away completely, but after a while it's not so overwhelming.",
	},
	{
		Author: "John Locke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477998201p2/51746.jpg",
		Content: "Reading furnishes the mind only with materials of knowledge; it is thinking that makes what we read ours.",
	},
	{
		Author: "John A. Shedd",
		AuthorAvatar: "",
		Content: "A ship is safe in harbor, but that's not what ships are for.",
	},
	{
		Author: "Chuck PalahniukChoke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "The unreal is more powerful than the real. Because nothing is as perfect as you can imagine it. Because its only intangible ideas, concepts, beliefs, fantasies that last. Stone crumbles. Wood rots. People, well, they die. But things as fragile as a thought, a dream, a legend, they can go on and on. If you can change the way people think. The way they see themselves. The way they see the world. You can change the way people live their lives. That's the only lasting thing you can create.",
	},
	{
		Author: "Rick RiordanThe Sea of Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "The real story of the Fleece: there were these two children of Zeus, Cadmus and Europa, okay? They were about to get offered up as human sacrifices, when they prayed to Zeus to save them. So Zeus sent this magical flying ram with golden wool, which picked them up in Greece and carried them all the way to Colchis in Asia Minor. Well, actually it carried Cadmus. Europa fell off and died along the way, but that's not important.\"\"It was probably important to her.",
	},
	{
		Author: "Henry Miller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1511445828p2/147.jpg",
		Content: "The one thing we can never get enough of is love. And the one thing we never give enough of is love.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Of all the trees we could've hit, we had to get one that hits back.",
	},
	{
		Author: "Jorge Luis Borges",
		AuthorAvatar: "https://images.gr-assets.com/authors/1496948506p2/500.jpg",
		Content: "I cannot sleep unless I am surrounded by books.",
	},
	{
		Author: "W.G. SebaldVertigo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465928875p2/6580622.jpg",
		Content: "It is thanks to my evening reading alone that I am still more or less sane.",
	},
	{
		Author: "Bill Watterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "Sometimes when I'm talking, my words can't keep up with my thoughts. I wonder why we think faster than we speak. Probably so we can think twice.",
	},
	{
		Author: "Sara Henderson",
		AuthorAvatar: "",
		Content: "Don't wait for a light to appear at the end of the tunnel, stride down there and light the bloody thing yourself.",
	},
	{
		Author: "Robert Jordan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1175475715p2/6252.jpg",
		Content: "Any fool knows men and women think differently at times, but the biggest difference is this. Men forget, but never forgive; women forgive, but never forget.",
	},
	{
		Author: "Lemony Snicket",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "All the secrets of the world are contained in books. Read at your own risk.",
	},
	{
		Author: "Alice Walker",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1423959229p2/15083.jpg",
		Content: "The most common way people give up their power is by thinking they don't have any.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "It is not that I'm so smart. But I stay with the questions much longer.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "When you do things from your soul, you feel a river moving in you, a joy.",
	},
	{
		Author: "Neil GaimanThe Ocean at the End of the Lane",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I lived in books more than I lived anywhere else.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "love the life you live.live the life you love.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "He was my mum and dad's best friend. He's a convicted murderer, but he's broken out of wizard prison and he's on the run. He likes to keep in touch with me, though...keep up with my news...check if I'm happy...",
	},
	{
		Author: "Saul Bellow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470607356p2/4391.jpg",
		Content: "People can lose their lives in libraries. They ought to be warned.",
	},
	{
		Author: "Shel Silverstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "If you are a dreamer come inIf you are a dreamer a wisher a liarA hoper a pray-er a magic-bean-buyerIf youre a pretender com sit by my fireFor we have some flax golden tales to spinCome in! Come in!",
	},
	{
		Author: "David Byrne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1492631911p2/27078.jpg",
		Content: "Sometimes it's a form of love just to talk to somebody that you have nothing in common with and still be fascinated by their presence.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "What are all these?\" Clary asked.\"Vials of holy water, blessed knives, steel and silver blades,\" Jace said, piling the weapons on the floor beside him, \"electrum wire - not much use at the moment but it's always good to have spares - silver bullets, charms of protetion, crucifixes, stars of David-\"\"Jesus,\" said Clary\"I doubt he'd fit.\"\"Jace.\" Clary was appalled.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "One can never have enough socks,\" said Dumbledore. \"Another Christmas has come and gone and I didn't get a single pair. People will insist on giving me books.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Don't feel sorry for yourself. Only assholes do that.",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "I crave your mouth, your voice, your hair.Silent and starving, I prowl through the streets. Bread does not nourish me, dawn disrupts me, all day I hunt for the liquid measure of your steps. I hunger for your sleek laugh, your hands the color of a savage harvest, hunger for the pale stones of your fingernails, I want to eat your skin like a whole almond. I want to eat the sunbeam flaring in your lovely body, the sovereign nose of your arrogant face, I want to eat the fleeting shade of your lashes, and I pace around hungry, sniffing the twilight, hunting for you, for your hot heart, Like a puma in the barrens of Quitratue.",
	},
	{
		Author: "Mary Oliver",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429857566p2/23988.jpg",
		Content: "You do not have to be good.You do not have to walk on your kneesfor a hundred miles through the desert, repenting.You only have to let the soft animal of your body love what it loves.Tell me about despair, yours, and I will tell you mine.Meanwhile the world goes on.Meanwhile the sun and the clear pebbles of the rainare moving across the landscapes,over the prairies and the deep trees,the mountains and the rivers.Meanwhile the wild geese, high in the clean blue air,are heading home again.Whoever you are, no matter how lonely,the world offers itself to your imagination,calls to you like the wild geese, harsh and exciting –over and over announcing your placein the family of things.",
	},
	{
		Author: "Ayn RandAtlas Shrugged",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "If you don't know, the thing to do is not to get scared, but to learn.",
	},
	{
		Author: "Chuck PalahniukSurvivor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "You realize that our mistrust of the future makes it hard to give up the past.",
	},
	{
		Author: "Émile Zola",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189528657p2/4750.jpg",
		Content: "If you ask me what I came to do in this world, I, an artist, will answer you: I am here to live out loud.",
	},
	{
		Author: "Will Rogers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1250044478p2/132444.jpg",
		Content: "Even if you are on the right track, you’ll get run over if you just sit there.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "You possess other people's...bodies.\"He accepted that statement with a nod.\"Do you want to possess my body?\"\"I want to do a lot of things to your body, but that's not one of them.",
	},
	{
		Author: "Albert CamusThe Fall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "You know what charm is: a way of getting the answer yes without having asked any clear question.",
	},
	{
		Author: "Jon KrakauerInto the Wild",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430856379p2/1235.jpg",
		Content: "Happiness [is] only real when shared",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "As if you were on fire from within.The moon lives in the lining of your skin.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "A book must be the axe for the frozen sea within us.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "From the moment I picked up your book until I put it down, I was convulsed with laughter. Some day I intend reading it.",
	},
	{
		Author: "Bram StokerDracula",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202438456p2/6988.jpg",
		Content: "There are darknesses in life and there are lights, and you are one of the lights, the light of all lights.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Education is the ability to listen to almost anything without losing your temper or your self-confidence.",
	},
	{
		Author: "Gilda Radner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212084245p2/145047.jpg",
		Content: "I wanted a perfect ending. Now I've learned, the hard way, that some poems don't rhyme, and some stories don't have a clear beginning, middle, and end.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "Many a book is like a key to unknown chambers within the castle of one’s own self.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Did you know that for pretty much the entire history of the human species, the average life span was less than thirty years? You could count on ten years or so of real adulthood, right? There was no planning for retirement, There was no planning for a career. There was no planning. No time for plannning. No time for a future. But then the life spans started getting longer, and people started having more and more future. And now life has become the future. Every moment of your life is lived for the future--you go to high school so you can go to college so you can get a good job so you can get a nice house so you can afford to send your kids to college so they can get a good job so they can get a nice house so they can afford to send their kids to college.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The measure of intelligence is the ability to change.",
	},
	{
		Author: "Diana GabaldonDragonfly in Amber",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213918339p2/3617.jpg",
		Content: "I stood still, vision blurring, and in that moment, I heard my heart break. It was a small, clean sound, like the snapping of a flower's stem.",
	},
	{
		Author: "Edward Albee",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206561590p2/9322.jpg",
		Content: "You're alive only once, as far as we know, and what could be worse than getting to the end of your life and realizing you hadn't lived it?",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Yes, frosting. The final defense of the dying.",
	},
	{
		Author: "Neil GaimanFragile Things: Short Fictions and Wonders",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "There are so many fragile things, after all. People break so easily, and so do dreams and hearts.",
	},
	{
		Author: "William WordsworthLyrical Ballads",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288772244p2/64845.jpg",
		Content: "The best portion of a good man's life: his little, nameless unremembered acts of kindness and love.",
	},
	{
		Author: "L. Frank Baum",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383720421p2/3242.jpg",
		Content: "Never give up... No one knows what's going to happen next.",
	},
	{
		Author: "James Baldwin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343346341p2/10427.jpg",
		Content: "You think your pain and your heartbreak are unprecedented in the history of the world, but then you read.",
	},
	{
		Author: "Mary Wollstonecraft ShelleyFrankenstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1386351586p2/11139.jpg",
		Content: "Beware; for I am fearless, and therefore powerful.",
	},
	{
		Author: "Cormac McCarthyNo Country for Old Men",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414695980p2/4178.jpg",
		Content: "You never know what worse luck your bad luck has saved you from.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "If I'd observed all the rules I'd never have got anywhere.",
	},
	{
		Author: "Deborah ReberChicken Soup for the Teenage Soul",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519577913p2/71869.jpg",
		Content: "Letting go doesn't mean that you don't care about someone anymore. It's just realizing that the only person you really have control over is yourself.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You gave me a forever within the numbered days, and I'm grateful.",
	},
	{
		Author: "Victor HugoHugo's Works: William Shakespeare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "Music expresses that which cannot be said and on which it is impossible to be silent.",
	},
	{
		Author: "Arthur C. Clarke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357191481p2/7779.jpg",
		Content: "I'm sure the universe is full of intelligent life. It's just been too intelligent to come here.",
	},
	{
		Author: "Tamora PierceMelting Stones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209044273p2/8596.jpg",
		Content: "Every now and then I like to do as I'm told, just to confuse people.",
	},
	{
		Author: "Douglas AdamsThe Salmon of Doubt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "A learning experience is one of those things that says, 'You know that thing you just did? Don't do that.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "He who has a why to live for can bear almost any how.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "When it is dark enough, you can see the stars.",
	},
	{
		Author: "Ernest HemingwayA Farewell to Arms",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "Maybe...you'll fall in love with me all over again.\"\"Hell,\" I said, \"I love you enough now. What do you want to do? Ruin me?\"\"Yes. I want to ruin you.\"\"Good,\" I said. \"That's what I want too.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "I am a cage, in search of a bird.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "We can easily forgive a child who is afraid of the dark; the real tragedy of life is when men are afraid of the light.",
	},
	{
		Author: "John GuareLandscape of the Body",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321641691p2/13977.jpg",
		Content: "It's amazing how a little tomorrow can make up for a whole lot of yesterday.",
	},
	{
		Author: "Noël CowardBlithe Spirit",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211318407p2/120035.jpg",
		Content: "It's discouraging to think how many people are shocked by honesty and how few by deceit.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "That's why they call it the American Dream, because you have to be asleep to believe it.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "He smiled understandingly-much more than understandingly. It was one of those rare smiles with a quality of eternal reassurance in it, that you may come across four or five times in life. It faced--or seemed to face--the whole eternal world for an instant, and then concentrated on you with an irresistible prejudice in your favor. It understood you just as far as you wanted to be understood, believed in you as you would like to believe in yourself, and assured you that it had precisely the impression of you that, at your best, you hoped to convey.",
	},
	{
		Author: "Johann Wolfgang von GoetheWilhelm Meister's Apprenticeship",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454800433p2/285217.jpg",
		Content: "One ought, every day at least, to hear a little song, read a good poem, see a fine picture, and, if it were possible, to speak a few reasonable words.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "You think my first instinct is to protect you. Because you're small, or a girl, or a Stiff. But you're wrong.\"He leans his face close to mine and wraps his fingers around my chin. His hand smells like metal. When was the last time he held a gun, or a knife? My skin tingles at the point of contact, like he's transmitting electricity through his skin. \"My first instinct is to push you until you break, just to see how hard I have to press.\" he says, his fingers squeezing at the word break. My body tenses at the edge in his voice, so I am coiled as tight as a spring, and I forget to breathe.His dark eyes lifting to mine, he adds, \"But I resist it.\" \"Why...\" I swallow hard. \"Why is that your first instinct?\"\"Fear doesn't shut you down; it wakes you up. I've seen it. It's fascinating.\" He releases me but doesn't pull away, his hand grazing my jaw, my neck. \"Sometimes I just want to see it again. Want to see you awake.",
	},
	{
		Author: "Alan BennettThe History Boys",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234747177p2/11781.jpg",
		Content: "The best moments in reading are when you come across something – a thought, a feeling, a way of looking at things – which you had thought special and particular to you. Now here it is, set down by someone else, a person you have never met, someone even who is long dead. And it is as if a hand has come out and taken yours.",
	},
	{
		Author: "Vincent van Gogh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1329405187p2/34583.jpg",
		Content: "Be clearly aware of the stars and infinity on high. Then life seems almost enchanted after all.",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "So when the moon's only partly full, you only feel a little wolfy?\" \"You could say that.\"\"Well, you can go ahead and hang your head out the car window if you feel like it.\"\"I'm a werewolf, not a golden retriever.",
	},
	{
		Author: "Jay McInerneyThe Last of the Savages",
		AuthorAvatar: "https://images.gr-assets.com/authors/1178138066p2/14079.jpg",
		Content: "The capacity for friendship is God's way of apologizing for our families.",
	},
	{
		Author: "Rick Riordan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "I nodded, looking at Rachel with respect. \"You hit the Lord of the Titans in the eye with a blue plastic hairbrush.",
	},
	{
		Author: "Mo WillemsGoldilocks and the Three Dinosaurs",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199684071p2/33274.jpg",
		Content: "If you ever find yourself in the wrong story, leave.",
	},
	{
		Author: "Benjamin FranklinPoor Richard's Almanack",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215314094p2/289513.jpg",
		Content: "Three may keep a secret, if two of them are dead.",
	},
	{
		Author: "Jay AsherThirteen Reasons Why",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243931536p2/569269.jpg",
		Content: "No one knows for certain how much impact they have on the lives of other people. Oftentimes, we have no clue. Yet we push it just the same.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "If you pick up a starving dog and make him prosperous he will not bite you. This is the principal difference between a dog and man.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Today was good. Today was fun. Tomorrow is another one.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "A DEFINITION NOT FOUNDIN THE DICTIONARYNot leaving: an act of trust and love,often deciphered by children",
	},
	{
		Author: "Annie DillardThe Living",
		AuthorAvatar: "https://images.gr-assets.com/authors/1240154705p2/5209.jpg",
		Content: "She read books as one would breathe air, to fill up and live.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Peeta, you said at the interview you’d had a crush on me forever. When did forever start?Oh, let’s see. I guess the first day of school. We were five. You had on a red plaid dress and your hair...it was in two braids instead of one. My father pointed you out when we were waiting to line up.\"Your father? Why?\"He said, ‘See that little girl? I wanted to marry her mother, but she ran off with a coal miner.'\"What? You’re making that up!\"No, true story. And I said, 'A coal miner? Why did she want a coal miner if she could’ve had you?' And he said, 'Because when he sings...even the birds stop to listen.",
	},
	{
		Author: "Madonna",
		AuthorAvatar: "https://images.gr-assets.com/authors/1423859021p2/104438.jpg",
		Content: "I'm tough, I'm ambitious, and I know exactly what I want. If that makes me a bitch, okay.",
	},
	{
		Author: "Golda Meir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1267489010p2/223411.jpg",
		Content: "Don't be so humble - you are not that great.",
	},
	{
		Author: "Markus ZusakI Am the Messenger",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Maybe everyone can live beyond what they're capable of.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "Don't Gain The World \u0026 Lose Your Soul, Wisdom Is Better Than Silver Or Gold.",
	},
	{
		Author: "Thomas MannEssays of Three Decades",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430109860p2/19405.jpg",
		Content: "A writer is someone for whom writing is more difficult than it is for other people.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "If things start happening, don't worry, don't stew, just go right along and you'll start happening too.",
	},
	{
		Author: "Niels Bohr",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206830940p2/821936.jpg",
		Content: "An expert is a person who has made all the mistakes that can be made in a very narrow field.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I didn’t need you, you idiot. I picked you. And then you picked me back.",
	},
	{
		Author: "Marcus AureliusMeditations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430676293p2/17212.jpg",
		Content: "You have power over your mind - not outside events. Realize this, and you will find strength.",
	},
	{
		Author: "Dan BrownThe Lost Symbol",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399396714p2/630.jpg",
		Content: "Great minds are always feared by lesser minds.",
	},
	{
		Author: "C.S. LewisThe Silver Chair",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Crying is all right in its way while it lasts. But you have to stop sooner or later, and then you still have to decide what to do.",
	},
	{
		Author: "Thomas MertonNo Man Is an Island",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201097624p2/1711.jpg",
		Content: "The beginning of love is the will to let those we love be perfectly themselves, the resolution not to twist them to fit our own image. If in loving them we do not love what they are, but only their potential likeness to ourselves, then we do not love them: we only love the reflection of ourselves we find in them",
	},
	{
		Author: "Karen Marie MoningBloodfever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1497612477p2/48206.jpg",
		Content: "One day you will kiss a man you can't breathe without, and find that breath is of little consequence.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "I think there's just one kind of folks. Folks.",
	},
	{
		Author: "Søren Kierkegaard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202410387p2/6172.jpg",
		Content: "The function of prayer is not to influence God, but rather to change the nature of the one who prays.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "I knew, in the silence that followed, that anything could happen here. It might be too late: again, I might have missed my chance. But I would at least know I tried, that I took my heart and extended my hand, whatever the outcome.\"Okay,\" he said. He took a breath. \"What would you do, if you could do anything?\"I took a step toward him, closing the space between us. \"This,\" I said. And then I kissed him.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "If you do not tell the truth about yourself you cannot tell it about other people.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "The moment you doubt whether you can fly, you cease for ever to be able to do it.",
	},
	{
		Author: "Lisa Schroeder",
		AuthorAvatar: "https://images.gr-assets.com/authors/1378418358p2/588558.jpg",
		Content: "Was it hard?\" I ask.Letting go?\"Not as hard as holding on to something that wasn't real.",
	},
	{
		Author: "Julian F. Fleron",
		AuthorAvatar: "",
		Content: "The creative adult is the child who has survived.",
	},
	{
		Author: "Johann Wolfgang von Goethe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454800433p2/285217.jpg",
		Content: "If you treat an individual as he is, he will remain how he is. But if you treat him as if he were what he ought to be and could be, he will become what he ought to be and could be.",
	},
	{
		Author: "Thomas Jefferson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400443180p2/1673.jpg",
		Content: "Do you want to know who you are? Don't ask. Act! Action will delineate and define you.",
	},
	{
		Author: "Marcus AureliusMeditations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430676293p2/17212.jpg",
		Content: "Dwell on the beauty of life. Watch the stars, and see yourself running with them.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Don't feel bad, I'm usually about to die.",
	},
	{
		Author: "Anne FrankThe Diary of a Young Girl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "I've found that there is always some beauty left -- in nature, sunshine, freedom, in yourself; these can all help you.",
	},
	{
		Author: "Truman Capote",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419249359p2/431149.jpg",
		Content: "Failure is the condiment that gives success its flavor.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Your friend's poetry is terrible,\" he said.Clary blinked, caught momentarily off guard. \"What?\"\"I said his poetry was terrible. It sounds like he ate a dictionary and started vomiting up words at random.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "What if it lines up like it did in the Trojan War ... Athena versus Poseidon?\"\"I don't know. But I just know that I'll be fighting next to you.\"\"Why?\"\"Because you're my friend, Seaweed Brain. Any more stupid questions?",
	},
	{
		Author: "Joshua SlocumSailing Alone around the World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1253681607p2/120675.jpg",
		Content: "I had already found that it was not good to be alone, and so made companionship with what there was around me, sometimes with the universe and sometimes with my own insignificant self; but my books were always my friends, let fail all else.",
	},
	{
		Author: "Neil GaimanStardust",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "She says nothing at all, but simply stares upward into the dark sky and watches, with sad eyes, the slow dance of the infinite stars.",
	},
	{
		Author: "Douglas AdamsDirk Gently's Holistic Detective Agency",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Let's think the unthinkable, let's do the undoable. Let us prepare to grapple with the ineffable itself, and see if we may not eff it after all.",
	},
	{
		Author: "William ShakespeareJulius Caesar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Cowards die many times before their deaths; The valiant never taste of death but once. Of all the wonders that I yet have heard, It seems to me most strange that men should fear; Seeing that death, a necessary end, Will come when it will come.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Nowadays people know the price of everything and the value of nothing.",
	},
	{
		Author: "Patricia Highsmith",
		AuthorAvatar: "https://images.gr-assets.com/authors/1418715271p2/7622.jpg",
		Content: "My imagination functions much better when I don't have to speak to people.",
	},
	{
		Author: "Jane Goodall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282766982p2/18163.jpg",
		Content: "What you do makes a difference, and you have to decide what kind of difference you want to make.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "You will find that it is necessary to let things go; simply for the reason that they are heavy. So let them go, let go of them. I tie no weights to my ankles.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I wasn't actually in love, but I felt a sort of tender curiosity.",
	},
	{
		Author: "أحمد خالد توفيق",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316682283p2/1479015.jpg",
		Content: "أتمنى أن أبكي و أرتجف , التصق بواحد من الكبار , لكن الحقيقة القاسية هي أنك الكبار! .. أنت من يجب أن يمنح القوة و الأمن للآخرين!",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Winter is coming.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I'll just have them change the entry in the demonology textbook from 'almost extinct' to 'not extinct enough for Alec. He prefers his monsters really, really extinct.' Will that make you happy?",
	},
	{
		Author: "Jane AustenPersuasion",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I can listen no longer in silence. I must speak to you by such means as are within my reach. You pierce my soul. I am half agony, half hope. Tell me not that I am too late, that such precious feelings are gone for ever. I offer myself to you again with a heart even more your own than when you almost broke it, eight years and a half ago. Dare not say that man forgets sooner than woman, that his love has an earlier death. I have loved none but you. Unjust I may have been, weak and resentful I have been, but never inconstant. You alone have brought me to Bath. For you alone, I think and plan. Have you not seen this? Can you fail to have understood my wishes? I had not waited even these ten days, could I have read your feelings, as I think you must have penetrated mine. I can hardly write. I am every instant hearing something which overpowers me. You sink your voice, but I can distinguish the tones of that voice when they would be lost on others. Too good, too excellent creature! You do us justice, indeed. You do believe that there is true attachment and constancy among men. Believe it to be most fervent, most undeviating, in F. W.I must go, uncertain of my fate; but I shall return hither, or follow your party, as soon as possible. A word, a look, will be enough to decide whether I enter your father's house this evening or never.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I was within and without, simultaneously enchanted and repelled by the inexhaustible variety of life.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Success makes so many people hate you. I wish it wasn't that way. It would be wonderful to enjoy success without seeing envy in the eyes of those around you.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Real loneliness is not necessarily limited to when you are alone.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "America will never be destroyed from the outside. If we falter and lose our freedoms, it will be because we destroyed ourselves.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "Keep in mind that people change, but the past doesn't.",
	},
	{
		Author: "Nick HornbyHigh Fidelity",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422915487p2/2929.jpg",
		Content: "People worry about kids playing with guns, and teenagers watching violent videos; we are scared that some sort of culture of violence will take them over. Nobody worries about kids listening to thousands - literally thousands - of songs about broken hearts and rejection and pain and misery and loss.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Promise me you'll never forget me because if I thought you would, I'd never leave.",
	},
	{
		Author: "Anne Lamott",
		AuthorAvatar: "https://images.gr-assets.com/authors/1489601640p2/7113.jpg",
		Content: "Lighthouses don’t go running all over an island looking for boats to save; they just stand there shining.",
	},
	{
		Author: "Barbara W. Tuchman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1229046503p2/137261.jpg",
		Content: "Books are the carriers of civilization...They are companions, teachers, magicians, bankers of the treasures of the mind. Books are humanity in print.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Keep smiling, because life is a beautiful thing and there's so much to smile about.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I could not tell you if I loved you the first moment I saw you, or if it was the second or third or fourth. But I remember the first moment I looked at you walking toward me and realized that somehow the rest of the world seemed to vanish when I was with you.",
	},
	{
		Author: "Charles DickensGreat Expectations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "Suffering has been stronger than all other teaching, and has taught me to understand what your heart used to be. I have been bent and broken, but - I hope - into a better shape.",
	},
	{
		Author: "Irina Dunn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1297526928p2/786140.jpg",
		Content: "A woman without a man is like a fish without a bicycle.",
	},
	{
		Author: "Jimi Hendrix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207679011p2/7268.jpg",
		Content: "Knowledge speaks, but wisdom listens",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "A dream you dream alone is only a dream. A dream you dream together is reality.",
	},
	{
		Author: "Mark TwainThe Innocents Abroad/Roughing It",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Travel is fatal to prejudice, bigotry, and narrow-mindedness, and many of our people need it sorely on these accounts. Broad, wholesome, charitable views of men and things cannot be acquired by vegetating in one little corner of the earth all one's lifetime.",
	},
	{
		Author: "Sarah DessenKeeping the Moon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "You should never be surprised when someone treats you with respect, you should expect it.",
	},
	{
		Author: "Patrick RothfussThe Name of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351307341p2/108424.jpg",
		Content: "Words are pale shadows of forgotten names. As names have power, words have power. Words can light fires in the minds of men. Words can wring tears from the hardest hearts.",
	},
	{
		Author: "Roald Dahl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1311554908p2/4273.jpg",
		Content: "So please, oh please, we beg, we pray, go throw your TV set away, and in its place you can install, a lovely bookshelf on the wall.",
	},
	{
		Author: "Jodi Picoult",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "Once you had put the pieces back together, even though you may look intact, you were never quite the same as you'd been before the fall.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "It's all in the view. That's what I mean about forever, too. For any one of us our forever could end in an hour, or a hundred years from now. You never know for sure, so you'd better make every second count.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "And he took her in his arms and kissed her under the sunlit sky, and he cared not that they stood high upon the walls in the sight of many.",
	},
	{
		Author: "Haruki Murakami",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Anyone who falls in love is searching for the missing pieces of themselves. So anyone who's in love gets sad when they think of their lover. It's like stepping back inside a room you have fond memories of, one you haven't seen in a long time.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Then I realize what it is. It's him. Something about him makes me feel like I am about to fall. Or turn to liquid. Or burst into flames.",
	},
	{
		Author: "Chuck KlostermanKilling Yourself to Live: 85% of a True Story",
		AuthorAvatar: "https://images.gr-assets.com/authors/1336060202p2/375.jpg",
		Content: "Art and love are the same thing: It’s the process of seeing yourself in things that are not you.",
	},
	{
		Author: "Dylan ThomasIn Country Sleep, and Other Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1323281470p2/57189.jpg",
		Content: "Do not go gentle into that good night.Rage, rage against the dying of the light.",
	},
	{
		Author: "Charles DickensOliver Twist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "There are books of which the backs and covers are by far the best parts.",
	},
	{
		Author: "Stephen KingOn Writing: A Memoir of the Craft",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "you can, you should, and if you’re brave enough to start, you will.",
	},
	{
		Author: "Zora Neale HurstonTheir Eyes Were Watching God",
		AuthorAvatar: "https://images.gr-assets.com/authors/1194472605p2/15151.jpg",
		Content: "There are years that ask questions and years that answer.",
	},
	{
		Author: "William ShakespeareA Midsummer Night's Dream",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Though she be but little, she is fierce!",
	},
	{
		Author: "Jeaniene FrostHalfway to the Grave",
		AuthorAvatar: "https://images.gr-assets.com/authors/1310062695p2/669810.jpg",
		Content: "I'm saying that I'm a moody, insecure, narrow-minded, jealous, borderline homicidal bitch, and I want you to promise me that you're okay with that, because it's who I am, and you're what I need.",
	},
	{
		Author: "Dorothy Parker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820565p2/24956.jpg",
		Content: "Beauty is only skin deep, but ugly goes clean to the bone.",
	},
	{
		Author: "Douglas AdamsMostly Harmless",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "A common mistake that people make when trying to design something completely foolproof is to underestimate the ingenuity of complete fools.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "That's why when major badasses greet each other in movies, they don't say anything, they just nod. The nod means, 'I' am a badass, and I recognize that you, too, are a badass,' but they don't say anything because they're Wolverine and Magneto and it would mess up their vibe to explain.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "My Best Friend is a person who will give me a book I have not read.",
	},
	{
		Author: "Isabel Allende",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341879973p2/2238.jpg",
		Content: "The library is inhabited by spirits that come out of the pages at night.",
	},
	{
		Author: "Henry David Thoreau",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "I went to the woods because I wished to live deliberately, to front only the essential facts of life, and see if I could not learn what it had to teach, and not, when I came to die, discover that I had not lived. I did not wish to live what was not life, living is so dear; nor did I wish to practise resignation, unless it was quite necessary. I wanted to live deep and suck out all the marrow of life, to live so sturdily and Spartan-like as to put to rout all that was not life, to cut a broad swath and shave close, to drive life into a corner, and reduce it to its lowest terms...",
	},
	{
		Author: "Margaret Lee Runbeck",
		AuthorAvatar: "",
		Content: "Silence make the real conversations between friends. Not the saying, but the never needing to say that counts.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "Rivers know this: there is no hurry. We shall get there some day.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It wouldn't be my move,\" Jace agreed. \"First the candy and flowers, then the apology letters, then the ravenous demon hordes. In that order.",
	},
	{
		Author: "Jean-Paul Sartre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475567078p2/1466.jpg",
		Content: "If you're lonely when you're alone, you're in bad company.",
	},
	{
		Author: "Dan BrownAngels \u0026 Demons",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399396714p2/630.jpg",
		Content: "Science and religion are not at odds. Science is simply too young to understand.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I don't know that love changes. People change. Circumstances change.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "The purpose of life is not to be happy. It is to be useful, to be honorable, to be compassionate, to have it make some difference that you have lived and lived well.",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "It is the time you have wasted for your rose that makes your rose so important.",
	},
	{
		Author: "Margaret Mead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198589352p2/61107.jpg",
		Content: "I was wise enough never to grow up, while fooling people into believing I had.",
	},
	{
		Author: "David LevithanWill Grayson, Will Grayson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1426529210p2/11664.jpg",
		Content: "I am constantly torn between killing myself and killing everyone around me.",
	},
	{
		Author: "Shmuley Boteach",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271797673p2/9383.jpg",
		Content: "There is greatness in doing something you hate for the sake of someone you love.",
	},
	{
		Author: "Eckhart Tolle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1505974741p2/4493.jpg",
		Content: "The past has no power over the present moment.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "May the forces of evil become confused on the way to your house.",
	},
	{
		Author: "Charles BukowskiLove Is a Dog from Hell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "I loved you like a man loves a woman he never touches, only writes to, keeps little photographs of.",
	},
	{
		Author: "Rita Mae Brown",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209493600p2/23511.jpg",
		Content: "I think the reward for conformity is that everyone likes you except yourself.",
	},
	{
		Author: "Hunter S. ThompsonThe Proud Highway: Saga of a Desperate Southern Gentleman, 1955-1967",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206560814p2/5237.jpg",
		Content: "We are all alone, born alone, die alone, and—in spite of True Romance magazines—we shall all someday look back on our lives and see that, in spite of our company, we were alone the whole way. I do not say lonely—at least, not all the time—but essentially, and finally, alone. This is what makes your self-respect so important, and I don't see how you can respect yourself if you must look in the hearts and minds of others for your happiness.",
	},
	{
		Author: "Bette Midler",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207797480p2/548179.jpg",
		Content: "The worst part of success is trying to find someone who is happy for you.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Knock, And He'll open the doorVanish, And He'll make you shine like the sunFall, And He'll raise you to the heavensBecome nothing, And He'll turn you into everything.",
	},
	{
		Author: "Carlos Ruiz ZafónThe Shadow of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444199853p2/815.jpg",
		Content: "Fools talk, cowards are silent, wise men listen.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "The beauty of a woman is not in the clothes she wears, the figure that she carries, or the way she combs her hair. The beauty of a woman is seen in her eyes, because that is the doorway to her heart, the place where love resides. True beauty in a woman is reflected in her soul. It's the caring that she lovingly gives, the passion that she shows \u0026 the beauty of a woman only grows with passing years.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "In a way, it's nice to know that there are Greek gods out there, because you have somebody to blame when things go wrong. For instance, when you're walking away from a bus that's just been attacked by monster hags and blown up by lightning, and it's raining on top of everything else, most people might think that's just really bad luck; when you're a half-blood, you understand that some devine force is really trying to mess up your day.",
	},
	{
		Author: "Anna Quindlen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387414577p2/3500.jpg",
		Content: "The thing that is really hard, and really amazing, is giving up on being perfect and beginning the work of becoming yourself.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Where there is love, there is often also hate. They can exist side by side.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "There is no exquisite beauty… without some strangeness in the proportion.",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "We can never judge the lives of others, because each person knows only their own pain and renunciation. It's one thing to feel that you are on the right path, but it's another to think that yours is the only path.",
	},
	{
		Author: "StendhalThe Red and the Black",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373142965p2/1481537.jpg",
		Content: "A good book is an event in my life.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "But it is the nature of stars to cross, and never was Shakespeare more wrong than when he has Cassius note, ‘The fault, dear Brutus, is not in our stars / But in ourselves.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Politeness is deception in pretty packaging.",
	},
	{
		Author: "Jane Austen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "Ah! There is nothing like staying at home, for real comfort.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Always do what is right. It will gratify half of mankind and astound the other.",
	},
	{
		Author: "Dalai Lama XIV",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241989386p2/570218.jpg",
		Content: "Love is the absence of judgment.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Nitwit! Blubber! Oddment! Tweak!",
	},
	{
		Author: "Joseph HellerCatch-22",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197308614p2/3167.jpg",
		Content: "He was going to live forever, or die in the attempt.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Would it save you a lot of time if I just gave up and went mad now?",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I want to grow old without facelifts. I want to have the courage to be loyal to the face I have made.",
	},
	{
		Author: "C.S. LewisMere Christianity",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I am trying here to prevent anyone saying the really foolish thing that people often say about Him: I’m ready to accept Jesus as a great moral teacher, but I don’t accept his claim to be God. That is the one thing we must not say. A man who was merely a man and said the sort of things Jesus said would not be a great moral teacher. He would either be a lunatic — on the level with the man who says he is a poached egg — or else he would be the Devil of Hell. You must make your choice. Either this man was, and is, the Son of God, or else a madman or something worse. You can shut him up for a fool, you can spit at him and kill him as a demon or you can fall at his feet and call him Lord and God, but let us not come with any patronizing nonsense about his being a great human teacher. He has not left that open to us. He did not intend to.",
	},
	{
		Author: "José N. HarrisMI VIDA: A Story of Faith, Hope and Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414207447p2/4631437.jpg",
		Content: "Tears shed for another person are not a sign of weakness. They are a sign of a pure heart.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "One thing you can't hide - is when you're crippled inside.",
	},
	{
		Author: "Paulo CoelhoEleven Minutes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "When I had nothing to lose, I had everything. When I stopped being who I am, I found myself.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "If you don't read the newspaper, you're uninformed. If you read the newspaper, you're mis-informed.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "There are some things you can't share without ending up liking each other, and knocking out a twelve-foot mountain troll is one of them.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "When I am with you, we stay up all night.When you're not here, I can't go to sleep.Praise God for those two insomnias!And the difference between them.",
	},
	{
		Author: "Anne Lamott",
		AuthorAvatar: "https://images.gr-assets.com/authors/1489601640p2/7113.jpg",
		Content: "You will lose someone you can’t live without,and your heart will be badly broken, and the bad news is that you never completely get over the loss of your beloved. But this is also the good news. They live forever in your broken heart that doesn’t seal back up. And you come through. It’s like having a broken leg that never heals perfectly—that still hurts when the weather gets cold, but you learn to dance with the limp.",
	},
	{
		Author: "غسان كنفاني",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202528011p2/72154.jpg",
		Content: "!لك شيء في هذا العالم.. فقم",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Some lose all mind and become soul,insane.some lose all soul and become mind, intellectual.some lose both and become accepted",
	},
	{
		Author: "E.E. Cummings100 Selected Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "For whatever we lose (like a you or a me),It's always our self we find in the sea.",
	},
	{
		Author: "George R.R. MartinA Song of Ice and Fire, 5-Book Boxed Set: A Game of Thrones, A Clash of Kings, A Storm of Swords, A Feast for Crows, A Dance with Dragons",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Never forget what you are, for surely the world will not. Make it your strength. Then it can never be your weakness.",
	},
	{
		Author: "Voltaire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393683411p2/5754446.jpg",
		Content: "It is forbidden to kill; therefore all murderers are punished unless they kill in large numbers and to the sound of trumpets.",
	},
	{
		Author: "Walt Whitman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392303683p2/1438.jpg",
		Content: "What is that you express in your eyes? It seems to me more than all the print I have read in my life.",
	},
	{
		Author: "A.J. Cronin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211475194p2/807271.jpg",
		Content: "Worry never robs tomorrow of its sorrow, but only saps today of its strength.",
	},
	{
		Author: "Fyodor DostoyevskyCrime and Punishment",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "The darker the night, the brighter the stars,  The deeper the grief, the closer is God!",
	},
	{
		Author: "Jojo MoyesMe Before You",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400624880p2/281810.jpg",
		Content: "You only get one life. It's actually your duty to live it as fully as possible.",
	},
	{
		Author: "أحلام مستغانميذاكرة الجسد",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351458215p2/1109606.jpg",
		Content: "أحسد الأطفال الرضّع، لأنهم يملكون وحدهم حق الصراخ والقدرة عليه، قبل أن تروض الحياة حبالهم الصوتية، وتعلِّمهم الصمت",
	},
	{
		Author: "Sherrilyn KenyonAcheron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1382066830p2/4430.jpg",
		Content: "Life isn't finding shelter in the storm. It's about learning to dance in the rain.",
	},
	{
		Author: "Laurie Halse AndersonSpeak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376224335p2/10003.jpg",
		Content: "You have to know what you stand for, not just what you stand against.",
	},
	{
		Author: "Steve Maraboli",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "Letting go means to come to the realization that some people are a part of your history, but not a part of your destiny.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Every time you smile at someone, it is an action of love, a gift to that person, a beautiful thing.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "No. I can survive well enough on my own— if given the proper reading material.",
	},
	{
		Author: "مصطفى محمودلغز الموت",
		AuthorAvatar: "https://images.gr-assets.com/authors/1458565195p2/871358.jpg",
		Content: "و كلما أمسكت بحالة من حالاتي و قلت هذا هو أنا .. ما تلبث هذه الحالة أن تفلت من أصابعي و تحل محلها حالة أخرى .. هي أنا .. أيضاً..",
	},
	{
		Author: "Veronica RothAllegiant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I suppose a fire that burns that bright is not meant to last.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I could easily forgive his pride, if he had not mortified mine.",
	},
	{
		Author: "Ralph Waldo EmersonSelf Reliance",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Is it so bad, then, to be misunderstood? Pythagoras was misunderstood, and Socrates, and Jesus, and Luther, and Copernicus, and Galileo, and Newton, and every pure and wise spirit that ever took flesh. To be great is to be misunderstood.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Time is an illusion. Lunchtime doubly so.",
	},
	{
		Author: "Kurt Vonnegut",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "Hello babies. Welcome to Earth. It's hot in the summer and cold in the winter. It's round and wet and crowded. On the outside, babies, you've got a hundred years here. There's only one rule that I know of, babies-\"God damn it, you've got to be kind.",
	},
	{
		Author: "E.L. JamesFifty Shades of Grey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308409727p2/4725841.jpg",
		Content: "- \"Why don't you like to be touched?\" - \"Because I'm fifty shades of fucked-up, Anastasia",
	},
	{
		Author: "Guy de Maupassant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188821941p2/18791.jpg",
		Content: "Our memory is a more perfect world than the universe: it gives back life to those who no longer exist.",
	},
	{
		Author: "S.E. HintonThe Outsiders",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206505616p2/762707.jpg",
		Content: "Stay gold, Ponyboy, stay gold.",
	},
	{
		Author: "Cynthia Ozick",
		AuthorAvatar: "https://images.gr-assets.com/authors/1255219148p2/43530.jpg",
		Content: "We take for granted the very things that most deserve our gratitude.",
	},
	{
		Author: "Joseph Campbell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114498p2/20105.jpg",
		Content: "We must be willing to let go of the life we planned so as to have the life that is waiting for us.",
	},
	{
		Author: "Oscar WildeThe Importance of Being Earnest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "All women become like their mothers. That is their tragedy. No man does, and that is his.",
	},
	{
		Author: "John SteinbeckThe Grapes of Wrath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182118389p2/585.jpg",
		Content: "There ain't no sin and there ain't no virtue. There's just stuff people do.",
	},
	{
		Author: "Lewis CarrollAlice's Adventures in Wonderland \u0026 Through the Looking-Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "If I had a world of my own, everything would be nonsense. Nothing would be what it is, because everything would be what it isn't. And contrary wise, what is, it wouldn't be. And what it wouldn't be, it would. You see?",
	},
	{
		Author: "Nicholas SparksA Walk to Remember",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Without suffering, there'd be no compassion.",
	},
	{
		Author: "Salvador Dalí",
		AuthorAvatar: "https://images.gr-assets.com/authors/1218295974p2/165858.jpg",
		Content: "Have no fear of perfection - you'll never reach it.",
	},
	{
		Author: "Sharon Salzberg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1441077018p2/17208.jpg",
		Content: "You yourself, as much as anybody in the entire universe, deserve your love and affection",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "There is more to sex appeal than just measurements. I don't need a bedroom to prove my womanliness. I can convey just as much sex appeal, picking apples off a tree or standing in the rain.",
	},
	{
		Author: "Zelda FitzgeraldThe Collected Writings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1427041106p2/28243.jpg",
		Content: "She refused to be bored chiefly because she wasn't boring.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "I cannot believe in a God who wants to be praised all the time.",
	},
	{
		Author: "Lady Gaga",
		AuthorAvatar: "https://images.gr-assets.com/authors/1304697345p2/2945649.jpg",
		Content: "Trust is like a mirror, you can fix it if it's broken, but you can still see the crack in that mother fucker's reflection.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "I have come to accept the feeling of not knowing where I am going. And I have trained myself to love it. Because it is only when we are suspended in mid-air with no landing in sight, that we force our wings to unravel and alas begin our flight. And as we fly, we still may not know where we are going to. But the miracle is in the unfolding of the wings. You may not know where you're going, but you know that so long as you spread your wings, the winds will carry you.",
	},
	{
		Author: "Lao Tzu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "A good traveler has no fixed plans and is not intent on arriving.",
	},
	{
		Author: "Laurie Halse AndersonSpeak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376224335p2/10003.jpg",
		Content: "When people don't express themselves, they die one piece at a time.",
	},
	{
		Author: "Jennifer L. ArmentroutObsidian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1290636579p2/4476934.jpg",
		Content: "Beautiful face. Beautiful body. Horrible attitude. It was the holy trinity of hot boys.",
	},
	{
		Author: "George R.R. Martin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "The best fantasy is written in the language of dreams. It is alive as dreams are alive, more real than real ... for a moment at least ... that long magic moment before we wake. Fantasy is silver and scarlet, indigo and azure, obsidian veined with gold and lapis lazuli. Reality is plywood and plastic, done up in mud brown and olive drab. Fantasy tastes of habaneros and honey, cinnamon and cloves, rare red meat and wines as sweet as summer. Reality is beans and tofu, and ashes at the end. Reality is the strip malls of Burbank, the smokestacks of Cleveland, a parking garage in Newark. Fantasy is the towers of Minas Tirith, the ancient stones of Gormenghast, the halls of Camelot. Fantasy flies on the wings of Icarus, reality on Southwest Airlines. Why do our dreams become so much smaller when they finally come true? We read fantasy to find the colors again, I think. To taste strong spices and hear the songs the sirens sang. There is something old and true in fantasy that speaks to something deep within us, to the child who dreamt that one day he would hunt the forests of the night, and feast beneath the hollow hills, and find a love to last forever somewhere south of Oz and north of Shangri-La. They can keep their heaven. When I die, I'd sooner go to middle Earth.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Creativity is intelligence having fun.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Youth can not know how age thinks and feels. But old men are guilty if they forget what it was to be young.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "But in the end one needs more courage to live than to kill himself.",
	},
	{
		Author: "Jay AsherThirteen Reasons Why",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243931536p2/569269.jpg",
		Content: "You can't stop the futureYou can't rewind the pastThe only way to learn the secret...is to press play.",
	},
	{
		Author: "Neil GaimanSmoke and Mirrors: Short Fiction and Illusions",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "When I was a child, adults would tell me not to make things up, warning me of what would happen if I did. As far as I can tell so far, it seems to involve lots of foreign travel and not having to get up too early in the morning.",
	},
	{
		Author: "John Muir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398092241p2/5297.jpg",
		Content: "The clearest way into the Universe is through a forest wilderness.",
	},
	{
		Author: "Temple Grandin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241222068p2/1567.jpg",
		Content: "I am different, not less.",
	},
	{
		Author: "Dr. SeussOne Fish, Two Fish, Red Fish, Blue Fish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "From there to here, from here to there, funny things are everywhere!",
	},
	{
		Author: "Steve Wozniak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376585301p2/3767.jpg",
		Content: "If you love what you do and are willing to do what it takes, it's within your reach. And it'll be worth every minute you spend alone at night, thinking and thinking about what it is you want to design or build. It'll be worth it, I promise.",
	},
	{
		Author: "George Gordon Byron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357459330p2/44407.jpg",
		Content: "And thus the heart will break, yet brokenly live on.",
	},
	{
		Author: "William ShakespeareThe Taming of the Shrew",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "My tongue will tell the anger of my heart, or else my heart concealing it will break.",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "A child can teach an adult three things: to be happy for no reason, to always be busy with something, and to know how to demand with all his might that which he desires.",
	},
	{
		Author: "Marguerite Duras",
		AuthorAvatar: "https://images.gr-assets.com/authors/1518609150p2/163.jpg",
		Content: "Our mothers always remain the strangest, craziest people we've ever met.",
	},
	{
		Author: "Edward Gorey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200338278p2/21578.jpg",
		Content: "The helpful thought for which you lookIs written somewhere in a book.",
	},
	{
		Author: "Richelle MeadLast Sacrifice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Ah, my daughter,ʺ he said. ʺEighteen, and already youʹve been accused of murder, aided felons, and acquired a death count higher than most guardians will ever see.ʺ He paused. ʺI couldnʹt be prouder.",
	},
	{
		Author: "Lauren OliverBefore I Fall",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1431467822p2/277505.jpg",
		Content: "So many things become beautiful when you really look.",
	},
	{
		Author: "D.H. LawrenceLady Chatterley's Lover",
		AuthorAvatar: "https://images.gr-assets.com/authors/1375811585p2/17623.jpg",
		Content: "We've got to live, no matter how many skies have fallen.",
	},
	{
		Author: "Kate DiCamilloThe Miraculous Journey of Edward Tulane",
		AuthorAvatar: "https://images.gr-assets.com/authors/1402414985p2/13663.jpg",
		Content: "Open your heart. Someone will come. Someone will come for you. But first you must open your heart.",
	},
	{
		Author: "Christopher Paolini",
		AuthorAvatar: "https://images.gr-assets.com/authors/1412963950p2/8349.jpg",
		Content: "Because you can't argue with all the fools in the world. It's easier to let them have their way, then trick them when they're not paying attention.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It's not gray,\" Clary felt compelled to point out. \"It's green.\"\"If there was such a thing as terminal literalism, you'd have died in childhood,\" said Jace.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "She leaned down and looked at his lifeless face and Leisel kissed her best friend, Rudy Steiner, soft and true on his lips. He tasted dusty and sweet. He tasted like regret in the shadows of trees and in the glow of the anarchist's suit collection. She kissed him long and soft, and when she pulled herself away, she touched his mouth with her fingers...She did not say goodbye. She was incapable, and after a few more minutes at his side, she was able to tear herself from the ground. It amazes me what humans can do, even when streams are flowing down their faces and they stagger on...",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Did you ever think that in a past life Alec was an old woman with ninety cats who was always yelling at the neighborhood kids to get off her lawn? Because I do,",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Must you go? I was rather hoping you'd stay and be a ministering angel, but if you must go, you must.\"\"I'll stay,\" Will said a bit crossly, and threw himself down in the armchair Tessa had just vacated. \"I can minister angelically.\"\"None too convincingly. And you're not as pretty to look at as Tessa is,\" Jem said, closing his eyes as he leaned back against the pillow.\"How rude. Many who have gazed upon me have compared the experience to gazing at the radiance of the sun.\"Jem still had his eyes closed. \"If they mean it gives you a headache, they aren't wrong.",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "The best books... are those that tell you what you know already.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "Let us learn to show our friendship for a man when he is alive and not after he is dead.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Honestly, if you were any slower, you’d be going backward.",
	},
	{
		Author: "John SteinbeckEast of Eden",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182118389p2/585.jpg",
		Content: "I believe a strong woman may be stronger than a man, particularly if she happens to have love in her heart. I guess a loving woman is indestructible.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Yesterday is gone. Tomorrow has not yet come. We have only today. Let us begin.",
	},
	{
		Author: "Barbara KingsolverThe Poisonwood Bible",
		AuthorAvatar: "https://images.gr-assets.com/authors/1350499031p2/3541.jpg",
		Content: "Don’t try to make life a mathematics problem with yourself in the center and everything coming out equal. When you’re good, bad things can still happen. And if you’re bad, you can still be lucky.",
	},
	{
		Author: "Carl SaganCosmos",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475953320p2/10538.jpg",
		Content: "Every one of us is, in the cosmic perspective, precious. If a human disagrees with you, let him live. In a hundred billion galaxies, you will not find another.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "And when at last you find someone to whom you feel you can pour out your soul, you stop in shock at the words you utter— they are so rusty, so ugly, so meaningless and feeble from being kept in the small cramped dark inside you so long.",
	},
	{
		Author: "Kurt VonnegutA Man Without a Country",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "Here is a lesson in creative writing. First rule: Do not use semicolons. They are transvestite hermaphrodites representing absolutely nothing. All they do is show you've been to college.",
	},
	{
		Author: "Charles M. Schulz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590750p2/209672.jpg",
		Content: "What's the good of living if you don't try a few things?",
	},
	{
		Author: "Margaret Mead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198589352p2/61107.jpg",
		Content: "Children must be taught how to think, not what to think.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I thought I'd lie on the floor and writhe in pain for a while,\" he grunted, \"It relaxes me.\"\"It does? Oh - you're being sarcastic. That's a good sign probably.",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "Out of suffering have emerged the strongest souls; the most massive characters are seared with scars.",
	},
	{
		Author: "Danielle SteelThe Gift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340133722p2/14255.jpg",
		Content: "Maybe some people just aren't meant to be in our lives forever. Maybe some people are just passing through. It's like some people just come through our lives to bring us something: a gift, a blessing, a lesson we need to learn. And that's why they're here. You'll have that gift forever.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The world is a dangerous place to live, not because of the people who are evil, but because of the people who don't do anything about it.",
	},
	{
		Author: "Thomas MertonNo Man Is an Island",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201097624p2/1711.jpg",
		Content: "Art enables us to find ourselves and lose ourselves at the same time.",
	},
	{
		Author: "Marion Zimmer BradleyThe Forest House",
		AuthorAvatar: "https://images.gr-assets.com/authors/1305483488p2/4841825.jpg",
		Content: "Remain true to yourself, child. If you know your own heart, you will always have one friend who does not lie.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "No matter what he does, every person on earth plays a central role in the history of the world. And normally he doesn't know it.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The saddest people I've ever met in life are the ones who don't care deeply about anything at all. Passion and satisfaction go hand in hand, and without them, any happiness is only temporary, because there's nothing to make it last.",
	},
	{
		Author: "Dr. SeussOh, the Places You'll Go! and The Lorax",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "All alone! Whether you like it or not, alone is something you'll be quite a lot!",
	},
	{
		Author: "Kazuo Ishiguro",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424906625p2/4280.jpg",
		Content: "There was another life that I might have had, but I am having this one.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "I have absolutely no pleasure in the stimulants in which I sometimes so madly indulge. It has not been in the pursuit of pleasure that I have periled life and reputation and reason. It has been the desperate attempt to escape from torturing memories, from a sense of insupportable loneliness and a dread of some strange impending doom.",
	},
	{
		Author: "Cassandra ClareCity of Fallen Angels",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You're just worried they'll hire a male instructor and he'll be hotter than you.\" Jace's eyebrows went up. \"Hotter than me?\" \"It could happen,\" Clary said, \"You know, theoretically.\"\"Theoretically the planet could suddenly crack in half, leaving me on one side and you on the other, forever and tragically parted, but I'm not worried about that either. Some things,\" Jace said, with his customary crooked smile, \"are just too unlikely to dwell upon.",
	},
	{
		Author: "Alan W. Watts",
		AuthorAvatar: "https://images.gr-assets.com/authors/1427892345p2/1501668.jpg",
		Content: "Trying to define yourself is like trying to bite your own teeth.",
	},
	{
		Author: "L.M. MontgomeryAnne of Green Gables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188896723p2/5350.jpg",
		Content: "It's been my experience that you can nearly always enjoy things if you make up your mind firmly that you will.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "let me live, love, and say it well in good sentences",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Most books on witchcraft will tell you that witches work naked. This is because most books on witchcraft were written by men.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1504718140p2/5722.jpg",
		Content: "Anyway, I keep picturing all these little kids playing some game in this big field of rye and all. Thousands of little kids, and nobody's around - nobody big, I mean - except me. And I'm standing on the edge of some crazy cliff. What I have to do, I have to catch everybody if they start to go over the cliff - I mean if they're running and they don't look where they're going I have to come out from somewhere and catch them. That's all I do all day. I'd just be the catcher in the rye and all. I know it's crazy, but that's the only thing I'd really like to be.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Everyone is a moon, and has a dark side which he never shows to anybody.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Your beliefs become your thoughts, Your thoughts become your words, Your words become your actions, Your actions become your habits, Your habits become your values, Your values become your destiny.",
	},
	{
		Author: "William Blake",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199069675p2/13453.jpg",
		Content: "It is easier to forgive an enemy than to forgive a friend.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "When did we see each other face-to-face? Not until you saw into my cracks and I saw into yours. Before that, we were just looking at ideas of each other, like looking at your window shade but never seeing inside. But once the vessel cracks, the light can get in. The light can get out.",
	},
	{
		Author: "Nicholas Sparks",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "If conversation was the lyrics, laughter was the music, making time spent together a melody that could be replayed over and over without getting stale.",
	},
	{
		Author: "Thomas MertonThe Way of Chuang Tzu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201097624p2/1711.jpg",
		Content: "The beginning of love is the will to let those we love be perfectly themselves, the resolution not to twist them to fit our own image.",
	},
	{
		Author: "Gordon B. Hinckley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202224343p2/313356.jpg",
		Content: "Try a little harder to be a little better.",
	},
	{
		Author: "Edward J. Stieglitz",
		AuthorAvatar: "",
		Content: "The important thing to you is not how many years in your life, but how much life in your years!",
	},
	{
		Author: "Alan BennettThe Uncommon Reader",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234747177p2/11781.jpg",
		Content: "What she was finding also was how one book led to another, doors kept opening wherever she turned and the days weren't long enough for the reading she wanted to do.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "God can't give us peace and happiness apart from Himself because there is no such thing.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "Success is stumbling from failure to failure with no loss of enthusiasm.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You could have fooled me. Everytime I called you, Luke said you were sick. I figured you were avoiding me. Again.\"\"I wasn't. I did want to talk to you. I've been thinking about you all the time.\"\"I've been thinking about you, too.\"\"I really was sick. I swear. I almost died back there on the ship, you know.\"\"I know. Everytime you almost die, I almost die myself.",
	},
	{
		Author: "Confucius",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407613261p2/15321.jpg",
		Content: "Everything has beauty, but not everyone sees it.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You'd think people had better things to gossip about,\" said Ginny as she sat on the common room floor, leaning against Harry’s legs and reading the Daily Prophet. \"Three Dementor attacks in a week, and all Romilda Vane does is ask me if it’s true you’ve got a Hippogriff tattooed across your chest.\"Ron and Hermione both roared with laughter. Harry ignored them.What did you tell her?\"I told her it's a Hungarian Horntail,\" said Ginny, turning a page of the newspaper idly. \"Much more macho.\"Thanks,\" said Harry, grinning. \"And what did you tell her Ron’s got?\"A Pygmy Puff, but I didn’t say where.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "It's funny how humans can wrap their mind around things and fit them into their version of reality.",
	},
	{
		Author: "Katherine Mansfield",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216670886p2/45712.jpg",
		Content: "The pleasure of all reading is doubled when one lives with another who shares the same books.",
	},
	{
		Author: "William GoldingLord of the Flies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198342496p2/306.jpg",
		Content: "Maybe there is a beast… maybe it's only us.",
	},
	{
		Author: "Ursula K. Le GuinThe Wave in the Mind: Talks \u0026 Essays on the Writer, the Reader \u0026 the Imagination",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244291425p2/874602.jpg",
		Content: "People who deny the existence of dragons are often eaten by dragons. From within.",
	},
	{
		Author: "William ShakespeareHamlet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "To be, or not to be: that is the question:Whether 'tis nobler in the mind to sufferThe slings and arrows of outrageous fortune,Or to take arms against a sea of troubles,And by opposing end them? To die: to sleep;No more; and by a sleep to say we endThe heart-ache and the thousand natural shocksThat flesh is heir to, 'tis a consummationDevoutly to be wish'd. To die, to sleep;To sleep: perchance to dream: ay, there's the rub;For in that sleep of death what dreams may comeWhen we have shuffled off this mortal coil,Must give us pause: there's the respectThat makes calamity of so long life;For who would bear the whips and scorns of time,The oppressor's wrong, the proud man's contumely,The pangs of despised love, the law's delay,The insolence of office and the spurnsThat patient merit of the unworthy takes,When he himself might his quietus makeWith a bare bodkin? who would fardels bear,To grunt and sweat under a weary life,But that the dread of something after death,The undiscover'd country from whose bournNo traveller returns, puzzles the willAnd makes us rather bear those ills we haveThan fly to others that we know not of?Thus conscience does make cowards of us all;And thus the native hue of resolutionIs sicklied o'er with the pale cast of thought,And enterprises of great pith and momentWith this regard their currents turn awry,And lose the name of action.--Soft you now!The fair Ophelia! Nymph, in thy orisonsBe all my sins remember'd!",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "I like to see people reunited, I like to see people run to each other, I like the kissing and the crying, I like the impatience, the stories that the mouth can't tell fast enough, the ears that aren't big enough, the eyes that can't take in all of the change, I like the hugging, the bringing together, the end of missing someone.",
	},
	{
		Author: "John SteinbeckEast of Eden",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182118389p2/585.jpg",
		Content: "All great and precious things are lonely.",
	},
	{
		Author: "François Mauriac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419495337p2/61023.jpg",
		Content: "If you would tell me the heart of a man, tell me not what he reads, but what he rereads.",
	},
	{
		Author: "Haruki MurakamiSouth of the Border, West of the Sun",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Sometimes when I look at you, I feel I'm gazing at a distant star. It's dazzling, but the light is from tens of thousands of years ago.Maybe the star doesn't even exist any more. Yet sometimes that light seems more real to me than anything.",
	},
	{
		Author: "Rick Riordan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Dreams like a podcast,Downloading truth in my ears.They tell me cool stuff.\"\"Apollo?\" I guess, because I figured nobody else could make a haiku that bad.He put his finger to his lips. \"I'm incognito. Call me Fred.\"\"A god named Fred?",
	},
	{
		Author: "Robert Orben",
		AuthorAvatar: "https://images.gr-assets.com/authors/1250696742p2/1032064.jpg",
		Content: "Illegal aliens have always been a problem in the United States. Ask any Indian.",
	},
	{
		Author: "François de La Rochefoucauld",
		AuthorAvatar: "https://images.gr-assets.com/authors/1410969118p2/7428903.jpg",
		Content: "No persons are more frequently wrong, than those who will not admit they are wrong.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "The only way to get rid of temptation is to yield to it.",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "Who in the world am I? Ah, that's the great puzzle.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I don't want tea,\" said Clary, with muffled force. \"I want to find my mother. And then I want to find out who took her in the first place, and I want to kill them.\"\"Unfortunately,\" said Hodge, \"we're all out of bitter revenge at the moment, so it's either tea or nothing.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Give her hell from us, Peeves.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Pain insists upon being attended to. God whispers to us in our pleasures, speaks in our consciences, but shouts in our pains. It is his megaphone to rouse a deaf world.",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Don't waste your time with explanations: people only hear what they want to hear.",
	},
	{
		Author: "Paulo CoelhoBrida",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Nothing in the world is ever completely wrong. Even a stopped clock is right twice a day.",
	},
	{
		Author: "Lois McMaster BujoldA Civil Campaign",
		AuthorAvatar: "https://images.gr-assets.com/authors/1377313786p2/16094.jpg",
		Content: "Reputation is what other people know about you. Honor is what you know about yourself.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The trouble is not in dying for a friend, but in finding a friend worth dying for.",
	},
	{
		Author: "Lao Tzu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "Life is a series of natural and spontaneous changes. Don't resist them; that only creates sorrow. Let reality be reality. Let things flow naturally forward in whatever way they like.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Quotation is a serviceable substitute for wit.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "We didn't talk about anything heavy or light. We were just there together. And that was enough",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "A diary with no drawings of me in it? Where are the torrid fantasies? The romance covers?",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "Atheism turns out to be too simple. If the whole universe has no meaning, we should never have found out that it has no meaning...",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Really, the combination of the scabs and the ointment looks hideous. I can't help enjoying his distress.\"Poor Finnick. Is this the first time in your life you haven't looked pretty?\" I say.\"It must be. The sensation's completely new. How have you managed it all these years?\" he asks.\"Just avoid mirrors. You'll forget about it,\" I say.\"Not if I keep looking at you,\" he says.",
	},
	{
		Author: "Keri HulmeThe Bone People",
		AuthorAvatar: "https://images.gr-assets.com/authors/1240845441p2/7451.jpg",
		Content: "You want to know about anybody? See what books they read, and how they've been read...",
	},
	{
		Author: "Veronica RothAllegiant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I fell in love with him. But I don't just stay with him by default as if there's no one else available to me. I stay with him because I choose to, every day that I wake up, every day that we fight or lie to each other or disappoint each other. I choose him over and over again, and he chooses me.",
	},
	{
		Author: "F. Scott FitzgeraldThis Side of Paradise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "I'm not sentimental--I'm as romantic as you are. The idea, you know,is that the sentimental person thinks things will last--the romanticperson has a desperate confidence that they won't.",
	},
	{
		Author: "Maxine Hong Kingston",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202527632p2/17290.jpg",
		Content: "In a time of destruction, create something.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "Dreams do come true, if only we wish hard enough. You can have anything in life if you will sacrifice everything else for it.",
	},
	{
		Author: "E.Y. Harburg",
		AuthorAvatar: "",
		Content: "Somewhere over the rainbow, skies are blue, and the dreams that you dare to dream really do come true.",
	},
	{
		Author: "Desmond Tutu",
		AuthorAvatar: "https://images.gr-assets.com/authors/1251315285p2/5943.jpg",
		Content: "Do your little bit of good where you are; it's those little bits of good put together that overwhelm the world.",
	},
	{
		Author: "Aldous Huxley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "You shall know the truth and the truth shall make you mad.",
	},
	{
		Author: "Anne Brontë",
		AuthorAvatar: "https://images.gr-assets.com/authors/1219762839p2/8249.jpg",
		Content: "But he who dares not grasp the thorn Should never crave the rose.",
	},
	{
		Author: "Søren Kierkegaard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202410387p2/6172.jpg",
		Content: "People demand freedom of speech as a compensation for the freedom of thought which they seldom use.",
	},
	{
		Author: "Walt Whitman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392303683p2/1438.jpg",
		Content: "This is what you shall do; Love the earth and sun and the animals, despise riches, give alms to every one that asks, stand up for the stupid and crazy, devote your income and labor to others, hate tyrants, argue not concerning God, have patience and indulgence toward the people, take off your hat to nothing known or unknown or to any man or number of men, go freely with powerful uneducated persons and with the young and with the mothers of families, read these leaves in the open air every season of every year of your life, re-examine all you have been told at school or church or in any book, dismiss whatever insults your own soul, and your very flesh shall be a great poem and have the richest fluency not only in its words but in the silent lines of its lips and face and between the lashes of your eyes and in every motion and joint of your body.",
	},
	{
		Author: "Deepak Chopra",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400537940p2/138207.jpg",
		Content: "Every time you are tempted to react in the same old way, ask if you want to be a prisoner of the past or a pioneer of the future.",
	},
	{
		Author: "Stephanie PerkinsAnna and the French Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407443106p2/3095893.jpg",
		Content: "The more you know who you are, and what you want, the less you let things upset you.",
	},
	{
		Author: "Oscar WildeLady Windermere's Fan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "It is absurd to divide people into good and bad. People are either charming or tedious.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "We do not grow absolutely, chronologically. We grow sometimes in one dimension, and not in another; unevenly. We grow partially. We are relative. We are mature in one realm, childish in another. The past, present, and future mingle and pull us backward, forward, or fix us in the present. We are made up of layers, cells, constellations.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Are you implying that shreds of my reputation remain intact?\" Will demanded with mock horror. \"Clearly I have been doing something wrong. Or not something wrong, as the case may be.\" He banged on the side of the carriage. \"Thomas! We must away at once to the nearest brothel. I seek scandal and low companionship.",
	},
	{
		Author: "L. Frank BaumThe Wonderful Wizard of Oz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383720421p2/3242.jpg",
		Content: "There is no place like home.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "May the odds be ever in your favor!",
	},
	{
		Author: "Conan O'Brien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207596891p2/140050.jpg",
		Content: "If you work really hard, and you're kind, amazing things will happen.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "But who can say what's best? That's why you need to grab whatever chance you have of happiness where you find it, and not worry about other people too much. My experience tells me that we get no more than two or three such chances in a life time, and if we let them go, we regret it for the rest of our lives.",
	},
	{
		Author: "Confucius",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407613261p2/15321.jpg",
		Content: "Wheresoever you go, go with all your heart.",
	},
	{
		Author: "Neil GaimanFragile Things: Short Fictions and Wonders",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "She seems so cool, so focused, so quiet, yet her eyes remain fixed upon the horizon. You think you know all there is to know about her immediately upon meeting her, but everything you think you know is wrong. Passion flows through her like a river of blood. She only looked away for a moment, and the mask slipped, and you fell. All your tomorrows start here.",
	},
	{
		Author: "Carl Sandburg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1225504449p2/16380.jpg",
		Content: "Time is the coin of your life. You spend it. Do not allow others to spend it for you.",
	},
	{
		Author: "Ayn Rand",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "Learn to value yourself, which means: fight for your happiness.",
	},
	{
		Author: "Coco ChanelThe Gospel According to Coco Chanel: Life Lessons from the World's Most Elegant Woman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "A girl should be two things: who and what she wants.",
	},
	{
		Author: "Mary Oliver",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429857566p2/23988.jpg",
		Content: "Someone I loved once gave me a box full of darkness. It took me years to understand that this too, was a gift.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You're an idiot.\"\"I've never claimed to be otherwise.",
	},
	{
		Author: "Dr. SeussHow the Grinch Stole Christmas!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Then the Grinch thought of something he hadn't before! What if Christmas, he thought, doesn't come from a store. What if Christmas...perhaps...means a little bit more!",
	},
	{
		Author: "Beatrix Potter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201022492p2/11593.jpg",
		Content: "There is something delicious about writing the first words of a story. You never quite know where they'll take you.",
	},
	{
		Author: "Toni MorrisonSong of Solomon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494211316p2/3534.jpg",
		Content: "You wanna fly, you got to give up the shit that weighs you down.",
	},
	{
		Author: "Paulo CoelhoBrida",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "When you find your path, you must not be afraid. You need to have sufficient courage to make mistakes. Disappointment, defeat, and despair are the tools God uses to show us the way.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "Youth is happy because it has the capacity to see beauty. Anyone who keeps the ability to see beauty never grows old.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "The only way that we can live, is if we grow. The only way that we can grow is if we change. The only way that we can change is if we learn. The only way we can learn is if we are exposed. And the only way that we can become exposed is if we throw ourselves out into the open. Do it. Throw yourself.",
	},
	{
		Author: "Stephenie MeyerBreaking Dawn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "Did you know that 'I told you so' has a brother,Jacob?\" she asked cutting me off. \"His name is 'Shut the hell up'.",
	},
	{
		Author: "William ShakespeareRomeo and Juliet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "These violent delights have violent endsAnd in their triump die, like fire and powderWhich, as they kiss, consume",
	},
	{
		Author: "Lev GrossmanThe Magicians",
		AuthorAvatar: "https://images.gr-assets.com/authors/1386343699p2/142270.jpg",
		Content: "If there's a single lesson that life teaches us, it's that wishing doesn't make it so.",
	},
	{
		Author: "Umberto EcoThe Name of the Rose",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455915753p2/1730.jpg",
		Content: "Books are not made to be believed, but to be subjected to inquiry. When we consider a book, we mustn't ask ourselves what it says but what it means...",
	},
	{
		Author: "David MitchellCloud Atlas",
		AuthorAvatar: "https://images.gr-assets.com/authors/1409248688p2/6538289.jpg",
		Content: "A half-read book is a half-finished love affair.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "I wrote the story myself. It's about a girl who lost her reputation and never missed it.",
	},
	{
		Author: "Richelle MeadVampire Academy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "The spell. Victor said you had to want me... to care about me... for it to work.\" When he didn't say anything, I tried to grip his shirt, but my fingers were too weak. \"Did you? Did you want me?\"His words came out thickly. \"Yes, Roza. I did want you. I still do. I wish... we could be together.\"\"Then why did you lie to me?\"We reached the clinic, and he managed to open the door while still holding me. As soon as he stepped inside, he began yelling for help. \"Why did you lie?\" I murmured again.Still holding me in his arms, he looked down at me. I could hear voices and footsteps getting closer.\"Because we can't be together.\"\"Because of the age thing, right?\" I asked. \"Because you're my mentor?\"His fingertip gently wiped away a tear that had escaped down my cheek. \"That's part of it,\" he said. \"But also... well, you and I will both be Lissa's gaurdians someday. I need to protect her at all cost. If a pack of Strogoi come, I need to throw my body between them and her.\"I know that. Of course that's what you have to do.\" The black sparkles were dancing in front of my eyes again. I was fading out.\"No. If I let myself love you, I won't throw myself in front of her. I'll throw myself in front of you.",
	},
	{
		Author: "Aldous HuxleyBrave New World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "But I don't want comfort. I want God, I want poetry, I want real danger, I want freedom, I want goodness. I want sin.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "It is a good rule after reading a new book, never to allow yourself another new one till you have read an old one in between.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Still round the corner there may wait A new road or a secret gateAnd though I oft have passed them by A day will come at last when IShall take the hidden paths that run West of the Moon, East of the Sun.",
	},
	{
		Author: "Benjamin FranklinMemoirs of the life \u0026 writings of Benjamin Franklin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215314094p2/289513.jpg",
		Content: "They who can give up essential liberty to obtain a little temporary safety deserve neither liberty nor safety.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "When we are tired, we are attacked by ideas we conquered long ago.",
	},
	{
		Author: "William Wordsworth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288772244p2/64845.jpg",
		Content: "Fill your paper with the breathings of your heart.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "Don't bend; don't water it down; don't try to make it logical; don't edit your own soul according to the fashion. Rather, follow your most intense obsessions mercilessly.",
	},
	{
		Author: "Michael OndaatjeThe English Patient",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390116915p2/4030.jpg",
		Content: "She had always wanted words, she loved them; grew up on them. Words gave her clarity, brought reason, shape.",
	},
	{
		Author: "Charles BukowskiWomen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "being alone never felt right. sometimes it felt good, but it never felt right.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "One of the advantages of being disorganized is that one is always having surprising discoveries.",
	},
	{
		Author: "John GrishamThe Rainmaker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1413390525p2/721.jpg",
		Content: "Don't compromise yourself - you're all you have.",
	},
	{
		Author: "Arthur Schopenhauer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1404836700p2/11682.jpg",
		Content: "Compassion is the basis of morality.",
	},
	{
		Author: "Orhan PamukSnow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1423469681p2/1728.jpg",
		Content: "Happiness is holding someone in your arms and knowing you hold the whole world.",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "The most beautiful things in the world cannot be seen or touched, they are felt with the heart.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Let me give you a piece of advice. The handsome young fellow who's trying to rescue you from a hideous fate is never wrong. Not even if he says the sky is purple and made of hedgehogs.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I always channel my emotions into my work. That way, I don't hurt anyone but myself.",
	},
	{
		Author: "Winston S. Churchill",
		AuthorAvatar: "https://images.gr-assets.com/authors/1306133803p2/14033.jpg",
		Content: "Tact is the ability to tell someone to go to hell in such a way that they look forward to the trip.",
	},
	{
		Author: "Jim Morrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429658456p2/7855.jpg",
		Content: "Expose yourself to your deepest fear; after that, fear has no power, and the fear of freedom shrinks and vanishes. You are free.",
	},
	{
		Author: "Dodie SmithI Capture the Castle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1196206427p2/6208.jpg",
		Content: "There is only one page left to write on. I will fill it with words of only one syllable. I love. I have loved. I will love.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "A snowball in the face is surely the perfect beginning to a lasting friendship.",
	},
	{
		Author: "Honoré de Balzac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206567834p2/228089.jpg",
		Content: "Reading brings us unknown friends",
	},
	{
		Author: "Emma DonoghueRoom",
		AuthorAvatar: "https://images.gr-assets.com/authors/1298686263p2/23613.jpg",
		Content: "Scared is what you're feeling. Brave is what you're doing.",
	},
	{
		Author: "Eric RothThe Curious Case of Benjamin Button Screenplay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400170590p2/80495.jpg",
		Content: "For what it’s worth: it’s never too late or, in my case, too early to be whoever you want to be. There’s no time limit, stop whenever you want. You can change or stay the same, there are no rules to this thing. We can make the best or the worst of it. I hope you make the best of it. And I hope you see things that startle you. I hope you feel things you never felt before. I hope you meet people with a different point of view. I hope you live a life you’re proud of. If you find that you’re not, I hope you have the courage to start all over again.",
	},
	{
		Author: "P.C. CastBetrayed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470091601p2/17015.jpg",
		Content: "Remember, darkness does not always equate to evil, just as light does not always bring good.",
	},
	{
		Author: "David LevithanThe Lover's Dictionary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1426529210p2/11664.jpg",
		Content: "It was a mistake,\" you said. But the cruel thing was, it felt like the mistake was mine, for trusting you.",
	},
	{
		Author: "Dalai Lama XIV",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241989386p2/570218.jpg",
		Content: "If you think you are too small to make a difference, try sleeping with a mosquito.",
	},
	{
		Author: "William BlakeAuguries of Innocence",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199069675p2/13453.jpg",
		Content: "A truth that's told with bad intentBeats all the lies you can invent.",
	},
	{
		Author: "Henry MillerThe Books in My Life",
		AuthorAvatar: "https://images.gr-assets.com/authors/1511445828p2/147.jpg",
		Content: "A book lying idle on a shelf is wasted ammunition.",
	},
	{
		Author: "Terry Pratchett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Wisdom comes from experience. Experience is often a result of lack of wisdom.",
	},
	{
		Author: "Jude DeverauxA Knight in Shining Armor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363874185p2/28574.jpg",
		Content: "My soul will find yours.",
	},
	{
		Author: "T.H. WhiteThe Once and Future King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246071532p2/426944.jpg",
		Content: "The best thing for being sad,\" replied Merlin, beginning to puff and blow, \"is to learn something. That's the only thing that never fails. You may grow old and trembling in your anatomies, you may lie awake at night listening to the disorder of your veins, you may miss your only love, you may see the world about you devastated by evil lunatics, or know your honour trampled in the sewers of baser minds. There is only one thing for it then — to learn. Learn why the world wags and what wags it. That is the only thing which the mind can never exhaust, never alienate, never be tortured by, never fear or distrust, and never dream of regretting. Learning is the only thing for you. Look what a lot of things there are to learn.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "In time, the hurt began to fade and it was easier to just let it go. At least I thought it was. But in every boy I met in the next few years, I found myself looking for you, and when the feelings got too strong, I'd write you another letter. But I never sent them for fear of what I might find. By then, you'd gone on with your life and I didn't want to think about you loving someone else. I wanted to remember us like we were that summer. I didn't ever want to lose that.",
	},
	{
		Author: "Margaret Walker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1245010919p2/9660.jpg",
		Content: "When I was about eight, I decided that the most wonderful thing, next to a human being, was a book.",
	},
	{
		Author: "Orson Welles",
		AuthorAvatar: "https://images.gr-assets.com/authors/1352282642p2/67899.jpg",
		Content: "Ask not what you can do for your country. Ask what’s for lunch.",
	},
	{
		Author: "John Boyne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1439062965p2/7195.jpg",
		Content: "There's things that happen in a person's life that are so scorched in the memory and burned into the heart that there's no forgetting them.",
	},
	{
		Author: "Charles Dickens",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "Have a heart that never hardens, and a temper that never tires, and a touch that never hurts.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Courage is resistance to fear, mastery of fear - not absence of fear.",
	},
	{
		Author: "Horace Mann",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495999004p2/279932.jpg",
		Content: "A house without books is like a room without windows.",
	},
	{
		Author: "H.P. Lovecraft",
		AuthorAvatar: "https://images.gr-assets.com/authors/1299165714p2/9494.jpg",
		Content: "The world is indeed comic, but the joke is on mankind.",
	},
	{
		Author: "Susan Sontag",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285821018p2/7907.jpg",
		Content: "I haven't been everywhere, but it's on my list.",
	},
	{
		Author: "Dr. SeussHorton Hatches the Egg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "I meant what I said and I said what I meant. An elephant's faithful one-hundred percent!",
	},
	{
		Author: "Oscar WildeAn Ideal Husband",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "To love oneself is the beginning of a lifelong romance.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I'm going to wake Peeta,\" I say.\"No, wait,\" says Finnick. \"Let's do it together. Put our faces right in front of his.\"Well, there's so little opportunity for fun left in my life, I agree. We position ourselves on either side of Peeta, lean over until our faces are inches frim his nose, and give him a shake. \"Peeta. Peeta, wake up,\" I say in a soft, singsong voice.His eyelids flutter open and then he jumps like we've stabbed him. \"Aa!\"Finnick and I fall back in the sand, laughing our heads off. Every time we try to stop, we look at Peeta's attempt to maintain a disdainful expression and it sets us off again.",
	},
	{
		Author: "Steve MaraboliLife, the Truth, and Being Free",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "If people refuse to look at you in a new light and they can only see you for what you were, only see you for the mistakes you've made, if they don't realize that you are not your mistakes, then they have to go.",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "When I get lonely these days, I think: So BE lonely, Liz. Learn your way around loneliness. Make a map of it. Sit with it, for once in your life. Welcome to the human experience. But never again use another person's body or emotions as a scratching post for your own unfulfilled yearnings.",
	},
	{
		Author: "Tennessee WilliamsCat on a Hot Tin Roof",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206504901p2/7751.jpg",
		Content: "I've got the guts to die. What I want to know is, have you got the guts to live?",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "My soul is from elsewhere, I'm sure of that, and I intend to end up there.",
	},
	{
		Author: "Frank Zappa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1315160559p2/22302.jpg",
		Content: "If you end up with a boring miserable life because you listened to your mom, your dad, your teacher, your priest, or some guy on television telling you how to do your shit, then you deserve it.",
	},
	{
		Author: "Galway Kinnell",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1423685779p2/215571.jpg",
		Content: "Let our scars fall in love.",
	},
	{
		Author: "EuripidesThe Bacchae",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195014632p2/973.jpg",
		Content: "Talk sense to a fool and he calls you foolish.",
	},
	{
		Author: "C.S. LewisThe Silver Chair",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I'm on Aslan's side even if there isn't any Aslan to lead it. I'm going to live as like a Narnian as I can even if there isn't any Narnia.",
	},
	{
		Author: "Alan MooreV for Vendetta",
		AuthorAvatar: "https://images.gr-assets.com/authors/1304944713p2/3961.jpg",
		Content: "People shouldn't be afraid of their government. Governments should be afraid of their people.",
	},
	{
		Author: "George Bernard ShawMan and Superman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "The reasonable man adapts himself to the world: the unreasonable one persists in trying to adapt the world to himself. Therefore all progress depends on the unreasonable man.",
	},
	{
		Author: "Stephanie PerkinsAnna and the French Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407443106p2/3095893.jpg",
		Content: "Is it possible for home to be a person and not a place?",
	},
	{
		Author: "Bill Cosby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198544055p2/3537.jpg",
		Content: "A word to the wise ain't necessary, it's the stupid ones who need advice.",
	},
	{
		Author: "Kurt VonnegutThe Sirens of Titan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "A purpose of human life, no matter who is controlling it, is to love whoever is around to be loved.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Whenever you read a cancer booklet or website or whatever, they always list depression among the side effects of cancer. But, in fact, depression is not a side effect of cancer. Depression is a side effect of dying.",
	},
	{
		Author: "Oprah Winfrey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1354837955p2/3518.jpg",
		Content: "You can have it all. Just not all at once.",
	},
	{
		Author: "Cassandra ClareClockwork Princess",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You endure what is unbearable, and you bear it. That is all.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Truth only means something when it's hard to admit.",
	},
	{
		Author: "Milan KunderaThe Unbearable Lightness of Being",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465275207p2/6343.jpg",
		Content: "When the heart speaks, the mind finds it indecent to object.",
	},
	{
		Author: "William Saroyan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207407131p2/4095.jpg",
		Content: "When you laugh, laugh like hell. And when you get angry, get good and angry. Try to be alive. You will be dead soon enough.",
	},
	{
		Author: "Eugene Field",
		AuthorAvatar: "https://images.gr-assets.com/authors/1248357745p2/91616.jpg",
		Content: "No book can be appreciated until it has been slept with and dreamed over.",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "You can cut all the flowers but you cannot keep Spring from coming.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "I drink to make other people more interesting.",
	},
	{
		Author: "Markus Zusak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Sometimes you read a book so special that you want to carry it around with you for months after you've finished just to stay near it.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Even in half demon hunter clothes, Clary thought, he looked like the kind of boy who'd come over your house to pick you up for a date and be polite to your parents and nice to your pets.Jace on the other hand, looked like the kind of boy who'd come over your house and burn it down just for kicks.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "I want to see you.Know your voice.Recognize you when youfirst come 'round the corner.Sense your scent when I come into a room you've just left.Know the lift of your heel,the glide of your foot.Become familiar with the way you purse your lipsthen let them part, just the slightest bit,when I lean in to your spaceand kiss you.I want to know the joy of how you whisper \"more",
	},
	{
		Author: "Yogi BerraWhen You Come to a Fork in the Road, Take It!: Inspiration and Wisdom from One of Baseball's Greatest Heroes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1315253191p2/79014.jpg",
		Content: "Always go to other people's funerals, otherwise they won't come to yours.",
	},
	{
		Author: "Karen Marie MoningDarkfever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1497612477p2/48206.jpg",
		Content: "I love books, by the way, way more than movies. Movies tell you what to think. A good book lets you choose a few thoughts for yourself. Movies show you the pink house. A good book tells you there's a pink house and lets you paint some of the finishing touches, maybe choose the roof style,park your own car out front. My imagination has always topped anything a movie could come up with. Case in point, those darned Harry Potter movies. That was so not what that part-Veela-chick, Fleur Delacour, looked like.",
	},
	{
		Author: "Doris Lessing",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457477725p2/7728.jpg",
		Content: "Whatever you're meant to do, do it now. The conditions are always impossible.",
	},
	{
		Author: "Herbert Hoover",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209349646p2/209936.jpg",
		Content: "Older men declare war. But it is youth that must fight and die.",
	},
	{
		Author: "Astrid Lindgren",
		AuthorAvatar: "https://images.gr-assets.com/authors/1194544214p2/410653.jpg",
		Content: "A childhood without books – that would be no childhood. That would be like being shut out from the enchanted place where you can go and find the rarest kind of joy.",
	},
	{
		Author: "Joseph HellerCatch-22",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197308614p2/3167.jpg",
		Content: "Just because you're paranoid doesn't mean they aren't after you.",
	},
	{
		Author: "Rick RiordanThe Sea of Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Families are messy. Immortal families are eternally messy. Sometimes the best we can do is to remind each other that we're related for better or for worse...and try to keep the maiming and killing to a minimum.",
	},
	{
		Author: "Eoin ColferArtemis Fowl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1254336426p2/10896.jpg",
		Content: "Confidence is ignorance. If you're feeling cocky, it's because there's something you don't know.",
	},
	{
		Author: "أحمد خالد توفيق",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316682283p2/1479015.jpg",
		Content: "إن المرأة تحب رجلها ليس لأنه أقوى الرجال، و لا أوسمهم، و لا أغناهم، بل لأنه هو.. بضعفه و قوته.. و الحب ليس إستعراض قوة لكنه طاقة عطاء دافئة مستمرة",
	},
	{
		Author: "Jay AsherThirteen Reasons Why",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243931536p2/569269.jpg",
		Content: "A lot of you cared, just not enough.",
	},
	{
		Author: "Tahereh MafiShatter Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "I'm oxygen and he's dying to breathe.",
	},
	{
		Author: "Lauren KateTorment",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502298656p2/2905297.jpg",
		Content: "Sometimes beautiful things come into our lives out of nowhere. We can't always understand them, but we have to trust in them. I know you want to question everything, but sometimes it pays to just have a little faith.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Adults are just obsolete children and the hell with them.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "He was gone, and I did not have time to tell him what I had just now realized: that I forgave him, and that she forgave us, and that we had to forgive to survive in the labyrinth. There were so many of us who would have to live with things done and things left undone that day. Things that did not go right, things that seemed okay at the time because we could not see the future. If only we could see the endless string of consequences that result from our smallest actions. But we can’t know better until knowing better is useless. And as I walked back to give Takumi’s note to the Colonel, I saw that I would never know. I would never know her well enough to know her thoughts in those last minutes, would never know if she left us on purpose. But the not-knowing would not keep me from caring, and I would always love Alaska Young, my crooked neighbor, with all my crooked heart.",
	},
	{
		Author: "Napoléon Bonaparte",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397332087p2/210910.jpg",
		Content: "Never interrupt your enemy when he is making a mistake.",
	},
	{
		Author: "Edna St. Vincent Millay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183231710p2/33998.jpg",
		Content: "I know I am but summer to your heart, and not the full four seasons of the year.",
	},
	{
		Author: "Max LucadoHe Still Moves Stones: Everyone Needs a Miracle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493056226p2/2737.jpg",
		Content: "Faith is not the belief that God will do what you want. It is the belief that God will do what is right.",
	},
	{
		Author: "Fernando PessoaThe Book of Disquiet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1427039317p2/7816.jpg",
		Content: "Literature is the most agreeable way of ignoring life.",
	},
	{
		Author: "Nicholas SparksThe Rescue",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "People come, people go – they’ll drift in and out of your life, almost like characters in a favorite book. When you finally close the cover, the characters have told their story and you start up again with another book, complete with new characters and adventures. Then you find yourself focusing on the new ones, not the ones from the past.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "And I guess I realized at that moment that I really did love her. Because there was nothing to gain, and that didn't matter.",
	},
	{
		Author: "Dorothy Parker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820565p2/24956.jpg",
		Content: "If you want to know what God thinks of money, just look at the people he gave it to.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It is a curious thing, Harry, but perhaps those who are best suited to power are those who have never sought it. Those who, like you, have leadership thrust upon them, and take up the mantle because they must, and find to their own surprise that they wear it well.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Tomorrow may be hell, but today was a good writing day, and on the good writing days nothing else matters.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "The pleasure of remembering had been taken from me, because there was no longer anyone to remember with. It felt like losing your co-rememberer meant losing the memory itself, as if the things we'd done were less real and important than they had been hours before.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Isn't it enough to see that a garden is beautiful without having to believe that there are fairies at the bottom of it too?",
	},
	{
		Author: "Hunter S. ThompsonThe Proud Highway: Saga of a Desperate Southern Gentleman, 1955-1967",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206560814p2/5237.jpg",
		Content: "A man who procrastinates in his choosing will inevitably have his choice made for him by circumstance.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "The purpose of life is to live it, to taste experience to the utmost, to reach out eagerly and without fear for newer and richer experience.",
	},
	{
		Author: "Ray BradburyFahrenheit 451",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "Stuff your eyes with wonder, he said, live as if you'd drop dead in ten seconds. See the world. It's more fantastic than any dream made or paid for in factories.",
	},
	{
		Author: "Laini TaylorDaughter of Smoke \u0026 Bone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224474224p2/324620.jpg",
		Content: "Once upon a time, an angel and a devil fell in love. It did not end well.",
	},
	{
		Author: "Milan KunderaThe Unbearable Lightness of Being",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465275207p2/6343.jpg",
		Content: "Anyone whose goal is 'something higher' must expect someday to suffer vertigo. What is vertigo? Fear of falling? No, Vertigo is something other than fear of falling. It is the voice of the emptiness below us which tempts and lures us, it is the desire to fall, against which, terrified, we defend ourselves.",
	},
	{
		Author: "Arthur O'ShaughnessyPoems of Arthur O'Shaughnessy",
		AuthorAvatar: "",
		Content: "We are the music makers, and we are the dreamers of dreams.",
	},
	{
		Author: "Terry PratchettLords and Ladies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "In the beginning there was nothing, which exploded.",
	},
	{
		Author: "Kurt VonnegutSlaughterhouse-Five",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "And I asked myself about the present: how wide it was, how deep it was, how much was mine to keep.",
	},
	{
		Author: "Mahmoud Darwish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414535868p2/75055.jpg",
		Content: "و كن من أنتَ حيث تكون و احمل عبءَ قلبِكَ وحدهُ",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "The scar had not pained Harry for nineteen years. All was well.",
	},
	{
		Author: "Nick Hornby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422915487p2/2929.jpg",
		Content: "It's no good pretending that any relationship has a future if your record collections disagree violently or if your favorite films wouldn't even speak to each other if they met at a party.",
	},
	{
		Author: "Walt Whitman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392303683p2/1438.jpg",
		Content: "Do I contradict myself? Very well, then, I contradict myself; I am large -- I contain multitudes.",
	},
	{
		Author: "Richelle MeadVampire Academy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "The only thing better than imagining Dimitri carrying me in his arms was imagining him shirtless while carrying me in his arms.",
	},
	{
		Author: "Lillian Hellman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1257464297p2/66241.jpg",
		Content: "People change and forget to tell each other.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "The glory of friendship is not the outstretched hand, not the kindly smile, nor the joy of companionship; it is the spiritual inspiration that comes to one when you discover that someone else believes in you and is willing to trust you with a friendship.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "If A is a success in life, then A equals x plus y plus z. Work is x; y is play; and z is keeping your mouth shut",
	},
	{
		Author: "Charles LambThe Life, Letters and Writings of Charles Lamb",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198641106p2/6237.jpg",
		Content: "Tis the privilege of friendship to talk nonsense, and to have her nonsense respected.",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "Animals are my friends...and I don't eat my friends.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "Nearly all men can stand adversity, but if you want to test a man's character, give him power.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Who, being loved, is poor?",
	},
	{
		Author: "Zig Ziglar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1236462809p2/50316.jpg",
		Content: "Lack of direction, not lack of time, is the problem. We all have twenty-four hour days.",
	},
	{
		Author: "Fannie FlaggFried Green Tomatoes at the Whistle Stop Cafe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1328803839p2/6125.jpg",
		Content: "Remember if people talk behind your back, it only means you are two steps ahead.",
	},
	{
		Author: "J.D. SalingerFranny and Zooey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "I'm sick of just liking people. I wish to God I could meet somebody I could respect.",
	},
	{
		Author: "Robert A. HeinleinStarship Troopers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192826560p2/205.jpg",
		Content: "Happiness consists in getting enough sleep. Just that, nothing more.",
	},
	{
		Author: "Benjamin Franklin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215314094p2/289513.jpg",
		Content: "Tell me and I forget, teach me and I may remember, involve me and I learn.",
	},
	{
		Author: "Thomas WolfeYou Can't Go Home Again",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188854076p2/7921.jpg",
		Content: "Make your mistakes, take your chances, look silly, but keep on going. Don’t freeze up.",
	},
	{
		Author: "Shel Silverstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "I cannot go to school today\"Said little Peggy Ann McKay.\"I have the measles and the mumps,A gash, a rash and purple bumps.My mouth is wet, my throat is dry.I'm going blind in my right eye.My tonsils are as big as rocks,I've counted sixteen chicken pox.And there's one more - that's seventeen,And don't you think my face looks green?My leg is cut, my eyes are blue,It might be the instamatic flu.I cough and sneeze and gasp and choke,I'm sure that my left leg is broke.My hip hurts when I move my chin,My belly button's caving in.My back is wrenched, my ankle's sprained,My 'pendix pains each time it rains.My toes are cold, my toes are numb,I have a sliver in my thumb.My neck is stiff, my voice is weak,I hardly whisper when I speak.My tongue is filling up my mouth,I think my hair is falling out.My elbow's bent, my spine ain't straight,My temperature is one-o-eight.My brain is shrunk, I cannot hear,There's a hole inside my ear.I have a hangnail, and my heart is ...What? What's that? What's that you say?You say today is .............. Saturday?G'bye, I'm going out to play!",
	},
	{
		Author: "Paulo CoelhoThe Devil and Miss Prym",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "When we least expect it, life sets us a challenge to test our courage and willingness to change; at such a moment, there is no point in pretending that nothing has happened or in saying that we are not yet ready. The challenge will not wait. Life does not look back. A week is more than enough time for us to decide whether or not to accept our destiny.",
	},
	{
		Author: "Terry PratchettSourcery",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "I meant,\" said Ipslore bitterly, \"what is there in this world that truly makes living worthwhile?\"Death thought about it.CATS, he said eventually. CATS ARE NICE.",
	},
	{
		Author: "Sarah DessenThe Truth About Forever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "I like flaws. I think they make things interesting.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Go on with what your heart tells you, or you will lose all.",
	},
	{
		Author: "Brian JacquesRedwall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1309635011p2/5329.jpg",
		Content: "Even the strongest and bravest must sometimes weep. It shows they have a great heart, one that can feel compassion for others.",
	},
	{
		Author: "Jon KrakauerInto the Wild",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430856379p2/1235.jpg",
		Content: "It's not always necessary to be strong, but to feel strong.",
	},
	{
		Author: "Bill Watterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "You know, Hobbes, some days even my lucky rocket ship underpants don't help.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I have learned now that while those who speak about one's miseries usually hurt, those who keep silence hurt more.",
	},
	{
		Author: "Emily BrontëWuthering Heights",
		AuthorAvatar: "https://images.gr-assets.com/authors/1473229007p2/4191.jpg",
		Content: "Be with me always - take any form - drive me mad! only do not leave me in this abyss, where I cannot find you! Oh, God! it is unutterable! I can not live without my life! I can not live without my soul!",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Jace shook his blond head in exasperation. \"You had to make a crazy jail friend, didn't you? You couldn't just count ceiling tiles or tame a pet mouse like normal prisoners do?",
	},
	{
		Author: "Brian AndreasStory People: Selected Stories \u0026 Drawings of Brian Andreas",
		AuthorAvatar: "https://images.gr-assets.com/authors/1242331314p2/32114.jpg",
		Content: "I read once that the ancient Egyptians had fifty words for sand \u0026 the Eskimos had a hundred words for snow. I wish I had a thousand words for love, but all that comes to mind is the way you move against me while you sleep \u0026 there are no words for that.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Have you really read all those books in your room?",
	},
	{
		Author: "Ben Okri",
		AuthorAvatar: "https://images.gr-assets.com/authors/1294662937p2/31425.jpg",
		Content: "Stories can conquer fear, you know. They can make the heart bigger.",
	},
	{
		Author: "Haruki MurakamiKafka on the Shore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "It's like Tolstoy said. Happiness is an allegory, unhappiness a story.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Kindness is a language which the deaf can hear and the blind can see.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I shut my eyes and all the world drops dead; I lift my eyes and all is born again.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I must have loved you a lot.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "He stepped down, trying not to look long at her, as if she were the sun, yet he saw her, like the sun, even without looking.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "If we couldn't laugh we would all go insane.",
	},
	{
		Author: "Kiera CassThe One",
		AuthorAvatar: "https://images.gr-assets.com/authors/1318605410p2/2987125.jpg",
		Content: "Break my heart. Break it a thousand times if you like. It was only ever yours to break anyway.",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "Sonnet XVIII do not love you as if you were salt-rose, or topaz,or the arrow of carnations the fire shoots off.I love you as certain dark things are to be loved,in secret, between the shadow and the soul.I love you as the plant that never bloomsbut carries in itself the light of hidden flowers;thanks to your love a certain solid fragrance,risen from the earth, lives darkly in my body.I love you without knowing how, or when, or from where.I love you straightforwardly, without complexities or pride;so I love you because I know no other way than this: where I does not exist, nor you,so close that your hand on my chest is my hand,so close that your eyes close as I fall asleep. ",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Peeta, how come I never know when you're having a nightmare?",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "The ships hung in the sky in much the same way that bricks don't.",
	},
	{
		Author: "Stanisław Lem",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246185166p2/10991.jpg",
		Content: "Good books tell the truth, even when they're about things that never have been and never will be. They're truthful in a different way.",
	},
	{
		Author: "Adrienne Rich",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198772292p2/29947.jpg",
		Content: "There must be those among whom we can sit down and weep and still be counted as warriors.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Crying is for plain women. Pretty women go shopping.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "What do you want?\"\"Just coffee. Black - like my soul.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "It must require bravery to be honest all the time.",
	},
	{
		Author: "Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Remember, remember, this is now, and now, and now. Live it, feel it, cling to it. I want to become acutely aware of all I’ve taken for granted.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "He accused me of being Dumbledore's man through and through.\"\"How very rude of him.\"\"I told him I was.\"Dumbledore opened his mouth to speak and then closed it again. Fawkes the phoenix let out a low, soft, musical cry. To Harry's intense embarrassment, he suddenly realized that Dumbledore's bright blue eyes looked rather watery, and stared hastily at his own knee. When Dumbledore spoke, however, his voice was quite steady. \"I am very touched, Harry.",
	},
	{
		Author: "Mitch Albom",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "Lost love is still love. It takes a different form, that's all. You can't see their smile or bring them food or tousle their hair or move them around a dance floor. But when those senses weaken another heightens. Memory. Memory becomes your partner. You nurture it. You hold it. You dance with it.",
	},
	{
		Author: "Sylvia PlathThe Bell Jar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "If neurotic is wanting two mutually exclusive things at one and the same time, then I'm neurotic as hell. I'll be flying back and forth between one mutually exclusive thing and another for the rest of my days.",
	},
	{
		Author: "Barack Obama",
		AuthorAvatar: "https://images.gr-assets.com/authors/1313512321p2/6356.jpg",
		Content: "Change will not come if we wait for some other person, or if we wait for some other time. We are the ones we've been waiting for. We are the change that we seek.",
	},
	{
		Author: "Jean de La Fontaine",
		AuthorAvatar: "https://images.gr-assets.com/authors/1247678674p2/82948.jpg",
		Content: "Rare as is true love, true friendship is rarer.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Never make someone a priority when all you are to them is an option.",
	},
	{
		Author: "Peter ShafferFive Finger Exercise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206548676p2/28785.jpg",
		Content: "The trouble is if you don’t spend your life yourself, other people spend it for you.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "You can only become truly accomplished at something you love. Don’t make money your goal. Instead pursue the things you love doing and then do them so well that people can’t take their eyes off of you.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Yes, I was infatuated with you: I am still. No one has ever heightened such a keen capacity of physical sensation in me. I cut you out because I couldn't stand being a passing fancy. Before I give my body, I must give my thoughts, my mind, my dreams. And you weren't having any of those.",
	},
	{
		Author: "Arrigo Boito",
		AuthorAvatar: "https://images.gr-assets.com/authors/1439755156p2/293421.jpg",
		Content: "When I saw you I fell in love, and you smiled because you knew.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "You are not your job, you're not how much money you have in the bank. You are not the car you drive. You're not the contents of your wallet. You are not your fucking khakis. You are all singing, all dancing crap of the world.",
	},
	{
		Author: "Phyllis Diller",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209492226p2/271137.jpg",
		Content: "A smile is a curve that sets everything straight.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Ginny!\" said Mr. Weasley, flabbergasted. \"Haven't I taught you anything? What have I always told you? Never trust anything that can think for itself if you can't see where it keeps its brain?",
	},
	{
		Author: "Pablo Picasso",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536109p2/3253.jpg",
		Content: "Art washes away from the soul the dust of everyday life.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "Respect was invented to cover the empty place where love should be.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "The greatness of a man is not in how much wealth he acquires, but in his integrity and his ability to affect those around him positively",
	},
	{
		Author: "William ShakespeareMacbeth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "By the pricking of my thumbs, Something wicked this way comes.",
	},
	{
		Author: "Charles BukowskiTales of Ordinary Madness",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "I felt like crying but nothing came out. it was just a sort of sad sickness, sick sad, when you can't feel any worse. I think you know it. I think everybody knows it now and then. but I think I have known it pretty often, too often.",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "My rapier wit hides my inner pain.",
	},
	{
		Author: "Boris Pasternak",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422699346p2/7902.jpg",
		Content: "I don't like people who have never fallen or stumbled. Their virtue is lifeless and it isn't of much value. Life hasn't revealed its beauty to them. ",
	},
	{
		Author: "Margaret MitchellGone with the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1185481392p2/11081.jpg",
		Content: "Burdens are for shoulders strong enough to carry them.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "You must not lose faith in humanity. Humanity is like an ocean; if a few drops of the ocean are dirty, the ocean does not become dirty.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "There is only one thing in the world worse than being talked about, and that is not being talked about.",
	},
	{
		Author: "Susan SontagAs Consciousness is Harnessed to Flesh: Journals and Notebooks, 1964-1980",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285821018p2/7907.jpg",
		Content: "My library is an archive of longings.",
	},
	{
		Author: "Evelyn WaughBrideshead Revisited: The Sacred and Profane Memories of Captain Charles Ryder",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357463949p2/11315.jpg",
		Content: "Sometimes, I feel the past and the future pressing so hard on either side that there's no room for the present at all.",
	},
	{
		Author: "A.A. Milne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "What day is it?\"It's today,\" squeaked Piglet.My favorite day,\" said Pooh.",
	},
	{
		Author: "George Eliot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525078524p2/173.jpg",
		Content: "What do we live for, if it is not to make life less difficult for each other?",
	},
	{
		Author: "Marissa MeyerCinder",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1455670017p2/488008.jpg",
		Content: "Even in the Future the Story Begins with Once Upon a Time.",
	},
	{
		Author: "Neil LaButeReasons to Be Pretty",
		AuthorAvatar: "https://images.gr-assets.com/authors/1220211917p2/17542.jpg",
		Content: "The future is now. It's time to grow up and be strong. Tomorrow may well be too late.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I fell in love with her when we were together, then fell deeper in love with her in the years we were apart.",
	},
	{
		Author: "Lewis CarrollAlice's Adventures in Wonderland \u0026 Through the Looking-Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "I wonder if the snow loves the trees and fields, that it kisses them so gently? And then it covers them up snug, you know, with a white quilt; and perhaps it says, \"Go to sleep, darlings, till the summer comes again.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I have the simplest tastes. I am always satisfied with the best.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "You're a painter. You're a baker. You like to sleep with the windows open. You never take sugar in your tea. And you always double-knot your shoelaces.",
	},
	{
		Author: "Anne Sexton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525637942p2/26814.jpg",
		Content: "As it has been said:Love and a coughcannot be concealed.Even a small cough.Even a small love.",
	},
	{
		Author: "Gustave Flaubert",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198541369p2/1461.jpg",
		Content: "Do not read as children do to enjoy themselves, or, as the ambitious do to educate themselves. No, read to live.",
	},
	{
		Author: "Gail Carson Levine",
		AuthorAvatar: "https://images.gr-assets.com/authors/1432495575p2/13677.jpg",
		Content: "A library is infinity under a roof.",
	},
	{
		Author: "W.C. Fields",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331712829p2/82951.jpg",
		Content: "It ain't what they call you, it's what you answer to.",
	},
	{
		Author: "Hermann HesseDemian. Die Geschichte von Emil Sinclairs Jugend",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499981916p2/1113469.jpg",
		Content: "If you hate a person, you hate something in him that is part of yourself. What isn't part of ourselves doesn't disturb us.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The reason it hurts so much to separate is because our souls are connected. Maybe they always have been and will be. Maybe we've lived a thousand lives before this one and in each of them we've found each other. And maybe each time, we've been forced apart for the same reasons. That means that this goodbye is both a goodbye for the past ten thousand years and a prelude to what will come.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Love can change a person the way a parent can change a baby- awkwardly, and often with a great deal of mess.",
	},
	{
		Author: "William ShakespeareRomeo and Juliet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "My bounty is as boundless as the sea,My love as deep; the more I give to thee,The more I have, for both are infinite.",
	},
	{
		Author: "Veronica RothAllegiant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "There are so many ways to be brave in this world. Sometimes bravery involves laying down your life for something bigger than yourself, or for someone else. Sometimes it involves giving up everything you have ever known, or everyone you have ever loved, for the sake of something greater.But sometimes it doesn't.Sometimes it is nothing more than gritting your teeth through pain, and the work of every day, the slow walk toward a better life. That is the sort of bravery I must have now.",
	},
	{
		Author: "Richelle MeadVampire Academy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Did you see that dress?",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Nothing in the world is more dangerous than sincere ignorance and conscientious stupidity.",
	},
	{
		Author: "John UpdikeMy Father's Tears and Other Stories",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419249254p2/6878.jpg",
		Content: "It is easy to love people in memory; the hard thing is to love them when they are there in front of you.",
	},
	{
		Author: "مصطفى محمودفي الحب والحياة",
		AuthorAvatar: "https://images.gr-assets.com/authors/1458565195p2/871358.jpg",
		Content: "الكراهية تكلف أكثر من الحب.. لأنها إحساس غير طبيعي.. إحساس عكسي مثل حركة الأجسام ضد جاذبية الأرض.. تحتاج إلى قوة إضافية وتستهلك وقوداً أكثر",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "The only way to deal with an unfree world is to become so absolutely free that your very existence is an act of rebellion.",
	},
	{
		Author: "Jack Kerouac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "I like too many things and get all confused and hung-up running from one falling star to another till i drop. This is the night, what it does to you. I had nothing to offer anybody except my own confusion.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "I will remember the kisses our lips raw with love and how you gave me everything you had and how I offered you what was left of me, and I will remember your small room the feel of you the light in the window your records your books our morning coffee our noons our nights our bodies spilled together sleeping the tiny flowing currents immediate and forever your leg my leg your arm my arm your smile and the warmth of you who made me laugh again.",
	},
	{
		Author: "John GreenWill Grayson, Will Grayson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "You like someone who can't like you back because unrequited love can be survived in a way that once-requited love cannot. ",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "You know how they say you only hurt the ones you love? Well, it works both ways.",
	},
	{
		Author: "Nikolai Gogol",
		AuthorAvatar: "https://images.gr-assets.com/authors/1303965602p2/232932.jpg",
		Content: "The longer and more carefully we look at a funny story, the sadder it becomes.",
	},
	{
		Author: "Daniel DefoeRobinson Crusoe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202587497p2/2007.jpg",
		Content: "It is never too late to be wise.",
	},
	{
		Author: "Octavia E. ButlerParable of the Talents",
		AuthorAvatar: "https://images.gr-assets.com/authors/1242244143p2/29535.jpg",
		Content: "In order to riseFrom its own ashesA phoenixFirstMustBurn.",
	},
	{
		Author: "Julian BarnesThe Sense of an Ending",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387175450p2/1462.jpg",
		Content: "This was another of our fears: that Life wouldn't turn out to be like Literature.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "Believe nothing you hear, and only one half that you see.",
	},
	{
		Author: "Maya AngelouPhenomenal Woman: Four Poems Celebrating Women",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Pretty women wonder where my secret lies.I'm not cute or built to suit a fashion model's sizeBut when I start to tell them,They think I'm telling lies.I say,It's in the reach of my armsThe span of my hips,The stride of my step,The curl of my lips.I'm a womanPhenomenally.Phenomenal woman,That's me.",
	},
	{
		Author: "Aung San Suu Kyi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1238124060p2/61546.jpg",
		Content: "The only real prison is fear, and the only real freedom is freedom from fear",
	},
	{
		Author: "Mitch AlbomThe Five People You Meet in Heaven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "All parents damage their children. It cannot be helped. Youth, like pristine glass, absorbs the prints of its handlers. Some parents smudge, others crack, a few shatter childhoods completely into jagged little pieces, beyond repair.",
	},
	{
		Author: "Michelle HodkinThe Unbecoming of Mara Dyer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454082708p2/4126827.jpg",
		Content: "Thinking something does not make it true. Wanting something does not make it real.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Make your own Bible. Select and collect all the words and sentences that in all your readings have been to you like the blast of a trumpet.",
	},
	{
		Author: "Vladimir NabokovLolita",
		AuthorAvatar: "https://images.gr-assets.com/authors/1482502806p2/5152.jpg",
		Content: "Lolita, light of my life, fire of my loins. My sin, my soul. Lo-lee-ta: the tip of the tongue taking a trip of three steps down the palate to tap, at three, on the teeth. Lo. Lee. Ta. She was Lo, plain Lo, in the morning, standing four feet ten in one sock. She was Lola in slacks. She was Dolly at school. She was Dolores on the dotted line. But in my arms she was always Lolita. Did she have a precursor? She did, indeed she did. In point of fact, there might have been no Lolita at all had I not loved, one summer, an initial girl-child. In a princedom by the sea. Oh when? About as many years before Lolita was born as my age was that summer. You can always count on a murderer for a fancy prose style. Ladies and gentlemen of the jury, exhibit number one is what the seraphs, the misinformed, simple, noble-winged seraphs, envied. Look at this tangle of thorns.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "Don't spend time beating on a wall, hoping to transform it into a door. ",
	},
	{
		Author: "Sarah Dessen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Anyone can hide. Facing up to things, working through them, that's what makes you strong.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "When you play a game of thrones you win or you die.",
	},
	{
		Author: "Rainer Maria RilkeTranslations from the Poetry of Rainer Maria Rilke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493785350p2/7906.jpg",
		Content: "We need, in love, to practice only this: letting each other go. For holding on comes easily; we do not need to learn it.",
	},
	{
		Author: "Julia Child",
		AuthorAvatar: "https://images.gr-assets.com/authors/1345004268p2/3465.jpg",
		Content: "A party without cake is just a meeting",
	},
	{
		Author: "Elie Wiesel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1255518412p2/1049.jpg",
		Content: "Friendship marks a life even more deeply than love. Love risks degenerating into obsession, friendship is never anything but sharing.",
	},
	{
		Author: "Robert Louis StevensonTravels with a Donkey in the Cevennes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192746024p2/854076.jpg",
		Content: "I travel not to go anywhere, but to go. I travel for travel's sake. The great affair is to move.",
	},
	{
		Author: "Robert Louis Stevenson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192746024p2/854076.jpg",
		Content: "Life is not a matter of holding good cards, but of playing a poor hand well.",
	},
	{
		Author: "Stieg LarssonThe Girl with the Dragon Tattoo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246466225p2/706255.jpg",
		Content: "What she had realized was that love was that moment when your heart was about to burst.",
	},
	{
		Author: "Steve Jobs",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322675535p2/5255891.jpg",
		Content: "Remembering that you are going to die is the best way I know to avoid the trap of thinking you have something to lose. You are already naked. There is no reason not to follow your heart.",
	},
	{
		Author: "John Muir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398092241p2/5297.jpg",
		Content: "When one tugs at a single thing in nature, he finds it attached to the rest of the world.",
	},
	{
		Author: "Jamie McGuireBeautiful Disaster",
		AuthorAvatar: "https://images.gr-assets.com/authors/1479315727p2/4464118.jpg",
		Content: "I know we're fucked up, alright? I'm impulsive, and hot tempered, and you get under my skin like no one else. You act like you hate me one minute, and then need me the next. I never get anything right, and I don't deserve you...but I fucking love you, Abby. I love you more than I loved anyone or anything ever. When you're around, I don't need booze, or money, or the fighting, or the one-night stands...",
	},
	{
		Author: "Anthony Powell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1331081008p2/9947.jpg",
		Content: "I get a warm feeling among my books.",
	},
	{
		Author: "Socrates",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390145726p2/275648.jpg",
		Content: "There is only one good, knowledge, and one evil, ignorance.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "Only the dead have seen the end of war.",
	},
	{
		Author: "John SteinbeckEast of Eden",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1473708026p2/184186.jpg",
		Content: "And now that you don't have to be perfect, you can be good.",
	},
	{
		Author: "Glen CookSweet Silver Blues",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207159752p2/13026.jpg",
		Content: "Morning is wonderful. Its only drawback is that it comes at such an inconvenient time of day.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "No one can construct for you the bridge upon which precisely you must cross the stream of life, no one but you yourself alone.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "If I’m honest I have to tell you I still read fairy-tales and I like them best of all.",
	},
	{
		Author: "Sophocles",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195014481p2/1002.jpg",
		Content: "One wordFrees us of all the weight and pain of life:That word is love.",
	},
	{
		Author: "William GoldmanThe Princess Bride",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198704782p2/12521.jpg",
		Content: "When I was your age, television was called books.",
	},
	{
		Author: "Margaret Atwood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "A word after a word after a word is power.",
	},
	{
		Author: "Charles Lamb",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198641106p2/6237.jpg",
		Content: "I always arrive late at the office, but I make up for it by leaving early.",
	},
	{
		Author: "Karen Blixen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1405181948p2/8147.jpg",
		Content: "I know of a cure for everything: salt water...in one way or the other. Sweat, or tears, or the salt sea.",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "I'd rather die my way than live yours.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Little by little, one travels far",
	},
	{
		Author: "Tennessee WilliamsThe Glass Menagerie",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206504901p2/7751.jpg",
		Content: "Time is the longest distance between two places.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Headline?\" he asked.\"'Swing Set Needs Home,'\" I said.\"'Desperately Lonely Swing Set Needs Loving Home,'\" he said.\"'Lonely, Vaguely Pedophilic Swing Set Seeks the Butts of Children,'\" I said.",
	},
	{
		Author: "Stephenie MeyerNew Moon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "Forbidden to remember, terrified to forget; it was a hard line to walk.",
	},
	{
		Author: "Thomas PaineA Letter Addressed to the Abbe Raynal on the Affairs of North America",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390497645p2/57639.jpg",
		Content: "The mind once enlightened cannot again become dark.",
	},
	{
		Author: "P.C. Cast",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470091601p2/17015.jpg",
		Content: "If you have good friends, no matter how much life is sucking , they can make you laugh.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Black holes are where God divided by zero.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "And all the books you've read have been read by other people. And all the songs you've loved have been heard by other people. And that girl that's pretty to you is pretty to other people. and that if you looked at these facts when you were happy, you would feel great because you are describing 'unity.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Ah, music,\" he said, wiping his eyes. \"A magic beyond all we do here!",
	},
	{
		Author: "Confucius",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407613261p2/15321.jpg",
		Content: "It does not matter how slowly you go as long as you do not stop.",
	},
	{
		Author: "Carol ShieldsThe Republic of Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1240080700p2/12034.jpg",
		Content: "Open a book this minute and start reading. Don’t move until you’ve reached page fifty. Until you’ve buried your thoughts in print. Cover yourself with words. Wash yourself away. Dissolve.",
	},
	{
		Author: "Michael Crichton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359042651p2/5194.jpg",
		Content: "If you don't know history, then you don't know anything. You are a leaf that doesn't know it is part of a tree. ",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Being a writer is a very peculiar sort of a job: it's always you versus a blank sheet of paper (or a blank screen) and quite often the blank piece of paper wins.",
	},
	{
		Author: "V.S. NaipaulIn a Free State",
		AuthorAvatar: "https://images.gr-assets.com/authors/1497418880p2/3989.jpg",
		Content: "The only lies for which we are truly punished are those we tell ourselves.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "Boring damned people. All over the earth. Propagating more boring damned people. What a horror show. The earth swarmed with them.",
	},
	{
		Author: "Joyce Carol Oates",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454307466p2/3524.jpg",
		Content: "Reading is the sole means by which we slip, involuntarily, often helplessly, into another's skin, another's voice, another's soul.",
	},
	{
		Author: "Yann MartelLife of Pi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454012903p2/811.jpg",
		Content: "To choose doubt as a philosophy of life is akin to choosing immobility as a means of transportation.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Patience, grasshopper,\" said Maia. \"Good things come to those who wait.\"\"I always thought that was 'Good things come to those who do the wave,'\" said Simon. \"No wonder I've been so confused all my life.",
	},
	{
		Author: "Anne Frankdiary of Anne Frank: the play",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "No one has ever become poor by giving.",
	},
	{
		Author: "Steve Martin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1410018030p2/7103.jpg",
		Content: "Some people have a way with words, and other people...oh, uh, not have way.",
	},
	{
		Author: "T.S. Eliot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1507887310p2/18540.jpg",
		Content: "Only those who will risk going too far can possibly find out how far one can go.",
	},
	{
		Author: "Emilie AutumnThe Asylum for Wayward Victorian Girls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510961475p2/2528119.jpg",
		Content: "It gives me strength to have somebody to fight for; I can never fight for myself, but, for others, I can kill.",
	},
	{
		Author: "Rainer Maria Rilke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493785350p2/7906.jpg",
		Content: "Let everything happen to youBeauty and terrorJust keep goingNo feeling is final",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Remember me and smile, for it's better to forget than to remember me and cry.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "The best way to find yourself is to lose yourself in the service of others.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "When we don't know who to hate, we hate ourselves.",
	},
	{
		Author: "Jasper FfordeSomething Rotten",
		AuthorAvatar: "https://images.gr-assets.com/authors/1350497674p2/4432.jpg",
		Content: "If the real world were a book, it would never find a publisher. Overlong, detailed to the point of distraction-and ultimately, without a major resolution.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I have faith that God will show you the answer. But you have to understand that sometimes it takes a while to be able to recognize what God wants you to do. That's how it often is. God's voice is usually nothing more than a whisper, and you have to listen very carefully to hear it. But other times, in those rarest of moments, the answer is obvious and rings as loud as a church bell.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "The snake which cannot cast its skin has to die. As well the minds which are prevented from changing their opinions; they cease to be mind.",
	},
	{
		Author: "Henry Wadsworth Longfellow",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211861766p2/2697.jpg",
		Content: "The love of learning, the sequestered nooks,And all the sweet serenity of books",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "Never underestimate the power of stupid people in large groups.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "Learn from the mistakes of others. You can never live long enough to make them all yourself.",
	},
	{
		Author: "Sarah DessenLock and Key",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "What is family? They were the people who claimed you. In good, in bad, in parts or in whole, they were the ones who showed up, who stayed in there, regardless. It wasn't just about blood relations or shared chromosomes, but something wider, bigger. We had many families over time. Our family of origin, the family we created, and the groups you moved through while all of this was happening: friends, lovers, sometimes even strangers. None of them perfect, and we couldn't expect them to be. You can't make any one person your world. The trick was to take what each could give you and build your world from it.",
	},
	{
		Author: "Nicholas SparksThe Last Song",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "you have to love something before you can hate it.",
	},
	{
		Author: "أحلام مستغانميعابر سرير",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351458215p2/1109606.jpg",
		Content: "أي علم هذا الذي لم يستطع حتى الآن أن يضع أصوات من نحب في أقراص ، أو زجاجة دواء نتناولها سرًّا ، عندما نصاب بوعكة عاطفية بدون أن يدري صاحبها كم نحن نحتاجه",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "The last enemy that shall be destroyed is death.",
	},
	{
		Author: "Alex Haley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1182806085p2/17434.jpg",
		Content: "Either you deal with what is the reality, or you can be sure that the reality is going to deal with you.",
	},
	{
		Author: "Laini TaylorDaughter of Smoke \u0026 Bone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224474224p2/324620.jpg",
		Content: "Hope can be a powerful force. Maybe there's no actual magic in it, but when you know what you hope for most and hold it like a light within you, you can make things happen, almost like magic.",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Tears are words that need to be written.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "I think that we're all mentally ill. Those of us outside the asylums only hide it a little better - and maybe not all that much better after all.",
	},
	{
		Author: "Roy T. BennettThe Light in the Heart",
		AuthorAvatar: "",
		Content: "Attitude is a choice. Happiness is a choice. Optimism is a choice. Kindness is a choice. Giving is a choice. Respect is a choice. Whatever choice you make makes you. Choose wisely.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "An Unbreakable Vow?\" said Ron, looking stunned. \"Nah, he can’t have.... Are you sure?\"\"Yes I’m sure,\" said Harry. \"Why, what does it mean?\"\"Well, you can’t break an Unbreakable Vow...\"\"I’d worked that much out for myself, funnily enough.",
	},
	{
		Author: "Laurell K. HamiltonA Lick of Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399387919p2/9550.jpg",
		Content: "I will love you always. When this red hair is white, I will still love you. When the smooth softness of youth is replaced by the delicate softness of age, I will still want to touch your skin. When your face is full of the lines of every smile you have ever smiled, of every surprise I have seen flash through your eyes, when every tear you have ever cried has left its mark upon your face,I will treasure you all the more, because I was there to see it all. I will share your life with you, Meredith, and I will love you until the last breath leaves your body or mine.",
	},
	{
		Author: "Nicholas SparksDear John",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "I love you, not just for now, but for always, and I dream of the day that you’ll take me in your arms again.",
	},
	{
		Author: "Lawrence Ferlinghetti",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289676757p2/29970.jpg",
		Content: "If you're too open-minded; your brains will fall out.",
	},
	{
		Author: "Phyllis McGinley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363243791p2/62056.jpg",
		Content: "A bit of trash now and then is good for the severest reader. It provides the necessary roughage in the literary diet.",
	},
	{
		Author: "Caitlyn SiehlLiterary Sexts: A Collection of Short \u0026 Sexy Love Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399734312p2/7145854.jpg",
		Content: "Do not fall in love with people like me. I will take you to museums, and parks, and monuments, and kiss you in every beautiful place, so that you can never go back to them without tasting me like blood in your mouth.I will destroy you in the most beautiful way possible. And when I leave you will finally understand, why storms are named after people.",
	},
	{
		Author: "Charles BukowskiTales of Ordinary Madness",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "the free soul is rare, but you know it when you see it - basically because you feel good, very good, when you are near or with them.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I am too fond of reading books to care to write them.",
	},
	{
		Author: "Nicole KraussThe History of Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285130428p2/2633.jpg",
		Content: "When will you learn that there isn't a word for everything?",
	},
	{
		Author: "Lemony SnicketThe Beatrice Letters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Strange as it may seem, I still hope for the best, even though the best, like an interesting piece of mail, so rarely arrives, and even when it does it can be lost so easily.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I am only responsible for my own heart, you offered yours up for the smashing my darling. Only a fool would give out such a vital organ",
	},
	{
		Author: "George S. Patton Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1226608647p2/370054.jpg",
		Content: "If everybody is thinking alike, then somebody isn't thinking.",
	},
	{
		Author: "Kurt Vonnegut",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "Of all the words of mice and men, the saddest are, \"It might have been.",
	},
	{
		Author: "Christopher HitchensThe Portable Atheist: Essential Readings for the Nonbeliever",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434046956p2/3956.jpg",
		Content: "Owners of dogs will have noticed that, if you provide them with food and water and shelter and affection, they will think you are god. Whereas owners of cats are compelled to realize that, if you provide them with food and water and shelter and affection, they draw the conclusion that they are gods.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Tess, Tess, Tessa. Was there ever a more beautiful sound than your name? To speak it aloud makes my heart ring like a bell. Strange to imagine that, isn’t it – a heart ringing – but when you touch me that is what it is like: as if my heart is ringing in my chest and the sound shivers down my veins and splinters my bones with joy.Why have I written these words in this book? Because of you. You taught me to love this book where I had scorned it. When I read it for the second time, with an open mind and heart, I felt the most complete despair and envy of Sydney Carton. Yes, Sydney, for even if he had no hope that the woman he loved would love him, at least he could tell her of his love. At least he could do something to prove his passion, even if that thing was to die.I would have chosen death for a chance to tell you the truth, Tessa, if I could have been assured that death would be my own. And that is why I envied Sydney, for he was free.And now at last I am free, and I can finally tell you, without fear of danger to you, all that I feel in my heart.You are not the last dream of my soul.You are the first dream, the only dream I ever was unable to stop myself from dreaming. You are the first dream of my soul, and from that dream I hope will come all other dreams, a lifetime’s worth.With hope at least,Will Herondale",
	},
	{
		Author: "Stieg LarssonThe Girl with the Dragon Tattoo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246466225p2/706255.jpg",
		Content: "Friendship- my definition- is built on two things. Respect and trust. Both elements have to be there. And it has to be mutual. You can have respect for someone, but if you don't have trust, the friendship will crumble.",
	},
	{
		Author: "Ray Bradbury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "You must write every single day of your life... You must lurk in libraries and climb the stacks like ladders to sniff books like perfumes and wear books like hats upon your crazy heads... may you be in love every day for the next 20,000 days. And out of that love, remake a world.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "It takes courage to love, but pain through love is the purifying fire which those who love generously know. We all know people who are so much afraid of pain that they shut themselves up like clams in a shell and, giving out nothing, receive nothing and therefore shrink until life is a mere living death.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Why are you worrying about YOU-KNOW-WHO, when you should be worrying about YOU-NO-POO? The constipation sensation that's gripping the nation!",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "A woman's heart should be so hidden in God that a man has to seek Him just to find her.",
	},
	{
		Author: "Arundhati RoyThe God of Small Things",
		AuthorAvatar: "https://images.gr-assets.com/authors/1496705394p2/6134.jpg",
		Content: "That's what careless words do. They make people love you a little less.",
	},
	{
		Author: "Dante Alighieri",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1426181539p2/372374.jpg",
		Content: "My course is set for an uncharted sea.",
	},
	{
		Author: "Tim Burton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316033812p2/5773.jpg",
		Content: "One person's craziness is another person's reality.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "silence is the language of god, all else is poor translation.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Knowing too much of your future is never a good thing.",
	},
	{
		Author: "Hunter S. Thompson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206560814p2/5237.jpg",
		Content: "I hate to advocate drugs, alcohol, violence, or insanity to anyone, but they've always worked for me.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "I'm a success today because I had a friend who believed in me and I didn't have the heart to let him down.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I don't know who invented high heels, but all women owe him a lot!",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Forgive, O Lord, my little jokes on TheeAnd I'll forgive Thy great big one on me.",
	},
	{
		Author: "Jo WaltonAmong Others",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353809579p2/107170.jpg",
		Content: "It doesn't matter. I have books, new books, and I can bear anything as long as there are books.",
	},
	{
		Author: "Rita Mae BrownRiding Shotgun",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209493600p2/23511.jpg",
		Content: "Sorrow is how we learn to love. Your heart isn't breaking. It hurts because it's getting larger. The larger it gets, the more love it holds.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Do not think that love in order to be genuine has to be extraordinary. What we need is to love without getting tired. Be faithful in small things because it is in them that your strength lies.",
	},
	{
		Author: "Khaled HosseiniThe Kite Runner",
		AuthorAvatar: "https://images.gr-assets.com/authors/1359753468p2/569.jpg",
		Content: "There is only one sin. and that is theft... when you tell a lie, you steal someones right to the truth.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Raise your words, not voice. It is rain that grows flowers, not thunder.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Sometimes things fall apart so that better things can fall together",
	},
	{
		Author: "Jack KornfieldBuddha's Little Instruction Book",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494612521p2/59705.jpg",
		Content: "In the endthese things matter most:How well did you love?How fully did you live?How deeply did you let go?",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "I am looking for friends. What does that mean -- tame?\"\"It is an act too often neglected,\" said the fox. \"It means to establish ties.\" \"To establish ties?\" \"Just that,\" said the fox. \"To me, you are still nothing more than a little boy who is just like a hundred thousand other little boys. And I have no need of you. And you, on your part, have no need of me. To you I am nothing more than a fox like a hundred thousand other foxes. But if you tame me, then we shall need each other. To me, you will be unique in all the world. To you, I shall be unique in all the world....",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "No matter how careful you are, there's going to be the sense you missed something, the collapsed feeling under your skin that you didn't experience it all. There's that fallen heart feeling that you rushed right through the moments where you should've been paying attention.Well, get used to that feeling. That's how your whole life will feel some day.This is all practice.",
	},
	{
		Author: "RumiMasnavi i Man'avi, the spiritual couplets of Maula",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Sell your cleverness and buy bewilderment.",
	},
	{
		Author: "Cornelia FunkeInkheart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437000100p2/15873.jpg",
		Content: "Books have to be heavy because the whole world's inside them.",
	},
	{
		Author: "Georg Wilhelm Friedrich Hegel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1221544639p2/6188.jpg",
		Content: "Nothing great in the world was accomplished without passion.",
	},
	{
		Author: "Edward Abbey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1516467601p2/37218.jpg",
		Content: "Better a cruel truth than a comfortable delusion.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "What happened down in the dungeons between you and Professor Quirrell is a complete secret, so, naturally the whole school knows.",
	},
	{
		Author: "Heinrich Heine",
		AuthorAvatar: "https://images.gr-assets.com/authors/1310406603p2/16071.jpg",
		Content: "Where words leave off, music begins.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "When people talk, listen completely. Most people never listen.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "You can make anything by writing.",
	},
	{
		Author: "Anne Rice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383250078p2/7577.jpg",
		Content: "You do have a story inside you; it lies articulate and waiting to be written — behind your silence and your suffering.",
	},
	{
		Author: "Robert Louis StevensonEssays of Robert Louis Stevenson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192746024p2/854076.jpg",
		Content: "I kept always two books in my pocket, one to read, one to write in.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "You see I usually find myself among strangers because I drift here and there trying to forget the sad things that happened to me.",
	},
	{
		Author: "Terry PratchettReaper Man",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Light thinks it travels faster than anything but it is wrong. No matter how fast light travels, it finds the darkness has always got there first, and is waiting for it.",
	},
	{
		Author: "Jodi PicoultMy Sister's Keeper",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "If you have a sister and she dies, do you stop saying you have one? Or are you always a sister, even when the other half of the equation is gone?",
	},
	{
		Author: "Patrick RothfussThe Name of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351307341p2/108424.jpg",
		Content: "Perhaps the greatest faculty our minds possess is the ability to cope with pain. Classic thinking teaches us of the four doors of the mind, which everyone moves through according to their need.First is the door of sleep. Sleep offers us a retreat from the world and all its pain. Sleep marks passing time, giving us distance from the things that have hurt us. When a person is wounded they will often fall unconscious. Similarly, someone who hears traumatic news will often swoon or faint. This is the mind's way of protecting itself from pain by stepping through the first door.Second is the door of forgetting. Some wounds are too deep to heal, or too deep to heal quickly. In addition, many memories are simply painful, and there is no healing to be done. The saying 'time heals all wounds' is false. Time heals most wounds. The rest are hidden behind this door.Third is the door of madness. There are times when the mind is dealt such a blow it hides itself in insanity. While this may not seem beneficial, it is. There are times when reality is nothing but pain, and to escape that pain the mind must leave reality behind. Last is the door of death. The final resort. Nothing can hurt us after we are dead, or so we have been told.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "It's not the size of the dog in the fight, it's the size of the fight in the dog.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "My concern is not whether God is on our side; my greatest concern is to be on God's side, for God is always right.",
	},
	{
		Author: "Jim Butcher",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400640324p2/10746.jpg",
		Content: "When everything goes to hell, the people who stand by you without flinching -- they are your family. ",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Kind words can be short and easy to speak, but their echoes are truly endless.",
	},
	{
		Author: "L.M. MontgomeryAnne of Green Gables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188896723p2/5350.jpg",
		Content: "Kindred spirits are not so scarce as I used to think. It's splendid to find out there are so many of them in the world.",
	},
	{
		Author: "Edgar Allan PoeThe Raven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "Once upon a midnight dreary, while I pondered, weak and weary,Over many a quaint and curious volume of forgotten lore,While I nodded, nearly napping, suddenly there came a tapping,As of some one gently rapping, rapping at my chamber door.Tis some visitor,\" I muttered, \"tapping at my chamber door — Only this, and nothing more.\"Ah, distinctly I remember it was in the bleak December,And each separate dying ember wrought its ghost upon the floor.Eagerly I wished the morrow; — vainly I had sought to borrowFrom my books surcease of sorrow — sorrow for the lost Lenore —For the rare and radiant maiden whom the angels name Lenore — Nameless here for evermore.And the silken sad uncertain rustling of each purple curtainThrilled me — filled me with fantastic terrors never felt before;So that now, to still the beating of my heart, I stood repeating,Tis some visitor entreating entrance at my chamber door —Some late visitor entreating entrance at my chamber door; — This it is, and nothing more.\"Presently my soul grew stronger; hesitating then no longer,Sir,\" said I, \"or Madam, truly your forgiveness I implore;But the fact is I was napping, and so gently you came rapping,And so faintly you came tapping, tapping at my chamber door,That I scarce was sure I heard you\"— here I opened wide the door; — Darkness there, and nothing more.Deep into that darkness peering, long I stood there wondering, fearing,Doubting, dreaming dreams no mortals ever dared to dream before;But the silence was unbroken, and the stillness gave no token,And the only word there spoken was the whispered word, \"Lenore?\"This I whispered, and an echo murmured back the word, \"Lenore!\" — Merely this, and nothing more.Back into the chamber turning, all my soul within me burning,Soon again I heard a tapping somewhat louder than before.Surely,\" said I, \"surely that is something at my window lattice:Let me see, then, what thereat is, and this mystery explore —Let my heart be still a moment and this mystery explore; — 'Tis the wind and nothing more.\"Open here I flung the shutter, when, with many a flirt and flutter,In there stepped a stately raven of the saintly days of yore;Not the least obeisance made he; not a minute stopped or stayed he;But, with mien of lord or lady, perched above my chamber door —Perched upon a bust of Pallas just above my chamber door — Perched, and sat, and nothing more.Then this ebony bird beguiling my sad fancy into smiling,By the grave and stern decorum of the countenance it wore.Though thy crest be shorn and shaven, thou,\" I said, \"art sure no craven,Ghastly grim and ancient raven wandering from the Nightly shore —Tell me what thy lordly name is on the Night's Plutonian shore!\" Quoth the Raven, \"Nevermore.\"Much I marveled this ungainly fowl to hear discourse so plainly,Though its answer little meaning— little relevancy bore;For we cannot help agreeing that no living human beingEver yet was blest with seeing bird above his chamber door —Bird or beast upon the sculptured bust above his chamber door, With such name as \"Nevermore.",
	},
	{
		Author: "Dante AlighieriInferno",
		AuthorAvatar: "https://images.gr-assets.com/authors/1310943198p2/5031312.jpg",
		Content: "Do not be afraid; our fateCannot be taken from us; it is a gift.",
	},
	{
		Author: "Pablo Neruda100 Love Sonnets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "so I wait for you like a lonely housetill you will see me again and live in me.Till then my windows ache.",
	},
	{
		Author: "Honoré de BalzacPhysiologie Du Mariage",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206567834p2/228089.jpg",
		Content: "The more one judges, the less one loves.",
	},
	{
		Author: "Blaise Pascal",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397937717p2/10994.jpg",
		Content: "The heart has its reasons which reason knows not.",
	},
	{
		Author: "Ransom RiggsMiss Peregrine’s Home for Peculiar Children",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475462673p2/3046613.jpg",
		Content: "I used to dream about escaping my ordinary life, but my life was never ordinary. I had simply failed to notice how extraordinary it was.",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "I guess that’s just part of loving people: You have to give things up. Sometimes you even have to give them up.",
	},
	{
		Author: "Frances Hodgson BurnettThe Secret Garden",
		AuthorAvatar: "https://images.gr-assets.com/authors/1197934848p2/2041.jpg",
		Content: "If you look the right way, you can see that the whole world is a garden.",
	},
	{
		Author: "Hal Borland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363222258p2/217906.jpg",
		Content: "No winter lasts forever; no spring skips its turn.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "When the first baby laughed for the first time, its laugh broke into a thousand pieces, and they all went skipping about, and that was the beginning of fairies.",
	},
	{
		Author: "Lemony SnicketThe Grim Grotto",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "It is one of life's bitterest truths that bedtime so often arrives just when things are really getting interesting.",
	},
	{
		Author: "Lorraine Hansberry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234149147p2/3732.jpg",
		Content: "The thing that makes you exceptional, if you are at all, is inevitably that which must also make you lonely.",
	},
	{
		Author: "Rebecca West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392646885p2/8111.jpg",
		Content: "I myself have never been able to find out precisely what feminism is: I only know that people call me a feminist whenever I express sentiments that differentiate me from a doormat.",
	},
	{
		Author: "Agatha ChristieThe Man in the Mist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321738793p2/123715.jpg",
		Content: "Very few of us are what we seem.",
	},
	{
		Author: "Kate ChopinThe Awakening",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224805370p2/5132.jpg",
		Content: "The voice of the sea speaks to the soul.",
	},
	{
		Author: "Terry Pratchett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "It's not worth doing something unless someone, somewhere, would much rather you weren't doing it.",
	},
	{
		Author: "Antoine de Saint-ExupéryThe Little Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1330853515p2/1020792.jpg",
		Content: "People have forgotten this truth,\" the fox said. \"But you mustn’t forget it. You become responsible forever for what you’ve tamed. You’re responsible for your rose.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Magnus, standing by the door, snapped his fingers impatiently. \"Move it along, teenagers. The only person who gets to canoodle in my bedroom is my magnificent self.\"\"Canoodle?\" repeated Clary, never having heard the word before.\"Magnificent?\" repeated Jace, who was just being nasty. Magnus growled. The growl sounded like \"Get out.",
	},
	{
		Author: "Socrates",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390145726p2/275648.jpg",
		Content: "I cannot teach anybody anything. I can only make them think",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "those who escape hellhowevernever talk aboutitand nothing muchbothers themafterthat.",
	},
	{
		Author: "Reginald Vincent HolmesFireside Fancies",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1485706007p2/96600.jpg",
		Content: "The earth has its music for those who will listen",
	},
	{
		Author: "Peter S. BeagleThe Last Unicorn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1360970921p2/1067608.jpg",
		Content: "Great heroes need great sorrows and burdens, or half their greatness goes unnoticed. It is all part of the fairy tale.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "[D]on't ever apologise to an author for buying something in paperback, or taking it out from a library (that's what they're there for. Use your library). Don't apologise to this author for buying books second hand, or getting them from bookcrossing or borrowing a friend's copy. What's important to me is that people read the books and enjoy them, and that, at some point in there, the book was bought by someone. And that people who like things, tell other people. The most important thing is that people read...",
	},
	{
		Author: "Albert Schweitzer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198534412p2/47146.jpg",
		Content: "There are two means of refuge from the misery of life — music and cats.",
	},
	{
		Author: "Jonathan Safran Foer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "Time was passing like a hand waving from a train I wanted to be on. I hope you never have to think about anything as much as I think about you.",
	},
	{
		Author: "Louise ErdrichThe Plague of Doves",
		AuthorAvatar: "https://images.gr-assets.com/authors/1462224430p2/9388.jpg",
		Content: "When we are young, the words are scattered all around us. As they are assembled by experience, so also are we, sentence by sentence, until the story takes shape.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You might want to lie down,\" Magnus advised. \"I find that it helps when the crushing sense of horrible realization sets in.",
	},
	{
		Author: "Nicholas SparksAt First Sight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "The emotion that can break your heart is sometimes the very one that heals it...",
	},
	{
		Author: "Daniel J. Boorstin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1238770710p2/10378.jpg",
		Content: "The greatest enemy of knowledge is not ignorance, it is the illusion of knowledge.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "please believe that things are good with me, and even when they're not, they will be soon enough. And i will always believe the same about you.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "I don't go out with strangers,\" I said.\"Good thing I do. I'll pick you up at five.",
	},
	{
		Author: "J.R.R. TolkienThe Return of the King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "I am glad you are here with me. Here at the end of all things, Sam.",
	},
	{
		Author: "Bob Dylan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466390809p2/8898.jpg",
		Content: "Behind every beautiful thing, there's some kind of pain.",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "There's nowhere you can be that isn't where you're meant to be...",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "Between two evils, I always pick the one I never tried before.",
	},
	{
		Author: "Langston Hughes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206669966p2/36910.jpg",
		Content: "Let the rain kiss you. Let the rain beat upon your head with silver liquid drops. Let the rain sing you a lullaby.",
	},
	{
		Author: "J.K. Rowling",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "We do not need magic to transform our world. We carry all of the power we need inside ourselves already.",
	},
	{
		Author: "Charlotte Brontë",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "The trouble is not that I am single and likely to stay single, but that I am lonely and likely to stay lonely.",
	},
	{
		Author: "Pittacus LoreI Am Number Four",
		AuthorAvatar: "https://images.gr-assets.com/authors/1291480497p2/3380908.jpg",
		Content: "When you have lost hope, you have lost everything. And when you think all is lost, when all is dire and bleak, there is always hope.",
	},
	{
		Author: "Julio Cortázar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1496948373p2/25824.jpg",
		Content: "Come sleep with me: We won't make Love, Love will make us.",
	},
	{
		Author: "Cassandra ClareCity of Fallen Angels",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You're the first Shadowhunter I've ever met.\" “That’s too bad,",
	},
	{
		Author: "T.S. Eliot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1507887310p2/18540.jpg",
		Content: "This is the way the world endsNot with a bang but a whimper.",
	},
	{
		Author: "Randy PauschThe Last Lecture",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213769716p2/287960.jpg",
		Content: "The brick walls are there for a reason. The brick walls are not there to keep us out. The brick walls are there to give us a chance to show how badly we want something. Because the brick walls are there to stop the people who don’t want it badly enough. They’re there to stop the other people.",
	},
	{
		Author: "J.R.R. Tolkien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "Courage is found in unlikely places.",
	},
	{
		Author: "Edgar Allan Poe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454522972p2/4624490.jpg",
		Content: "From childhood's hour I have not been. As others were, I have not seen. As others saw, I could not awaken. My heart to joy at the same tone. And all I loved, I loved alone.",
	},
	{
		Author: "Carlos Ruiz ZafónThe Shadow of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444199853p2/815.jpg",
		Content: "Every book, every volume you see here, has a soul. The soul of the person who wrote it and of those who read it and lived and dreamed with it. Every time a book changes hands, every time someone runs his eyes down its pages, its spirit grows and strengthens.",
	},
	{
		Author: "Edna St. Vincent Millay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183231710p2/33998.jpg",
		Content: "Where you used to be, there is a hole in the world, which I find myself constantly walking around in the daytime, and falling in at night. I miss you like hell.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "You may encounter many defeats, but you must not be defeated. In fact, it may be necessary to encounter the defeats, so you can know who you are, what you can rise from, how you can still come out of it.",
	},
	{
		Author: "Anne RiceInterview with the Vampire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383250078p2/7577.jpg",
		Content: "Evil is always possible. And goodness is eternally difficult.",
	},
	{
		Author: "Rainbow RowellEleanor \u0026 Park",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342324527p2/4208569.jpg",
		Content: "Holding Eleanor's hand was like holding a butterfly. Or a heartbeat. Like holding something complete, and completely alive.",
	},
	{
		Author: "Sun TzuThe Art of War",
		AuthorAvatar: "https://images.gr-assets.com/authors/1185211202p2/1771.jpg",
		Content: "Appear weak when you are strong, and strong when you are weak.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "The ultimate measure of a man is not where he stands in moments of comfort and convenience, but where he stands at times of challenge and controversy.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "I hope you're pleased with yourselves. We could all have been killed - or worse, expelled. Now if you don't mind, I'm going to bed.",
	},
	{
		Author: "Christopher Hitchens",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434046956p2/3956.jpg",
		Content: "That which can be asserted without evidence, can be dismissed without evidence.",
	},
	{
		Author: "Ashleigh Brilliant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215239131p2/15038.jpg",
		Content: "My life has a superb cast, but I cannot figure out the plot.",
	},
	{
		Author: "J.R. WardLover Awakened",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437768069p2/20248.jpg",
		Content: "I was dead until you found me, though I breathed. I was sightless, though I could see. And then you came...and I was awakened.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "They say I'm old-fashioned, and live in the past, but sometimes I think progress progresses too fast!",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "I am always saying \"Glad to've met you\" to somebody I'm not at all glad I met. If you want to stay alive, you have to say that stuff, though.",
	},
	{
		Author: "Guy de Maupassant",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188821941p2/18791.jpg",
		Content: "Words dazzle and deceive because they are mimed by the face. But black words on a white page are the soul laid bare.",
	},
	{
		Author: "Lewis CarrollAlice in Wonderland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "Curiouser and curiouser!",
	},
	{
		Author: "E.L. JamesFifty Shades of Grey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308409727p2/4725841.jpg",
		Content: "Sometimes I wonder if there's something wrong with me. Perhaps I've spent too long in the company of my literary romantic heroes, and consequently my ideals and expectations are far too high.",
	},
	{
		Author: "Bob Marley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207771636p2/25241.jpg",
		Content: "Better to die fighting for freedom then be a prisoner all the days of your life.",
	},
	{
		Author: "Henry James",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468309415p2/159.jpg",
		Content: "She feels in italics and thinks in CAPITALS.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "I've lived through some terrible things in my life, some of which actually happened.",
	},
	{
		Author: "Tom Robbins",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430660478p2/197.jpg",
		Content: "When we're incomplete, we're always searching for somebody to complete us. When, after a few years or a few months of a relationship, we find that we're still unfulfilled, we blame our partners and take up with somebody more promising. This can go on and on--series polygamy--until we admit that while a partner can add sweet dimensions to our lives, we, each of us, are responsible for our own fulfillment. Nobody else can provide it for us, and to believe otherwise is to delude ourselves dangerously and to program for eventual failure every relationship we enter.",
	},
	{
		Author: "Maya AngelouLetter to My Daughter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Try to be a rainbow in someone's cloud.",
	},
	{
		Author: "Kurt VonnegutSlaughterhouse-Five",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "How nice -- to feel nothing, and still get full credit for being alive.",
	},
	{
		Author: "Charles DarwinThe Autobiography of Charles Darwin, 1809–82",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398693802p2/12793.jpg",
		Content: "If I had my life to live over again, I would have made a rule to read some poetry and listen to some music at least once every week.",
	},
	{
		Author: "Thomas Babington MacaulayThe Selected Letters Of Thomas Babington Macaulay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1293651286p2/234077.jpg",
		Content: "What a blessing it is to love books as I love them;- to be able to converse with the dead, and to live amidst the unreal!",
	},
	{
		Author: "Paulo Coelho",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "If I am really a part of your dream, you'll come back one day.",
	},
	{
		Author: "Johnny Depp",
		AuthorAvatar: "https://images.gr-assets.com/authors/1317565242p2/88280.jpg",
		Content: "If someone were to harm my family or a friend or somebody I love, I would eat them. I might end up in jail for 500 years, but I would eat them.",
	},
	{
		Author: "Simone ElkelesRules of Attraction",
		AuthorAvatar: "https://images.gr-assets.com/authors/1350327003p2/274533.jpg",
		Content: "You’re dangerous,",
	},
	{
		Author: "Kurt Cobain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437780927p2/33041.jpg",
		Content: "I'd rather be hated for who I am, than loved for who I am not.",
	},
	{
		Author: "Neil GaimanDream Country",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Things need not have happened to be true. Tales and dreams are the shadow-truths that will endure when mere facts are dust and ashes, and forgot.",
	},
	{
		Author: "Ray BradburyFahrenheit 451",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "Everyone must leave something behind when he dies, my grandfather said. A child or a book or a painting or a house or a wall built or a pair of shoes made. Or a garden planted. Something your hand touched some way so your soul has somewhere to go when you die, and when people look at that tree or that flower you planted, you're there. It doesn't matter what you do, he said, so long as you change something from the way it was before you touched it into something that's like you after you take your hands away. The difference between the man who just cuts lawns and a real gardener is in the touching, he said. The lawn-cutter might just as well not have been there at all; the gardener will be there a lifetime.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Selfishness is not living as one wishes to live, it is asking others to live as one wishes to live.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "You are not special. You're not a beautiful and unique snowflake. You're the same decaying organic matter as everything else. We're all part of the same compost heap. We're all singing, all dancing crap of the world.",
	},
	{
		Author: "Emily BrontëWuthering Heights",
		AuthorAvatar: "https://images.gr-assets.com/authors/1473229007p2/4191.jpg",
		Content: "I wish I were a girl again, half-savage and hardy, and free.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "Trust your heart if the seas catch fire, live by love though the stars walk backward.",
	},
	{
		Author: "Richelle MeadShadow Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "You will lose what you value most, so treasure it while you can.",
	},
	{
		Author: "Woody Allen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501608198p2/10356.jpg",
		Content: "In my next life I want to live my life backwards. You start out dead and get that out of the way. Then you wake up in an old people's home feeling better every day. You get kicked out for being too healthy, go collect your pension, and then when you start work, you get a gold watch and a party on your first day. You work for 40 years until you're young enough to enjoy your retirement. You party, drink alcohol, and are generally promiscuous, then you are ready for high school. You then go to primary school, you become a kid, you play. You have no responsibilities, you become a baby until you are born. And then you spend your last 9 months floating in luxurious spa-like conditions with central heating and room service on tap, larger quarters every day and then Voila! You finish off as an orgasm!",
	},
	{
		Author: "James Baldwin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343346341p2/10427.jpg",
		Content: "Children have never been very good at listening to their elders, but they have never failed to imitate them.",
	},
	{
		Author: "Steve Maraboli",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "Cry. Forgive. Learn. Move on. Let your tears water the seeds of your future happiness.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I am lonely, yet not everybody will do. I don't know why, some people fill the gaps and others emphasize my loneliness. In reality those who satisfy me are those who simply allow me to live with my ''idea of them.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "There are only the pursued, the pursuing, the busy and the tired.",
	},
	{
		Author: "R.A. SalvatoreSojourn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207159077p2/1023510.jpg",
		Content: "It is better, I think, to grab at the stars than to sit flustered because you know you cannot reach them.",
	},
	{
		Author: "James Joyce",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517863935p2/5144.jpg",
		Content: "Shut your eyes and see.",
	},
	{
		Author: "Joss Whedon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1302721520p2/18015.jpg",
		Content: "I write to give myself strength. I write to be the characters that I am not. I write to explore all the things I'm afraid of. ",
	},
	{
		Author: "Harvey Fierstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206735024p2/4775.jpg",
		Content: "Never be bullied into silence. Never allow yourself to be made a victim. Accept no one's definition of your life, but define yourself.",
	},
	{
		Author: "Terry McMillanDisappearing Acts",
		AuthorAvatar: "https://images.gr-assets.com/authors/1462816905p2/20704.jpg",
		Content: "Too many of us are hung up on what we don't have, can't have, or won't ever have. We spend too much energy being down, when we could use that same energy – if not less of it – doing, or at least trying to do, some of the things we really want to do.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I know it's wrong - God, it's all kinds of wrong - but I just want to lie down with you and wake up with you, just once, just once ever in my life.",
	},
	{
		Author: "Louisa May AlcottLittle Women",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200326665p2/1315.jpg",
		Content: "I like good strong words that mean something…",
	},
	{
		Author: "René Descartes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198124305p2/36556.jpg",
		Content: "I think; therefore I am.",
	},
	{
		Author: "Victor Hugo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "To put everything in balance is good, to put everything in harmony is better.",
	},
	{
		Author: "Tahereh MafiUnravel Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "Books are easily destroyed. But words will live as long as people can remember them.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Alec looked at her and shook his head. \"How do you manage never to get mud on your clothes?\"Isabelle shrugged philosophically. \"I'm pure at heart. It repels the dirt.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "We sit silently and watch the world around us. This has taken a lifetime to learn. It seems only the old are able to sit next to one another and not say anything and still feel content. The young, brash and impatient, must always break the silence. It is a waste, for silence is pure. Silence is holy. It draws people together because only those who are comfortable with each other can sit without speaking. This is the great paradox.",
	},
	{
		Author: "Sigmund Freud",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406688955p2/10017.jpg",
		Content: "Being entirely honest with oneself is a good exercise.",
	},
	{
		Author: "Richelle MeadBlood Promise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "I know how devastated you must be to miss me, but leave a message, and I'll try to ease your agony",
	},
	{
		Author: "Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Is there no way out of the mind?",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Only you can control your future.",
	},
	{
		Author: "Katharine Graham",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371404590p2/54931.jpg",
		Content: "The longer I live, the more I observe that carrying around anger is the most debilitating to the person who bears it.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "Writing is like sex. First you do it for love, then you do it for your friends, and then you do it for money.",
	},
	{
		Author: "Lao TzuTao Teh Ching",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "Those who know do not speak. Those who speak do not know.",
	},
	{
		Author: "Seán O'Casey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1251317358p2/11582.jpg",
		Content: "Laughter is wine for the soul - laughter soft, or loud and deep, tinged through with seriousness - the hilarious declaration made by man that life is worth living.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "And then something invisible snapped insider her, and that which had come together commenced to fall apart.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I hope you will have a wonderful year, that you'll dream dangerously and outrageously, that you'll make something that didn't exist before you made it, that you will be loved and that you will be liked, and that you will have people to love and to like in return. And, most importantly (because I think there should be more kindness and more wisdom in the world right now), that you will, when you need to be, be wise, and that you will always be kind.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Sometimes, when I have to do something I don't want to do, I pretend I'm a character from a book. It's easier to know what they would do.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Trains are great dirty smoky things,\" said Will. \"You won't like it.\" Tessa was unmoved. \"I won't know if I like it until I try it, will I?\" \"I've never swum naked in the Thames before, but I know I wouldn't like it.\" \"But think how entertaining for sightseers,\" said Tessa, and she saw Jem duck his head to hide the quick flash of his grin.",
	},
	{
		Author: "Jonathan Safran FoerEverything Is Illuminated",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "He awoke each morning with the desire to do right, to be a good and meaningful person, to be, as simple as it sounded and as impossible as it actually was, happy. And during the course of each day his heart would descend from his chest into his stomach. By early afternoon he was overcome by the feeling that nothing was right, or nothing was right for him, and by the desire to be alone. By evening he was fulfilled: alone in the magnitude of his grief, alone in his aimless guilt, alone even in his loneliness. I am not sad, he would repeat to himself over and over, I am not sad. As if he might one day convince himself. Or fool himself. Or convince others--the only thing worse than being sad is for others to know that you are sad. I am not sad. I am not sad. Because his life had unlimited potential for happiness, insofar as it was an empty white room. He would fall asleep with his heart at the foot of his bed, like some domesticated animal that was no part of him at all. And each morning he would wake with it again in the cupboard of his rib cage, having become a little heavier, a little weaker, but still pumping. And by the midafternoon he was again overcome with the desire to be somewhere else, someone else, someone else somewhere else. I am not sad.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Gus: \"It tastes like...\"Me: \"Food.\"Gus: \"Yes, precisely. It tastes like food, excellently prepared. But it does not taste, how do I put this delicately...?\"Me: \"It does not taste like God Himself cooked heaven into a series of five dishes which were then served to you accompanied by several luminous balls of fermented, bubbly plasma while actual and literal flower petals floated down around your canal-side dinner table.\"Gus: \"Nicely phrased.\"Gus's father: \"Our children are weird.\"My dad: \"Nicely phrased.",
	},
	{
		Author: "Alexandre DumasThe Count of Monte Cristo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1279049943p2/4785.jpg",
		Content: "I am not proud, but I am happy; and happiness blinds, I think, more than pride.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "All a girl really wants is for one guy to prove to her that they are not all the same.",
	},
	{
		Author: "David Henry HwangM. Butterfly",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335646977p2/55111.jpg",
		Content: "I'm happy. Which often looks like crazy.",
	},
	{
		Author: "Donald E. WestlakeTwo Much",
		AuthorAvatar: "https://images.gr-assets.com/authors/1336863543p2/30953.jpg",
		Content: "Nobody gets everything in this life. You decide your priorities and you make your choices. I'd decided long ago that any cake I had would be eaten.",
	},
	{
		Author: "Dr. SeussI Can Read With My Eyes Shut!",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "You’ll miss the best things if you keep your eyes shut.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Let the first act of every morning be to make the following resolve for the day:- I shall not fear anyone on Earth. - I shall fear only God. - I shall not bear ill will toward anyone. - I shall not submit to injustice from anyone. - I shall conquer untruth by truth. And in resisting untruth, I shall put up with all suffering.",
	},
	{
		Author: "Rick RiordanThe Lightning Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Even strength must bow to wisdom sometimes.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "If you never did you should. These things are fun and fun is good.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You could say sorry,\" suggested Harry bluntly. \"What, and get attacked by another flock of canaries?\" muttered Ron.\"What did you have to imitate her for?\"\"She laughed at my mustache!\"\"So did I, it was the stupidest thing I've ever seen.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "We all bear scars,... Mine just happen to be more visible than most.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "It's much easier to not know things sometimes. Things change and friends leave. And life doesn't stop for anybody. I wanted to laugh. Or maybe get mad. Or maybe shrug at how strange everybody was, especially me. I think the idea is that every person has to live for his or her own life and than make the choice to share it with other people. You can't just sit their and put everybody's lives ahead of yours and think that counts as love. You just can't. You have to do things. I'm going to do what I want to do. I'm going to be who I really am. And I'm going to figure out what that is. And we could all sit around and wonder and feel bad about each other and blame a lot of people for what they did or didn't do or what they didn't know. I don't know. I guess there could always be someone to blame. It's just different. Maybe it's good to put things in perspective, but sometimes, I think that the only perspective is to really be there. Because it's okay to feel things. I was really there. And that was enough to make me feel infinite. I feel infinite.",
	},
	{
		Author: "Edmund Burke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216482471p2/17142.jpg",
		Content: "The only thing necessary for the triumph of evil is for good men to do nothing.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "If you're texting Magnus to say 'I think u r kewl,' I'm going to kill you.\"\"Who's Magnus?\" Max inquired.\"He's a warlock,\" said Alec.\"A sexy, sexy warlock,\" Isabelle told Max, ignoring Alec's look of total fury.\"But warlocks are bad,\" protested Max, looking baffled.\"Exactly,\" said Isabelle.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "There are nights when the wolves are silent and only the moon howls. ",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Half-Blood Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "It's going to be all right, sir,\" Harry said over and over again, more worried by Dumbledore's silence than he had been by his weakened voice. \"We're nearly there ... I can Apparate us both back ... don't worry ...\"\"I am not worried, Harry,\" said Dumbledore, his voice a little stronger despite the freezing water. \"I am with you.",
	},
	{
		Author: "Melissa MarrInk Exchange",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192302741p2/175855.jpg",
		Content: "Sometimes love means letting go when you want to hold on tighter.",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "Ever has it been that love knows not its own depth until the hour of separation.",
	},
	{
		Author: "Warsan Shire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1358876300p2/5431334.jpg",
		Content: "My alone feels so good, I'll only have you if you're sweeter than my solitude.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Don't give in to your fears. If you do, you won't be able to talk to your heart.",
	},
	{
		Author: "Andy WarholThe Philosophy of Andy Warhol",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206565218p2/1203.jpg",
		Content: "They always say time changes things, but you actually have to change them yourself.",
	},
	{
		Author: "Gary PaulsenThe Winter Room",
		AuthorAvatar: "https://images.gr-assets.com/authors/1309159225p2/18.jpg",
		Content: "If books could have more, give more, be more, show more, they would still need readers who bring to them sound and smell and light and all the rest that can’t be in books.\tThe book needs you.",
	},
	{
		Author: "Pat  Monahan",
		AuthorAvatar: "",
		Content: "You can only be young once. But you can always be immature.",
	},
	{
		Author: "Terry PratchettGood Omens: The Nice and Accurate Prophecies of Agnes Nutter, Witch",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "DON'T THINK OF IT AS DYING, said Death. JUST THINK OF IT AS LEAVING EARLY TO AVOID THE RUSH.",
	},
	{
		Author: "Erma BombeckWhen God Created Mothers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208791191p2/11882.jpg",
		Content: "When God Created Mothers\"When the Good Lord was creating mothers, He was into His sixth day of \"overtime\" when the angel appeared and said. \"You're doing a lot of fiddling around on this one.\" And God said, \"Have you read the specs on this order?\" She has to be completely washable, but not plastic. Have 180 moveable parts...all replaceable. Run on black coffee and leftovers. Have a lap that disappears when she stands up. A kiss that can cure anything from a broken leg to a disappointed love affair. And six pairs of hands.\" The angel shook her head slowly and said. \"Six pairs of hands.... no way.\" It's not the hands that are causing me problems,\" God remarked, \"it's the three pairs of eyes that mothers have to have.\" That's on the standard model?\" asked the angel. God nodded. One pair that sees through closed doors when she asks, 'What are you kids doing in there?' when she already knows. Another here in the back of her head that sees what she shouldn't but what she has to know, and of course the ones here in front that can look at a child when he goofs up and say. 'I understand and I love you' without so much as uttering a word.\" God,\" said the angel touching his sleeve gently, \"Get some rest tomorrow....\" I can't,\" said God, \"I'm so close to creating something so close to myself. Already I have one who heals herself when she is sick...can feed a family of six on one pound of hamburger...and can get a nine year old to stand under a shower.\" The angel circled the model of a mother very slowly. \"It's too soft,\" she sighed. But tough!\" said God excitedly. \"You can imagine what this mother can do or endure.\" Can it think?\" Not only can it think, but it can reason and compromise,\" said the Creator. Finally, the angel bent over and ran her finger across the cheek. There's a leak,\" she pronounced. \"I told You that You were trying to put too much into this model.\" It's not a leak,\" said the Lord, \"It's a tear.\" What's it for?\" It's for joy, sadness, disappointment, pain, loneliness, and pride.\" You are a genius, \" said the angel. Somberly, God said, \"I didn't put it there.",
	},
	{
		Author: "Kazuo IshiguroNever Let Me Go",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424906625p2/4280.jpg",
		Content: "Memories, even your most precious ones, fade surprisingly quickly. But I don’t go along with that. The memories I value most, I don’t ever see them fading.",
	},
	{
		Author: "Annie Proulx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1219720509p2/1262010.jpg",
		Content: "You should write because you love the shape of stories and sentences and the creation of different words on a page. Writing comes from reading, and reading is the finest teacher of how to write.",
	},
	{
		Author: "Viktor E. FranklMan's Search for Meaning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1218372404p2/2782.jpg",
		Content: "When we are no longer able to change a situation, we are challenged to change ourselves.",
	},
	{
		Author: "Sheng Wang",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376842453p2/7009636.jpg",
		Content: "Why do people say \"grow some balls\"? Balls are weak and sensitive. If you wanna be tough, grow a vagina. Those things can take a pounding.",
	},
	{
		Author: "Jane AustenLove and Friendship",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "The Very first moment I beheld him, my heart was irrevocably gone.",
	},
	{
		Author: "Geraldine BrooksMarch",
		AuthorAvatar: "https://images.gr-assets.com/authors/1303284528p2/211268.jpg",
		Content: "For to know a man's library is, in some measure, to know his mind.",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "I hope that one day you will have the experience of doing something you do not understand for someone you love.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Talking to a drunk person was like talking to an extremely happy, severely brain-damaged three-year-old.",
	},
	{
		Author: "Nicholas SparksThree Weeks With My Brother",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Dreams are always crushing when they don't come true. But it's the simple dreams that are often the most painful because they seem so personal, so reasonable, so attainable. You're always close enough to touch, but never quite close enough to hold and it's enough to break your heart.",
	},
	{
		Author: "Pablo Neruda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982432p2/4026.jpg",
		Content: "But I love your feet only because they walked upon the earth and upon the wind and upon the waters, until they found me.",
	},
	{
		Author: "Molière",
		AuthorAvatar: "https://images.gr-assets.com/authors/1219168412p2/29837.jpg",
		Content: "Trees that are slow to grow bear the best fruit.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "Life is too short to waste any amount of time on wondering what other people think about you. In the first place, if they had better things going on in their lives, they wouldn't have the time to sit around and talk about you. What's important to me is not others' opinions of me, but what's important to me is my opinion of myself.",
	},
	{
		Author: "James  Burke",
		AuthorAvatar: "https://images.gr-assets.com/authors/1222103880p2/989318.jpg",
		Content: "When you read a book, you hold another's mind in your hands.",
	},
	{
		Author: "Ayn RandThe Fountainhead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "I could die for you. But I couldn't, and wouldn't, live for you.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Some walks you have to take alone.",
	},
	{
		Author: "Eleanor H. Porter",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1402418285p2/28504.jpg",
		Content: "Be glad. Be good. Be brave.",
	},
	{
		Author: "Colleen HooverSlammed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1464032240p2/5430144.jpg",
		Content: "I got schooled this yearbyaboy.A boy that I'm seriously, deeply, madly, incredibly, and undeniably in \n  love\n with.And he taught me the most important thing of all...To put the emphasisOn \n  life\n.",
	},
	{
		Author: "Haruki MurakamiKafka on the Shore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "Lost opportunities, lost possibilities, feelings we can never get back. That's part of what it means to be alive. But inside our heads - at least that's where I imagine it - there's a little room where we store those memories. A room like the stacks in this library. And to understand the workings of our own heart we have to keep on making new reference cards. We have to dust things off every once in awhile, let in fresh air, change the water in the flower vases. In other words, you'll live forever in your own private library.",
	},
	{
		Author: "Bill Watterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "I'm killing time while I wait for life to shower me with meaning and happiness.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I don't mind making jokes, but I don't want to look like one.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "I have nothing to declare except my genius.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Everything must be made as simple as possible. But not simpler.",
	},
	{
		Author: "Steven Spielberg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1205428805p2/174694.jpg",
		Content: "Only a generation of readers will spawn a generation of writers.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Have you tried talking to her?\" \"No. We've been punching her in the face repeatedly. What? You don't think that will work?",
	},
	{
		Author: "Santosh KalwarQuote Me Everyday",
		AuthorAvatar: "https://images.gr-assets.com/authors/1275309637p2/2894169.jpg",
		Content: "We are addicted to our thoughts. We cannot change anything if we cannot change our thinking.",
	},
	{
		Author: "Jean de La FontaineFables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1247678674p2/82948.jpg",
		Content: "A person often meets his destiny on the road he took to avoid it.",
	},
	{
		Author: "Leo TolstoyWar and Peace",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "We can know only that we know nothing. And that is the highest degree of human wisdom.",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "If you liked being a teenager, there's something really wrong with you.",
	},
	{
		Author: "Rainbow RowellFangirl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342324527p2/4208569.jpg",
		Content: "I don’t trust anybody. Not anybody. And the more that I care about someone, the more sure I am they’re going to get tired of me and take off.",
	},
	{
		Author: "Anne FrankThe Diary of a Young Girl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "I don't think of all the misery, but of the beauty that still remains.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "I have for the first time found what I can truly love–I have found you. You are my sympathy–my better self–my good angel–I am bound to you with a strong attachment. I think you good, gifted, lovely: a fervent, a solemn passion is conceived in my heart; it leans to you, draws you to my centre and spring of life, wrap my existence about you–and, kindling in pure, powerful flame, fuses you and me in one.",
	},
	{
		Author: "Bei Dao",
		AuthorAvatar: "https://images.gr-assets.com/authors/1319659394p2/3011429.jpg",
		Content: "In the world I amAlways a strangerI do not understand its languageIt does not understand my silence",
	},
	{
		Author: "Alexander PopeAn Essay on Criticism",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206647570p2/25157.jpg",
		Content: "To err is human, to forgive, divine.",
	},
	{
		Author: "Paula Poundstone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216752505p2/38161.jpg",
		Content: "Adults are always asking little kids what they want to be when they grow up ’cause they’re looking for ideas.",
	},
	{
		Author: "William Makepeace Thackeray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204650429p2/3953.jpg",
		Content: "Life is a mirror: if you frown at it, it frowns back; if you smile, it returns the greeting.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I’m on a roller coaster that only goes up, my friend.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The best way to cheer yourself is to cheer somebody else up.",
	},
	{
		Author: "Michel de Montaigne",
		AuthorAvatar: "https://images.gr-assets.com/authors/1332479195p2/17241.jpg",
		Content: "The most certain sign of wisdom is cheerfulness. ",
	},
	{
		Author: "Edith Wharton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1484512230p2/16.jpg",
		Content: "There are two ways of spreading light: to be the candle or the mirror that receives it.",
	},
	{
		Author: "Carl SaganCosmos",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475953320p2/10538.jpg",
		Content: "The nitrogen in our DNA, the calcium in our teeth, the iron in our blood, the carbon in our apple pies were made in the interiors of collapsing stars. We are made of starstuff.",
	},
	{
		Author: "Candace BushnellSex and the City",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435063350p2/4415.jpg",
		Content: "Maybe our girlfriends are our soulmates and guys are just people to have fun with.",
	},
	{
		Author: "Charles DickensA Tale of Two Cities",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "A wonderful fact to reflect upon, that every human creature is constituted to be that profound secret and mystery to every other.",
	},
	{
		Author: "J.M. BarriePeter Pan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1519029719p2/5255014.jpg",
		Content: "Second star to the right and straight on 'til morning. ",
	},
	{
		Author: "Richard MathesonWhat Dreams May Come",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200467797p2/8726.jpg",
		Content: "That which you believe becomes your world.",
	},
	{
		Author: "Gayle FormanIf I Stay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1472502468p2/295178.jpg",
		Content: "If you stay, I'll do whatever you want. I'll quit the band, go with you to New York. But if you need me to go away, I'll do that, too. I was talking to Liz and she said maybe coming back to your old life would be too painful, that maybe it'd be easier for you to erase us. And that would suck, but I'd do it. I can lose you like that if I don't lose you today. I'll let you go. If you stay.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "I feel like someone breathed new air into my lungs. I am not Abnegation. I am not Dauntless. I am Divergent.",
	},
	{
		Author: "Patrick RothfussThe Wise Man's Fear",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351307341p2/108424.jpg",
		Content: "It's the questions we can't answer that teach us the most. They teach us how to think. If you give a man an answer, all he gains is a little fact. But give him a question and he'll look for his own answers.",
	},
	{
		Author: "Caroline Gordon",
		AuthorAvatar: "",
		Content: "A well-composed book is a magic carpet on which we are wafted to a world that we cannot enter in any other way.",
	},
	{
		Author: "Mitch AlbomFor One More Day",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "When someone is in your heart, they're never truly gone. They can come back to you, even at unlikely times.",
	},
	{
		Author: "Jane YolenTouch Magic: Fantasy, Faerie \u0026 Folklore in the Literature of Childhood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1413465729p2/5989.jpg",
		Content: "Literature is a textually transmitted disease, normally contracted in childhood.",
	},
	{
		Author: "Ray Bradbury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "There is more than one way to burn a book. And the world is full of people running about with lit matches.",
	},
	{
		Author: "Laura Ingalls Wilder",
		AuthorAvatar: "https://images.gr-assets.com/authors/1347574987p2/5300.jpg",
		Content: "I am beginning to learn that it is the sweet, simple things of life which are the real ones after all.",
	},
	{
		Author: "Louis L'AmourMatagorda/The First Fast Draw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343675199p2/858.jpg",
		Content: "Once you have read a book you care about, some part of it is always with you.",
	},
	{
		Author: "Mitch AlbomThe Five People You Meet in Heaven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "Holding anger is a poison...It eats you from inside...We think that by hating someone we hurt them...But hatred is a curved blade...and the harm we do to others...we also do to ourselves.",
	},
	{
		Author: "C.G. Jung",
		AuthorAvatar: "https://images.gr-assets.com/authors/1471024439p2/38285.jpg",
		Content: "Everything that irritates us about others can lead us to an understanding of ourselves.",
	},
	{
		Author: "Richelle MeadShadow Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "What's up?\" I asked.You tell me,\" he said. \"You were the one about ready to start making out with Adrian.\"It was an experiment,\" I said. \"It was part of my therapy.\"What the hell kind of therapy are you in?",
	},
	{
		Author: "Ursula K. Le GuinThe Left Hand of Darkness",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244291425p2/874602.jpg",
		Content: "It is good to have an end to journey toward; but it is the journey that matters, in the end.",
	},
	{
		Author: "Charles M. Schulz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590750p2/209672.jpg",
		Content: "Nothing takes the taste out of peanut butter quite like unrequited love.",
	},
	{
		Author: "Lauren OliverBefore I Fall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "I guess that's what saying good-bye is always like--like jumping off an edge. The worst part is making the choice to do it. Once you're in the air, there's nothing you can do but let go.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "No sight so sad as that of a naughty child,\" he began, \"especially a naughty little girl. Do you know where the wicked go after death?\"\"They go to hell,\" was my ready and orthodox answer.\"And what is hell? Can you tell me that?\"\"A pit full of fire.\"\"And should you like to fall into that pit, and to be burning there for ever?\"\"No, sir.\"\"What must you do to avoid it?\"I deliberated a moment: my answer, when it did come was objectionable: \"I must keep in good health and not die.",
	},
	{
		Author: "Laurie Halse AndersonWintergirls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376224335p2/10003.jpg",
		Content: "In one aspect, yes, I believe in ghosts, but we create them. We haunt ourselves.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Education: the path from cocky ignorance to miserable uncertainty.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "As a woman I have no country. As a woman I want no country. As a woman, my country is the whole world.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "This is how you do it: you sit down at the keyboard and you put one word after another until its done. It's that easy, and that hard.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Ignore those that make you fearful and sad, that degrade you back towards disease and death.",
	},
	{
		Author: "Donna TarttThe Secret History",
		AuthorAvatar: "https://images.gr-assets.com/authors/1409871301p2/8719.jpg",
		Content: "Beauty is terror. Whatever we call beautiful, we quiver before it.",
	},
	{
		Author: "Washington Irving",
		AuthorAvatar: "https://images.gr-assets.com/authors/1218187394p2/28525.jpg",
		Content: "A mother is the truest friend we have, when trials heavy and sudden fall upon us; when adversity takes the place of prosperity; when friends desert us; when trouble thickens around us, still will she cling to us, and endeavor by her kind precepts and counsels to dissipate the clouds of darkness, and cause peace to return to our hearts.",
	},
	{
		Author: "Nora RobertsVision in White",
		AuthorAvatar: "https://images.gr-assets.com/authors/1505847251p2/625.jpg",
		Content: "Some things in life are out of your control. You can make it a party or a tragedy.",
	},
	{
		Author: "Lauren Conrad",
		AuthorAvatar: "https://images.gr-assets.com/authors/1276283295p2/1980985.jpg",
		Content: "Don't cry over someone who wouldn't cry over you.",
	},
	{
		Author: "Katherine Anne Porter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1245828857p2/74572.jpg",
		Content: "The past is never where you think you left it.",
	},
	{
		Author: "George OrwellAnimal Farm",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "The creatures outside looked from pig to man, and from man to pig, and from pig to man again; but already it was impossible to say which was which.",
	},
	{
		Author: "W.B. Yeats",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440689155p2/29963.jpg",
		Content: "The world is full of magic things, patiently waiting for our senses to grow sharper.",
	},
	{
		Author: "Miguel Ruiz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1428604894p2/4402.jpg",
		Content: "1. Be Impeccable With Your WordSpeak with integrity. Say only what you mean. Avoid using the word to speak against yourself or to gossip about others. Use the power of your word in the direction of truth and love. 2. Don't Take Anything PersonallyNothing others do is because of you. What others say and do is a projection of their own reality, their own dream. When you are immune to the opinions and actions of others, you won't be the victim of needless suffering. 3. Don't Make AssumptionsFind the courage to ask questions and to express what you really want. Communicate with others as clearly as you can to avoid misunderstandings, sadness and drama. With just this one agreement, you can completely transform your life. 4. Always Do Your BestYour best is going to change from moment to moment; it will be different when you are healthy as opposed to sick. Under any circumstance, simply do your best, and you will avoid self-judgment, self-abuse and regret.",
	},
	{
		Author: "Amy TanThe Hundred Secret Senses",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437607346p2/5246.jpg",
		Content: "We dream to give ourselves hope. To stop dreaming - well, that’s like saying you can never change your fate.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "You deal with mythological stuff for a few years, you learn that paradises are usually places where you get killed.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Sex is part of nature. I go along with nature.",
	},
	{
		Author: "Mitch AlbomTuesdays with Morrie",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "So many people walk around with a meaningless life. They seem half-asleep, even when they're busy doing things they think are important. This is because they're chasing the wrong things. The way you get meaning into your life is to devote yourself to loving others, devote yourself to your community around you, and devote yourself to creating something that gives you purpose and meaning.",
	},
	{
		Author: "Shannon L. Alder",
		AuthorAvatar: "https://images.gr-assets.com/authors/1446155656p2/1391130.jpg",
		Content: "Dignity/ˈdignitē/ noun 1. The moment you realize that the person you cared for has nothing intellectually or spiritually to offer you, but a headache. 2. The moment you realize God had greater plans for you that don’t involve crying at night or sad Pinterest quotes.3. The moment you stop comparing yourself to others because it undermines your worth, education and your parent’s wisdom.4. The moment you live your dreams, not because of what it will prove or get you, but because that is all you want to do. People’s opinions don’t matter. 5. The moment you realize that no one is your enemy, except yourself.6. The moment you realize that you can have everything you want in life. However, it takes timing, the right heart, the right actions, the right passion and a willingness to risk it all. If it is not yours, it is because you really didn’t want it, need it or God prevented it. 7. The moment you realize the ghost of your ancestors stood between you and the person you loved. They really don't want you mucking up the family line with someone that acts anything less than honorable.8. The moment you realize that happiness was never about getting a person. They are only a helpmate towards achieving your life mission.9. The moment you believe that love is not about losing or winning. It is just a few moments in time, followed by an eternity of situations to grow from.10. The moment you realize that you were always the right person. Only ignorant people walk away from greatness.",
	},
	{
		Author: "Harlan Ellison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1377708311p2/7415.jpg",
		Content: "You are not entitled to your opinion. You are entitled to your informed opinion. No one is entitled to be ignorant.",
	},
	{
		Author: "Honoré de Balzac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206567834p2/228089.jpg",
		Content: "All happiness depends on courage and work.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "If death meant just leaving the stage long enough to change costume and come back as a new character...Would you slow down? Or speed up?",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "The future depends on what you do today.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "Wanting to be someone else is a waste of the person you are.",
	},
	{
		Author: "Art SpiegelmanMaus I: A Survivor's Tale: My Father Bleeds History",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206557373p2/5117.jpg",
		Content: "To die, it's easy. But you have to struggle for life.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1504718173p2/122049.jpg",
		Content: "I like it when somebody gets excited about something. It's nice.",
	},
	{
		Author: "François de La RochefoucauldMaxims",
		AuthorAvatar: "https://images.gr-assets.com/authors/1410969118p2/7428903.jpg",
		Content: "Absence diminishes small loves and increases great ones, as the wind blows out the candle and fans the bonfire.",
	},
	{
		Author: "Janet FitchWhite Oleander",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202586900p2/3540.jpg",
		Content: "\u2028Loneliness is the human condition. Cultivate it. The way it tunnels into you allows your soul room to grow. Never expect to outgrow loneliness. Never hope to find people who will understand you, someone to fill that space. An intelligent, sensitive person is the exception, the very great exception. If you expect to find people who will understand you, you will grow murderous with disappointment. The best you'll ever do is to understand yourself, know what it is that you want, and not let the cattle stand in your way.",
	},
	{
		Author: "Jennifer EganA Visit from the Goon Squad",
		AuthorAvatar: "https://images.gr-assets.com/authors/1231143470p2/49625.jpg",
		Content: "I'm always happy,\" Sasha said. \"Sometimes I just forget.",
	},
	{
		Author: "Margaret Mitchell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1185481392p2/11081.jpg",
		Content: "With enough courage, you can do without a reputation.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The best way to cheer yourself is to try to cheer someone else up.",
	},
	{
		Author: "Zadie Smith",
		AuthorAvatar: "https://images.gr-assets.com/authors/1478188567p2/2522.jpg",
		Content: "The past is always tense, the future perfect.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "In individuals, insanity is rare; but in groups, parties, nations and epochs, it is the rule.",
	},
	{
		Author: "Frida Kahlo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217453723p2/52760.jpg",
		Content: "At the end of the day, we can endure much more than we think we can.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "I do not think, sir, you have any right to command me, merely because you are older than I, or because you have seen more of the world than I have; your claim to superiority depends on the use you have made of your time and experience.",
	},
	{
		Author: "Jack Kerouac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "The only truth is music.",
	},
	{
		Author: "Philip José Farmer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234714074p2/10089.jpg",
		Content: "Imagination is like a muscle. I found out that the more I wrote, the bigger it got.",
	},
	{
		Author: "Fannie FlaggFried Green Tomatoes at the Whistle Stop Cafe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1328803839p2/6125.jpg",
		Content: "I wonder how many people don't get the one they want, but end up with the one they're supposed to be with.",
	},
	{
		Author: "Gore Vidal",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400912972p2/5657.jpg",
		Content: "The unfed mind devours itself.",
	},
	{
		Author: "Ernest Hemingway",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "All good books are alike in that they are truer than if they had really happened and after you are finished reading one you will feel that all that happened to you and afterwards it all belongs to you: the good and the bad, the ecstasy, the remorse and sorrow, the people and the places and how the weather was. If you can get so that you can give that to people, then you are a writer.",
	},
	{
		Author: "Paulo CoelhoVeronika Decides to Die",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "The two hardest tests on the spiritual road are the patience to wait for the right moment and the courage not to be disappointed with what we encounter.",
	},
	{
		Author: "Elizabeth WurtzelProzac Nation",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209360163p2/4370.jpg",
		Content: "That's the thing about depression: A human being can survive almost anything, as long as she sees the end in sight. But depression is so insidious, and it compounds daily, that it's impossible to ever see the end.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I desire the things that will destroy me in the end.",
	},
	{
		Author: "Jonathan Kellerman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1320951950p2/43626.jpg",
		Content: "Life is like a prism. What you see depends on how you turn the glass.",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "We do not believe in ourselves until someone reveals that deep inside us something is valuable, worth listening to, worthy of our trust, sacred to our touch. Once we believe in ourselves we can risk curiosity, wonder, spontaneous delight or any experience that reveals the human spirit.",
	},
	{
		Author: "Nelson Mandela",
		AuthorAvatar: "https://images.gr-assets.com/authors/1308928296p2/367338.jpg",
		Content: "I learned that courage was not the absence of fear, but the triumph over it. The brave man is not he who does not feel afraid, but he who conquers that fear.",
	},
	{
		Author: "Oprah Winfrey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1354837955p2/3518.jpg",
		Content: "Be thankful for what you have; you'll end up having more. If you concentrate on what you don't have, you will never, ever have enough",
	},
	{
		Author: "John W. Campbell Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1333994237p2/5410853.jpg",
		Content: "History does not always repeat itself. Sometimes it just yells, 'Can't you remember anything I told you?' and lets fly with a club.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Well, she's not responding to my advances,\" he observed more brightly than he felt, \"so she must be dead.\" \"Or she's a woman of good taste and sense.",
	},
	{
		Author: "Mitch AlbomFor One More Day",
		AuthorAvatar: "https://images.gr-assets.com/authors/1368640552p2/2331.jpg",
		Content: "Have you ever lost someone you love and wanted one more conversation, one more chance to make up for the time when you thought they would be here forever? If so, then you know you can go your whole life collecting days, and none will outweigh the one you wish you had back.",
	},
	{
		Author: "W.E.B. Du Bois",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211293877p2/10710.jpg",
		Content: "Children learn more from what you are than what you teach.",
	},
	{
		Author: "Jack KerouacOn the Road",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "Nothing behind me, everything ahead of me, as is ever so on the road.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "I suffer from girlnextdooritis where the guy is friends with you and that's it.",
	},
	{
		Author: "Anne Frank",
		AuthorAvatar: "https://images.gr-assets.com/authors/1343271406p2/3720.jpg",
		Content: "Parents can only give good advice or put them on the right paths, but the final forming of a person's character lies in their own hands.",
	},
	{
		Author: "Anaïs NinIncest: From a Journal of Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Reality doesn't impress me. I only believe in intoxication, in ecstasy, and when ordinary life shackles me, I escape, one way or another. No more walls.",
	},
	{
		Author: "L.M. MontgomeryAnne of Green Gables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188896723p2/5350.jpg",
		Content: "I'm so glad I live in a world where there are Octobers.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "You are okay?\" he asked. \"Not eaten by monsters?\"\"Not even a little bit.\" I showed him that I still had both arms and both legs, and Tyson clapped happily.\"Yay!\" he said. \"Now we can eat peanut butter sandwiches and ride fish ponies! We can fight monsters and see Annabeth and make things go BOOM!\"I hoped he didn't mean all at the same time, but I told him absolutely, we'd have a lot of fun this summer.",
	},
	{
		Author: "Rainbow RowellEleanor \u0026 Park",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342324527p2/4208569.jpg",
		Content: "I want everyone to meet you. You're my favorite person of all time.",
	},
	{
		Author: "Tahereh MafiShatter Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "The moon is a loyal companion.It never leaves. It’s always there, watching, steadfast, knowing us in our light and dark moments, changing forever just as we do. Every day it’s a different version of itself. Sometimes weak and wan, sometimes strong and full of light. The moon understands what it means to be human.Uncertain. Alone. Cratered by imperfections.",
	},
	{
		Author: "F. Scott Fitzgerald",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "Whenever you feel like criticizing any one...just remember that all the people in this world haven't had the advantages that you've had.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "Name the greatest of all inventors. Accident.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Sorcerer's Stone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "So light a fire!\" Harry choked. \"Yes...of course...but there's no wood!\" ... \"HAVE YOU GONE MAD!\" Ron bellowed. \"ARE YOU A WITCH OR NOT!",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "A man is but the product of his thoughts. What he thinks, he becomes.",
	},
	{
		Author: "Becca FitzpatrickCrescendo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "Being with you never felt wrong. It's the one thing I did right. You're the one thing I did right.",
	},
	{
		Author: "Rick RiordanThe Son of Neptune",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Um...is that thing tame?\" Frank said.The horse whinnied angrily.\"I don't think so,\" Percy guessed. \"He just said, 'I will trample you to death, silly Chinese Canadian baby man'.",
	},
	{
		Author: "Seamus HeaneyStepping Stones: Interviews with Seamus Heaney",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200407647p2/29574.jpg",
		Content: "If you have the words, there's always a chance that you'll find the way.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Will rolled up his sleeves. \"We'll probably have to knock down the door--\" \"Or,\" said Jem, reaching out and giving the knob a twist, \"not.\" The door swung open onto a rectangle of darkness. \"Now, that's simply laziness,\" said Will.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "They're certainly entitled to think that, and they're entitled to full respect for their opinions... but before I can live with other folks I've got to live with myself. The one thing that doesn't abide by majority rule is a person's conscience.",
	},
	{
		Author: "Lemony SnicketThe Bad Beginning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "...you know that a good, long session of weeping can often make you feel better, even if your circumstances have not changed one bit.",
	},
	{
		Author: "Rick RiordanThe Last Olympian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "It's hard to enjoy practical jokes when your whole life feels like one.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "But I believe in true love, you know? I don't believe that everybody gets to keep their eyes or not get sick or whatever, but everybody should have true love, and it should last at least as long as your life does.",
	},
	{
		Author: "Richelle MeadBlood Promise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Stop fighting me!\" he said, trying to pull on the arm he held.He was in a precarious position himself, straddling the rail as he tried to lean over far enough to get me and actually hold onto me.“Let go of me!",
	},
	{
		Author: "Rick Riordan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Can you surf really well, then?\"I looked at Grover, who was trying hard not to laugh.\"Jeez, Nico,\" I said. \"I've never really tried.\"He went on asking questions. Did I fight a lot with Thalia, since she was a daughter of Zeus? (I didn't answer that one.) If Annabeth's mother was Athena, the goddess of wisdom, then why didn't Annabeth know better than to fall off a cliff? (I tried not to strangle Nico for asking that one.) Was Annabeth my girlfriend? (At this point, I was ready to stick the kid in a meat-flavored sack and throw him to the wolves.)",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I must be a mermaid, Rango. I have no fear of depths and a great fear of shallow living.",
	},
	{
		Author: "Margaret AtwoodThe Handmaid's Tale",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "Don't let the bastards grind you down.",
	},
	{
		Author: "Roy T. BennettThe Light in the Heart",
		AuthorAvatar: "",
		Content: "If you want to be happy, do not dwell in the past, do not worry about the future, focus on living fully in the present.",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "If you want to keep a secret, you must also hide it from yourself.",
	},
	{
		Author: "Noam Chomsky",
		AuthorAvatar: "https://images.gr-assets.com/authors/1485473965p2/2476.jpg",
		Content: "We shouldn't be looking for heroes, we should be looking for good ideas.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "You never know what life is like, until you have lived it.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Fashion is a form of ugliness so intolerable that we have to alter it every six months.",
	},
	{
		Author: "Julie KagawaThe Iron King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1257816454p2/2995873.jpg",
		Content: "Oh, we're playing nice now? Shall we have tea first? Brew up a nice pot of kiss-my-ass?",
	},
	{
		Author: "عباس محمود العقاد",
		AuthorAvatar: "https://images.gr-assets.com/authors/1509897931p2/1409619.jpg",
		Content: "ليس هناك كتابا أقرأه و لا أستفيد منه شيئا جديدا ، فحتى الكتاب التافه أستفيد من قراءته ، أني تعلمت شيئا جديدا هو ما هي التفاهة ؟ و كيف يكتب الكتاب التافهون ؟ و فيم يفكرون ؟",
	},
	{
		Author: "Walt Disney Company",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289112683p2/3510823.jpg",
		Content: "All our dreams can come true, if we have the courage to pursue them.",
	},
	{
		Author: "Elizabeth Barrett Browning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206559375p2/67032.jpg",
		Content: "How do I love thee? Let me count the ways.I love thee to the depth and breadth and heightMy soul can reach",
	},
	{
		Author: "Doris LessingThe Golden Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1457477725p2/7728.jpg",
		Content: "What's terrible is to pretend that second-rate is first-rate. To pretend that you don't need love when you do; or you like your work when you know quite well you're capable of better.",
	},
	{
		Author: "Nadezhda MandelstamHope Against Hope",
		AuthorAvatar: "https://images.gr-assets.com/authors/1278457514p2/61637.jpg",
		Content: "I decided it is better to scream. Silence is the real crime against humanity.",
	},
	{
		Author: "David LevithanEvery Day",
		AuthorAvatar: "https://images.gr-assets.com/authors/1426529210p2/11664.jpg",
		Content: "If there's one thing I've learned, it's this: We all want everything to be okay. We don't even wish so much for fantastic or marvelous or outstanding. We will happily settle for okay, because most of the time, okay is enough.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Behind every exquisite thing that existed, there was something tragic.",
	},
	{
		Author: "Charles BukowskiThe People Look Like Flowers at Last",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "A love like that was a serious illness, an illness from which you never entirely recover.",
	},
	{
		Author: "Tennessee WilliamsA Streetcar Named Desire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206504901p2/7751.jpg",
		Content: "What is straight? A line can be straight, or a street, but the human heart, oh, no, it's curved like a road through mountains.",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "Books serve to show a man that those original thoughts of his aren't very new after all.",
	},
	{
		Author: "J.R.R. TolkienThe Lord of the Rings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "It's the job that's never started as takes longest to finish.",
	},
	{
		Author: "Terry Pratchett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "If you have enough book space, I don't want to talk to you.",
	},
	{
		Author: "Zig ZiglarRaising Positive Kids in a Negative World",
		AuthorAvatar: "https://images.gr-assets.com/authors/1236462809p2/50316.jpg",
		Content: "Of course motivation is not permanent. But then, neither is bathing; but it is something you should do on a regular basis.",
	},
	{
		Author: "Rose Wilder Lane",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195976419p2/5456.jpg",
		Content: "Happiness is something that comes into our lives through doors we don't even remember leaving open.",
	},
	{
		Author: "Sherman AlexieThe Absolutely True Diary of a Part-Time Indian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1333515890p2/4174.jpg",
		Content: "If you let people into your life a little bit, they can be pretty damn amazing.",
	},
	{
		Author: "Rachel CaineGlass Houses",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373144795p2/15292.jpg",
		Content: "Damn, Claire. Warn a guy before you do a face-plant on the floor next time. I could have looked all heroic and caught you or something -Shane",
	},
	{
		Author: "Elizabeth GilbertCommitted: A Skeptic Makes Peace with Marriage",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "To be fully seen by somebody, then, and be loved anyhow - this is a human offering that can border on miraculous.",
	},
	{
		Author: "David Foster WallaceInfinite Jest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466019433p2/4339.jpg",
		Content: "The truth will set you free. But not until it is finished with you.",
	},
	{
		Author: "Selma Lagerlöf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1237932710p2/38266.jpg",
		Content: "Nothing on earth can make up for the loss of one who has loved you.",
	},
	{
		Author: "Nicole KraussThe History of Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1285130428p2/2633.jpg",
		Content: "there are two types of people in the world: those who prefer to be sad among others, and those who prefer to be sad alone.",
	},
	{
		Author: "Frank Zappa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1315160559p2/22302.jpg",
		Content: "If you want to get laid, go to college. If you want an education, go to the library.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "Many people will walk in and out of your life, but only true friends will leave footprints in your heart",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "All I have seen teaches me to trust the Creator for all I have not seen.",
	},
	{
		Author: "Flannery O'Connor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1469878767p2/22694.jpg",
		Content: "She looked at nice young men as if she could smell their stupidity.",
	},
	{
		Author: "Neil Gaiman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "I hope that in this year to come, you make mistakes.Because if you are making mistakes, then you are making new things, trying new things, learning, living, pushing yourself, changing yourself, changing your world. You're doing things you've never done before, and more importantly, you're Doing Something.So that's my wish for you, and all of us, and my wish for myself. Make New Mistakes. Make glorious, amazing mistakes. Make mistakes nobody's ever made before. Don't freeze, don't stop, don't worry that it isn't good enough, or it isn't perfect, whatever it is: art, or love, or work or family or life.Whatever it is you're scared of doing, Do it.Make your mistakes, next year and forever.",
	},
	{
		Author: "Jon KrakauerInto the Wild",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430856379p2/1235.jpg",
		Content: "make a radical change in your lifestyle and begin to boldly do things which you may previously never have thought of doing, or been too hesitant to attempt. So many people live within unhappy circumstances and yet will not take the initiative to change their situation because they are conditioned to a life of security, conformity, and conservation, all of which may appear to give one peace of mind, but in reality nothing is more damaging to the adventurous spirit within a man than a secure future. The very basic core of a man's living spirit is his passion for adventure. The joy of life comes from our encounters with new experiences, and hence there is no greater joy than to have an endlessly changing horizon, for each day to have a new and different sun. If you want to get more out of life, you must lose your inclination for monotonous security and adopt a helter-skelter style of life that will at first appear to you to be crazy. But once you become accustomed to such a life you will see its full meaning and its incredible beauty.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "When you trip over love, it is easy to get up. But when you fall in love, it is impossible to stand again.",
	},
	{
		Author: "Joseph ConradAn Outcast of the Islands",
		AuthorAvatar: "https://images.gr-assets.com/authors/1403814208p2/3345.jpg",
		Content: "It's only those who do nothing that make no mistakes, I suppose.",
	},
	{
		Author: "Ellen Bass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1452839679p2/73434.jpg",
		Content: "to love life, to love it evenwhen you have no stomach for itand everything you've held dearcrumbles like burnt paper in your hands,your throat filled with the silt of it.When grief sits with you, its tropical heatthickening the air, heavy as watermore fit for gills than lungs;when grief weights you like your own fleshonly more of it, an obesity of grief,you think, How can a body withstand this?Then you hold life like a facebetween your palms, a plain face,no charming smile, no violet eyes,and you say, yes, I will take youI will love you, again.",
	},
	{
		Author: "Colleen HooverSlammed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1464032240p2/5430144.jpg",
		Content: "There are three questions every woman should be able to answer yes to before they commit to a man. If you answer no to any of the three questions, run like hell.\"[...]\"Does he treat you with respect at all times? That's the first question. The second question is, if he is the exact same person twenty years from now that he is today, would you still want to marry him? And finally, does he inspire to be a better person? You find someone you can answer yes to all three, then you've found a good man.",
	},
	{
		Author: "Nathaniel HawthorneThe Scarlet Letter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1291476587p2/7799.jpg",
		Content: "We dream in our waking moments, and walk in our sleep.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "The things you used to own, now they own you.",
	},
	{
		Author: "Paulo CoelhoThe Fifth Mountain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "There are moments when troubles enter our lives and we can do nothing to avoid them.But they are there for a reason. Only when we have overcome them will we understand why they were there.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "Atticus, he was real nice.\"\"Most people are, Scout, when you finally see them.",
	},
	{
		Author: "Laini TaylorDaughter of Smoke \u0026 Bone",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224474224p2/324620.jpg",
		Content: "Have you ever asked yourself, do monsters make war, or does war make monsters?",
	},
	{
		Author: "Albert DietrichArmy GI, Pacifist Co: The World War II Letters of Frank Dietrich and Albert Dietrich",
		AuthorAvatar: "",
		Content: "There are perhaps many causes worth dying for, but to me, certainly, there are none worth killing for.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "A man's face is his autobiography. A woman's face is her work of fiction.",
	},
	{
		Author: "Elizabeth Barrett Browning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206559375p2/67032.jpg",
		Content: "You're something between a dream and a miracle.",
	},
	{
		Author: "Lao TzuTao Te Ching",
		AuthorAvatar: "https://images.gr-assets.com/authors/1435903703p2/2622245.jpg",
		Content: "When you are content to be simply yourself and don't compare or compete, everyone will respect you.",
	},
	{
		Author: "Stephanie  LennoxI Don't Remember You",
		AuthorAvatar: "",
		Content: "I’ve been fighting to be who I am all my life. What’s the point of being who I am, if I can’t have the person who was worth all the fighting for?",
	},
	{
		Author: "Sylvia PlathThe Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "Can you understand? Someone, somewhere, can you understand me a little, love me a little? For all my despair, for all my ideals, for all that - I love life. But it is hard, and I have so much - so very much to learn.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Sometimes, the best way to help someone is just to be near them.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Seventeen, eh!\" said Hagrid as he accepted a bucket-sized glass of wine from Fred.\"Six years to the day we met, Harry, d’yeh remember it?\"\"Vaguely,\" said Harry, grinning up at him. \"Didn’t you smash down the front door, give Dudley a pig’s tail, and tell me I was a wizard?\"\"I forge’ the details,\" Hagrid chortled.",
	},
	{
		Author: "Aristotle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390143800p2/2192.jpg",
		Content: "Hope is a waking dream.",
	},
	{
		Author: "Carlos Castaneda",
		AuthorAvatar: "https://images.gr-assets.com/authors/1223890018p2/8088.jpg",
		Content: "You have everything needed for the extravagant journey that is your life.",
	},
	{
		Author: "Nicholas SparksMessage in a Bottle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Without you in my arms, I feel an emptiness in my soul. I find myself searching the crowds for your face - I know it's an impossibility, but I cannot help myself.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "This is a new year. A new beginning. And things will change.",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "Why do beautiful songs make you sad?' 'Because they aren't true.' 'Never?' 'Nothing is beautiful and true.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Demon pox, oh demon poxJust how is it acquired?One must go down to the bad part of townUntil one is very tired.Demon pox, oh demon pox, I had it all along—Not the pox, you foolish blocks,I mean this very song—For I was right, and you were wrong!\"\"Will!\" Charlotte shouted over the noise, \"Have you LOST YOUR MIND? CEASE THAT INFERNAL RACKET! Jem—\" Jem, rising to his feet, clapped his hands over Will's mouth. \"Do you promise to be quiet?\" he hissed into his friend's ear.Will nodded, blue eyes blazing. Tessa was staring at him in amazement; they all were. She had seen Will many things—amused, bitter, condescending, angry, pitying—but never giddy before.Jem let him go. \"All right, then.\"Will slid to the floor, his back against the armchair, and threw up his arms. \"A demon pox on all your houses!\" he announced, and yawned.\"Oh, God, weeks of pox jokes,\" said Jem. \"We're in for it now.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "In your light I learn how to love. In your beauty, how to make poems. You dance inside my chest where no-one sees you, but sometimes I do, and that sight becomes this art.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "You believe in a book that has talking animals, wizards, witches, demons, sticks turning into snakes, burning bushes, food falling from the sky, people walking on water, and all sorts of magical, absurd and primitive stories, and you say that we are the ones that need help?",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "You alone are enough. You have nothing to prove to anybody.",
	},
	{
		Author: "Sarah DessenThis Lullaby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Holding people away from you, and denying yourself love, that doesn't make you strong. if anything, it makes you weaker. Because you're doing it out of fear.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "He died not for men, but for each man. If each man had been the only man made, He would have done no less.",
	},
	{
		Author: "Walter Tevis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1296773197p2/4448408.jpg",
		Content: "I feel free and strong. If I were not a reader of books I could not feel this way.",
	},
	{
		Author: "David NichollsOne Day",
		AuthorAvatar: "https://images.gr-assets.com/authors/1283277096p2/46118.jpg",
		Content: "This is where it all begins. Everything starts here, today.",
	},
	{
		Author: "Erin MorgensternThe Night Circus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316440077p2/4370565.jpg",
		Content: "The finest of pleasures are always the unexpected ones.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "What is right is not always popular and what is popular is not always right.",
	},
	{
		Author: "Jonathan AmesMy Less Than Secret Life: A Diary, Fiction, Essays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1221071651p2/2285.jpg",
		Content: "I live for coincidences. They briefly give to me the illusion or the hope that there's a pattern to my life, and if there's a pattern, then maybe I'm moving toward some kind of destiny where it's all explained.",
	},
	{
		Author: "Roy T. BennettThe Light in the Heart",
		AuthorAvatar: "",
		Content: "We are all different. Don’t judge, understand instead.",
	},
	{
		Author: "Agatha ChristieMurder on the Orient Express",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321738793p2/123715.jpg",
		Content: "The impossible could not have happened, therefore the impossible must be possible in spite of appearances.",
	},
	{
		Author: "Kiera CassThe Selection",
		AuthorAvatar: "https://images.gr-assets.com/authors/1318605410p2/2987125.jpg",
		Content: "True love is usually the most inconvenient kind.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "I'm afraid that sometimes you'll play lonely games too. Games you can't win 'cause you'll play against you.",
	},
	{
		Author: "H.L. Mencken",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399814051p2/7805.jpg",
		Content: "I know some who are constantly drunk on books as other men are drunk on whiskey.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "To give pleasure to a single heart by a single act is better than a thousand heads bowing in prayer.",
	},
	{
		Author: "Leonard CohenSelected Poems, 1956-1968",
		AuthorAvatar: "https://images.gr-assets.com/authors/1458346890p2/52060.jpg",
		Content: "There is a crack in everything.That's how the light gets in.",
	},
	{
		Author: "Kurt VonnegutA Man Without a Country",
		AuthorAvatar: "https://images.gr-assets.com/authors/1433582280p2/2778055.jpg",
		Content: "And I urge you to please notice when you are happy, and exclaim or murmur or think at some point, 'If this isn't nice, I don't know what is.",
	},
	{
		Author: "Ellen DeGeneres",
		AuthorAvatar: "https://images.gr-assets.com/authors/1328758901p2/40648.jpg",
		Content: "My grandmother started walking five miles a day when she was sixty. She's ninety-seven now, and we don't know where the heck she is.",
	},
	{
		Author: "Steve Jobs",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322675535p2/5255891.jpg",
		Content: "Being the richest man in the cemetery doesn't matter to me. Going to bed at night saying we've done something wonderful... that's what matters to me.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Whatever you are physically...male or female, strong or weak, ill or healthy--all those things matter less than what your heart contains. If you have the soul of a warrior, you are a warrior. All those other things, they are the glass that contains the lamp, but you are the light inside.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "You're a guardian angel now.\" I was still too much in awe to wrap my mind around it, but at the same time I felt amazement, curiosity...happiness.\"I'm your guardian angel,\" he said.\"I get my very own guardian angel? What, exactly, is your job description?\"\"Guard your body.\" His smile tipped higher. \"I take my job seriously, which means I'm going to need to get acquainted with the subject matter on a personal level.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Moths,\" repeats Will. \"You're afraid of moths?\" \"Not just a cloud of moths,\" she says, \"like...a swarm of them. Everywhere. All those wings and legs and...\" She shudders and shakes her head.\"Terrifying,\" Will says with mock seriousness. \"That's my girl. Tough as cotton balls.\"\"Oh, Shut up.",
	},
	{
		Author: "Roald DahlThe Twits",
		AuthorAvatar: "https://images.gr-assets.com/authors/1311554908p2/4273.jpg",
		Content: "A person who has good thoughts cannot ever be ugly. You can have a wonky nose and a crooked mouth and a double chin and stick-out teeth, but if you have good thoughts they will shine out of your face like sunbeams and you will always look lovely.",
	},
	{
		Author: "Gertrude Stein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422989334p2/9325.jpg",
		Content: "We are always the same age inside. ",
	},
	{
		Author: "Zora Neale Hurston",
		AuthorAvatar: "https://images.gr-assets.com/authors/1194472605p2/15151.jpg",
		Content: "Love makes your soul crawl out from its hiding place.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "Sometimes the Bible in the hand of one man is worse than a whisky bottle in the hand of (another)... There are just some kind of men who - who're so busy worrying about the next world they've never learned to live in this one, and you can look down the street and see the results.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I just want to be wonderful.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "Hang on, did you just call me Angel?\" I asked.\"If I did?\"\"I don't like it.\"He grinned. \"It stays, Angel.",
	},
	{
		Author: "Philip Pullman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1396622492p2/3618.jpg",
		Content: "We don’t need a list of rights and wrongs, tables of dos and don’ts: we need books, time, and silence. Thou shalt not is soon forgotten, but Once upon a time lasts forever.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I am going to take this bucket of water and pour it on the flames of hell, and then I am going to use this torch to burn down the gates of paradise so that people will not love God for want of heaven or fear of hell, but because He is God.",
	},
	{
		Author: "T.S. EliotThe Use of Poetry and the Use of Criticism",
		AuthorAvatar: "https://images.gr-assets.com/authors/1507887310p2/18540.jpg",
		Content: "To do the useful thing, to say the courageous thing, to contemplate the beautiful thing: that is enough for one man's life.",
	},
	{
		Author: "J.R.R. TolkienThe Fellowship of the Ring",
		AuthorAvatar: "https://images.gr-assets.com/authors/1434625177p2/656983.jpg",
		Content: "May it be a light to you in dark places, when all other lights go out.",
	},
	{
		Author: "Hafiz",
		AuthorAvatar: "",
		Content: "Even After All this timeThe Sun never says to the Earth,\"You owe me.\"LookWhat happensWith a love like that,It lights the whole sky.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "The homemaker has the ultimate career. All other careers exist for one purpose only - and that is to support the ultimate career. ",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "There are much worse games to play.",
	},
	{
		Author: "John GreenWill Grayson, Will Grayson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "When things break, it's not the actual breaking that prevents them from getting back together again. It's because a little piece gets lost - the two remaining ends couldn't fit together even if they wanted to. The whole shape has changed.",
	},
	{
		Author: "Daphne du Maurier",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422444467p2/2001717.jpg",
		Content: "Women want love to be a novel. Men, a short story.",
	},
	{
		Author: "Haruki Murakami",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "But I didn't understand then. That I could hurt somebody so badly she would never recover. That a person can, just by living, damage another human being beyond repair.",
	},
	{
		Author: "Audre LordeSister Outsider: Essays and Speeches",
		AuthorAvatar: "https://images.gr-assets.com/authors/1507572634p2/18486.jpg",
		Content: "Your silence will not protect you.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "The man who passes the sentence should swing the sword. If you would take a man's life, you owe it to him to look into his eyes and hear his final words. And if you cannot bear to do that, then perhaps the man does not deserve to die.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "District 12: Where you can starve to death in safety.",
	},
	{
		Author: "Steve MaraboliLife, the Truth, and Being Free",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "I want to be in a relationship where you telling me you love me is just a ceremonious validation of what you already show me.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "When you go into the ER, one of the first things they ask you to do is rate your pain on a scale of one to ten, and from there they decide which drugs to use and how quickly to use them. I'd been asked this question hundreds of times over the years, and I remember once early on when I couldn't get my breath and it felt like my chest was on fire, flames licking the inside of my ribs fighting for a way to burn out of my body, my parents took me to the ER. nurse asked me about the pain, and I couldn't even speak, so I held up nine fingers.Later, after they'd given me something, the nurse came in and she was kind of stroking my head while she took my blood pressure and said, \"You know how I know you're a fighter? You called a ten a nine.\"But that wasn't quite right. I called it a nine because I was saving my ten. And here it was, the great and terrible ten, slamming me again and again as I lay still and alone in my bed staring at the ceiling, the waves tossing me against the rocks then pulling me back out to sea so they could launch me again into the jagged face of the cliff, leaving me floating faceup on the water, undrowned.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "Good sex is like good bridge. If you don't have a good partner, you'd better have a good hand.",
	},
	{
		Author: "Gertrude Stein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1422989334p2/9325.jpg",
		Content: "One must dare to be happy. ",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Sometimes you lose a battle. But mischief always wins the war",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Freedom lies in being bold.",
	},
	{
		Author: "Martin Luther King Jr.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518558p2/23924.jpg",
		Content: "Intelligence plus character-that is the goal of true education.",
	},
	{
		Author: "Alyson NoelEvermore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1360452950p2/200317.jpg",
		Content: "I guess by now I should know enough about loss to realize that you never really stop missing someone-you just learn to live around the huge gaping hole of their absence.",
	},
	{
		Author: "Edward Abbey",
		AuthorAvatar: "https://images.gr-assets.com/authors/1516467601p2/37218.jpg",
		Content: "May your trails be crooked, winding, lonesome, dangerous, leading to the most amazing view.",
	},
	{
		Author: "Leo TolstoyEssays, Letters and Miscellanies",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "If, then, I were asked for the most important advice I could give, that which I considered to be the most useful to the men of our century, I should simply say: in the name of God, stop a moment, cease your work, look around you.",
	},
	{
		Author: "Henry Thomas Buckle",
		AuthorAvatar: "",
		Content: "Great minds discuss ideas. Average minds discuss events. Small minds discuss people.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "I never forget a face, but in your case I'll be glad to make an exception.",
	},
	{
		Author: "Thomas Pynchon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1465361157p2/235.jpg",
		Content: "Every weirdo in the world is on my wavelength.",
	},
	{
		Author: "Rachel Carson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397487410p2/15332.jpg",
		Content: "The more clearly we can focus our attention on the wonders and realities of the universe about us, the less taste we shall have for destruction.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "They say time heals all wounds, but that presumes the source of the grief is finite",
	},
	{
		Author: "Umberto EcoThe Island of the Day Before",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455915753p2/1730.jpg",
		Content: "To survive, you must tell stories.",
	},
	{
		Author: "Charles DickensGreat Expectations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "Heaven knows we need never be ashamed of our tears, for they are rain upon the blinding dust of earth, overlying our hard hearts. I was better after I had cried, than before--more sorry, more aware of my own ingratitude, more gentle.",
	},
	{
		Author: "Stephen Fry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400162446p2/10917.jpg",
		Content: "Books are no more threatened by Kindle than stairs by elevators.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "It is easy to forget how full the world is of people, full to bursting, and each of them imaginable and consistently misimagined.",
	},
	{
		Author: "Jane AustenPride and Prejudice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "I must learn to be content with being happier than I deserve.",
	},
	{
		Author: "Douglas AdamsThe Restaurant at the End of the Universe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "The major problem—one of the major problems, for there are several—one of the many major problems with governing people is that of whom you get to do it; or rather of who manages to get people to let them do it to them. To summarize: it is a well-known fact that those people who must want to rule people are, ipso facto, those least suited to do it. To summarize the summary: anyone who is capable of getting themselves made President should on no account be allowed to do the job.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "I restore myself when I'm alone.",
	},
	{
		Author: "Ernest HemingwayThe Old Man and the Sea",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "Every day is a new day. It is better to be lucky. But I would rather be exact. Then when luck comes you are ready.",
	},
	{
		Author: "Victor HugoThe Hunchback of Notre-Dame",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "Love is like a tree: it grows by itself, roots itself deeply in our being and continues to flourish over a heart in ruin. The inexplicable fact is that the blinder it is, the more tenacious it is. It is never stronger than when it is completely unreasonable.",
	},
	{
		Author: "Søren KierkegaardThe Concept of Anxiety: A Simple Psychologically Orienting Deliberation on the Dogmatic Issue of Hereditary Sin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202410387p2/6172.jpg",
		Content: "Anxiety is the dizziness of freedom.",
	},
	{
		Author: "Stephenie MeyerTwilight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "I decided as long as I'm going to hell, I might as well do it thoroughly.",
	},
	{
		Author: "Sarah VowellThe Partly Cloudy Patriot",
		AuthorAvatar: "https://images.gr-assets.com/authors/1297911965p2/2122.jpg",
		Content: "Being a nerd, which is to say going too far and caring too much about a subject, is the best way to make friends I know.",
	},
	{
		Author: "Neil JordanThe Dream of a Beast",
		AuthorAvatar: "https://images.gr-assets.com/authors/1488059098p2/29997.jpg",
		Content: "Is it fair to have given us the memory of what was and the desire of what could be when we must suffer what is?",
	},
	{
		Author: "Ian FlemingYou Only Live Twice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1364532740p2/2565.jpg",
		Content: "You only live twice:Once when you are bornAnd once when you look death in the face",
	},
	{
		Author: "Jarod KintzThis Book is Not for Sale",
		AuthorAvatar: "https://images.gr-assets.com/authors/1460877358p2/4157885.jpg",
		Content: "I find out a lot about myself by sleeping. Dreams, they are who I am when I’m too tired to be me.",
	},
	{
		Author: "Dan BrownThe Da Vinci Code",
		AuthorAvatar: "https://images.gr-assets.com/authors/1399396714p2/630.jpg",
		Content: "Men go to far greater lengths to avoid what they fear than to obtain what they desire.",
	},
	{
		Author: "Judy Garland",
		AuthorAvatar: "https://images.gr-assets.com/authors/1275327407p2/179335.jpg",
		Content: "Always be a first rate version of yourself and not a second rate version of someone else.",
	},
	{
		Author: "Jandy NelsonI'll Give You the Sun",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1441828634p2/1324179.jpg",
		Content: "Maybe some people are just meant to be in the same story.",
	},
	{
		Author: "C.J. Cherryh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244675150p2/989968.jpg",
		Content: "Ignorance killed the cat; curiosity was framed!",
	},
	{
		Author: "Viktor E. FranklMan's Search for Meaning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1218372404p2/2782.jpg",
		Content: "Everything can be taken from a man but one thing: the last of the human freedoms—to choose one’s attitude in any given set of circumstances, to choose one’s own way.",
	},
	{
		Author: "Alice McDermott",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208206092p2/7505.jpg",
		Content: "We are surrounded by story.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "If Peeta and I were both to die, or they thought we were....My fingers fumble with the pouch on my belt, freeing it. Peeta sees it and his hand clamps on my wrist. \"No, I won't let you.\" \"Trust me,\" I whisper. He holds my gaze for a long moment then lets go. I loosen the top of the pouch and pour a few spoonfuls of berries into his palm. Then I fill my own. \"On the count of three?\" Peeta leans down and kisses me once, very gently. \"The count of three,\" he says. We stand, our backs pressed together, our empty hands locked tight. \"Hold them out. I want everyone to see,\" he says. I spread out my fingers, and the dark berries glisten in the sun. I give Peeta's hand one last squeeze as a signal, as a good-bye, and we begin counting. \"One.\" Maybe I'm wrong. \"Two.\" Maybe they don't care if we both die. \"Three!\" It's too late to change my mind. I lift my hand to my mouth taking one last look at the world. The berries have just passed my lips when the trumpets begin to blare. The frantic voice of Claudius Templesmith shouts above them. \"Stop! Stop! Ladies and gentlemen, I am pleased to present the victors of the 74th Hunger Games, Katniss Everdeen and Peeta Mellark! I give you - the tributes of District 12!",
	},
	{
		Author: "Victor HugoLes Misérables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "The power of a glance has been so much abused in love stories, that it has come to be disbelieved in. Few people dare now to say that two beings have fallen in love because they have looked at each other. Yet it is in this way that love begins, and in this way only.",
	},
	{
		Author: "Marcel Proust",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392271688p2/233619.jpg",
		Content: "The real voyage of discovery consists not in seeking new landscapes, but in having new eyes.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Anybody can sympathise with the sufferings of a friend, but it requires a very fine nature to sympathise with a friend's success.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "We need to talk. All of us About what we're going to do now.\"\"I was going to watch Project Runway.",
	},
	{
		Author: "Sinclair Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1205204856p2/7330.jpg",
		Content: "We'd get sick on too many cookies, but ever so much sicker on no cookies at all.",
	},
	{
		Author: "Jean-Paul Sartre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475567078p2/1466.jpg",
		Content: "Do you think that I count the days? There is only one day left, always starting over: it is given to us at dawn and taken away from us at dusk.",
	},
	{
		Author: "Marvin Bell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407068635p2/121818.jpg",
		Content: "Learn the rules, break the rules, make up new rules, break the new rules.",
	},
	{
		Author: "Fred Rogers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208803867p2/32106.jpg",
		Content: "When I say it's you I like, I'm talking about that part of you that knows that life is far more than anything you can ever see or hear or touch. That deep part of you that allows you to stand for those things without which humankind cannot survive. Love that conquers hate, peace that rises triumphant over war, and justice that proves more powerful than greed.",
	},
	{
		Author: "Veronica RothDivergent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363910238p2/4039811.jpg",
		Content: "Can you be a girl for a few seconds?\"\"I'm always a girl\" I frown.\"You know what I mean. Like a silly, annoying girl\"I twirl my hair around my finger. \"Kay.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "We can't be afraid of change. You may feel very secure in the pond that you are in, but if you never venture out of it, you will never know that there is such a thing as an ocean, a sea. Holding onto something that is good for you now, may be the very reason why you don't have something better.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I figured all your classes were stuff like Slaughter 101 and Beheading for Beginners.\"Jace flipped a page. \"Very funny, Fray.",
	},
	{
		Author: "Roberto Bolaño2666",
		AuthorAvatar: "https://images.gr-assets.com/authors/1464727787p2/72039.jpg",
		Content: "Reading is like thinking, like praying, like talking to a friend, like expressing your ideas, like listening to other people's ideas, like listening to music, like looking at the view, like taking a walk on the beach.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Deathly Hallows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "I'm going to keep going until I succeed — or die. Don't think I don't know how this might end. I've known it for years.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "It crosses my mind that Cinna's calm and normal demeanor masks a complete madman.",
	},
	{
		Author: "Rick RiordanThe Mark of Athena",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "We're staying together,\" he promised. \"You're not getting away from me. Never again.",
	},
	{
		Author: "Katherine Mansfield",
		AuthorAvatar: "https://images.gr-assets.com/authors/1216670886p2/45712.jpg",
		Content: "The mind I love must have wild places.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "Let no one ever come to you without leaving better and happier. Be the living expression of God's kindness: kindness in your face, kindness in your eyes, kindness in your smile.",
	},
	{
		Author: "David MitchellCloud Atlas",
		AuthorAvatar: "https://images.gr-assets.com/authors/1409248688p2/6538289.jpg",
		Content: "Our lives are not our own. We are bound to others, past and present, and by each crime and every kindness, we birth our future.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "Usually we walk around constantly believing ourselves. \"I'm okay\" we say. \"I'm alright\". But sometimes the truth arrives on you and you can't get it off. That's when you realize that sometimes it isn't even an answer--it's a question. Even now, I wonder how much of my life is convinced.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "When I'm good, I'm very good, but when I'm bad, I'm better. ",
	},
	{
		Author: "Jeanette WintersonWritten on the Body",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406177070p2/9399.jpg",
		Content: "You’ll get over it…",
	},
	{
		Author: "Sherrilyn KenyonDance with the Devil",
		AuthorAvatar: "https://images.gr-assets.com/authors/1382066830p2/4430.jpg",
		Content: "I don't suffer from my insanity -- I enjoy every minute of it.",
	},
	{
		Author: "Carlos Ruiz ZafónThe Shadow of the Wind",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444199853p2/815.jpg",
		Content: "People tend to complicate their own lives, as if living weren't already complicated enough.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Luxury is not a necessity to me, but beautiful and good things are.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I am both happy and sad at the same time, and I'm still trying to figure out how that could be.",
	},
	{
		Author: "Lois Lowry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1348162077p2/2493.jpg",
		Content: "It is very risky. But each time a child opens a book, he pushes open the gate that separates him from Elsewhere.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "He was the crazy one who had painted himself black and defeated the world.She was the book thief without the words.Trust me, though, the words were on their way, and when they arrived, Liesel would hold them in her hands like the clouds, and she would wring them out like rain.",
	},
	{
		Author: "Paulo CoelhoEleven Minutes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Now that she had nothing to lose, she was free.",
	},
	{
		Author: "Helen KellerThe Open Door",
		AuthorAvatar: "https://images.gr-assets.com/authors/1266096039p2/7275.jpg",
		Content: "Life is either a daring adventure or nothing at all.",
	},
	{
		Author: "Charlie Chaplin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1248334782p2/48136.jpg",
		Content: "Life is a beautiful magnificent thing, even to a jellyfish.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "What draws people to be friends is that they see the same truth. They share it.",
	},
	{
		Author: "Toni MorrisonBeloved",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494211316p2/3534.jpg",
		Content: "Freeing yourself was one thing, claiming ownership of that freed self was another.",
	},
	{
		Author: "Dean KoontzFear Nothing",
		AuthorAvatar: "https://images.gr-assets.com/authors/1487353807p2/9355.jpg",
		Content: "Never leave a friend behind. Friends are all we have to get us through this life--and they are the only things from this world that we could hope to see in the next.",
	},
	{
		Author: "Neil GaimanThe Ocean at the End of the Lane",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Grown-ups don't look like grown-ups on the inside either. Outside, they're big and thoughtless and they always know what they're doing. Inside, they look just like they always have. Like they did when they were your age. Truth is, there aren't any grown-ups. Not one, in the whole wide world.",
	},
	{
		Author: "Robert A. HeinleinTime Enough for Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192826560p2/205.jpg",
		Content: "Never attempt to teach a pig to sing; it wastes your time and annoys the pig.",
	},
	{
		Author: "Jack London",
		AuthorAvatar: "https://images.gr-assets.com/authors/1508674808p2/1240.jpg",
		Content: "You can't wait for inspiration. You have to go after it with a club.",
	},
	{
		Author: "Nora Ephron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1366180104p2/5691.jpg",
		Content: "What are you going to do? Everything, is my guess. It will be a little messy, but embrace the mess. It will be complicated, but rejoice in the complications.",
	},
	{
		Author: "Stephen Fry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400162446p2/10917.jpg",
		Content: "If you know someone who’s depressed, please resolve never to ask them why. Depression isn’t a straightforward response to a bad situation; depression just is, like the weather.Try to understand the blackness, lethargy, hopelessness, and loneliness they’re going through. Be there for them when they come through the other side. It’s hard to be a friend to someone who’s depressed, but it is one of the kindest, noblest, and best things you will ever do.",
	},
	{
		Author: "Groucho Marx",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590177p2/43244.jpg",
		Content: "Humor is reason gone mad.",
	},
	{
		Author: "Victor HugoLes Misérables",
		AuthorAvatar: "https://images.gr-assets.com/authors/1415946858p2/13661.jpg",
		Content: "To love another person is to see the face of God.",
	},
	{
		Author: "Marilyn Monroe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1436929110p2/82952.jpg",
		Content: "We are all born sexual creatures,thank God, but it's a pity so many people despise and crush this natural gift.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Hello, Harry\" said George, beaming at him. \"We thought we heard your dulcet tones.\"\"You don't want to bottle up your anger like that, Harry, let it all out,\" said Fred, also beaming. \"There might be a couple of people fifty miles away who didn't hear you.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "No good deed goes unpunished.",
	},
	{
		Author: "Joseph Mitchell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1265044165p2/123376.jpg",
		Content: "...you can hate a place with all your heart and soul and still be homesick for it.",
	},
	{
		Author: "Tahereh MafiIgnite Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1444252799p2/4637539.jpg",
		Content: "Words, I think, are such unpredictable creatures. No gun, no sword, no army or king will ever be more powerful than a sentence. Swords may cut and kill, but words will stab and stay, burying themselves in our bones to become corpses we carry into the future, all the time digging and failing to rip their skeletons from our flesh.",
	},
	{
		Author: "Margaret AtwoodCat's Eye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "Another belief of mine: that everyone else my age is an adult, whereas I am merely in disguise.",
	},
	{
		Author: "Gautama Buddha",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356856893p2/2167493.jpg",
		Content: "However many holy words you read, however many you speak, what good will they do you if you do not act on upon them?",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "But-\" Maia, still looking at Alec and Magnus, broke off and rasied her eyebrows. Simon turned to see what she was looking at - and stared.Alec had his arms around Magnus and was kissing him full on the mouth. Magnus, who appeared to be in a state of shock, stood frozen. Several groups of people - Shadowhunters and Downworlders alike - were staring and whispering. Glancing to the side, Simon saw the Lightwoods, their eyes widen, gaping at the display. Maryse had her hand over her mouth.Maia looked perplexed. \"Wait a second,\" she said. \"Do we all have to do that, too?",
	},
	{
		Author: "John KeatsOde on a Grecian Urn and Other Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198548090p2/11978.jpg",
		Content: "Heard melodies are sweet, but those unheard, are sweeter",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Jace?\"\"Yeah?\"\"How did you know I had Shadowhunter blood? Was there some way you could tell?\"The elevator arrived with a final groan. Jace unlatched the gate and slid it open. The inside reminded Clary of a birdcage, all black metal and decorative bits of gilt. \"I guessed,\" he said, latching the door behind them. \"It seemed like the most likely explanation.\"\"You guessed? You must have been pretty sure, considering you could have killed me.\"He pressed a button in the wall, and the elevator lurched into action with a vibrating groan that she felt all through the bones in her feet. \"I was ninety percent sure.\"\"I see,\" Clary said.There must have been something in her voice, because he turned to look at her. Her hand cracked across his face, a slap that rocked him back on his heels. He put a hand to his cheek, more in surprise than pain. \"What the hell was that for?\"The other ten percent,\" she said, and they rode the rest of the way down to the street in silence.",
	},
	{
		Author: "Betty SmithA Tree Grows in Brooklyn",
		AuthorAvatar: "",
		Content: "Look at everything always as though you were seeing it either for the first or last time: Thus is your time on earth filled with glory.",
	},
	{
		Author: "أحلام مستغانميالأسود يليق بك",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351458215p2/1109606.jpg",
		Content: "الحداد ليس في ما نرتديه بل في ما نراه . إنّه يكمن في نظرتناللأشياء . بإمكان عيون قلبنا أن تكون في حداد... ولا أحد يدري بذلك",
	},
	{
		Author: "William ShakespeareMacbeth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "To-morrow, and to-morrow, and to-morrow,Creeps in this petty pace from day to day,To the last syllable of recorded time;And all our yesterdays have lighted foolsThe way to dusty death. Out, out, brief candle!Life's but a walking shadow, a poor player,That struts and frets his hour upon the stage,And then is heard no more. It is a taleTold by an idiot, full of sound and fury,Signifying nothing.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "We are each our own devil, and we make this world our hell.",
	},
	{
		Author: "Harper LeeTo Kill a Mockingbird",
		AuthorAvatar: "https://images.gr-assets.com/authors/1188820730p2/1825.jpg",
		Content: "Atticus said to Jem one day, \"I’d rather you shot at tin cans in the backyard, but I know you’ll go after birds. Shoot all the blue jays you want, if you can hit ‘em, but remember it’s a sin to kill a mockingbird.\" That was the only time I ever heard Atticus say it was a sin to do something, and I asked Miss Maudie about it. \"Your father’s right,\" she said. \"Mockingbirds don’t do one thing except make music for us to enjoy. They don’t eat up people’s gardens, don’t nest in corn cribs, they don’t do one thing but sing their hearts out for us. That’s why it’s a sin to kill a mockingbird.",
	},
	{
		Author: "William S. Burroughs",
		AuthorAvatar: "https://images.gr-assets.com/authors/1459243207p2/4462369.jpg",
		Content: "A paranoid is someone who knows a little of what's going on. ",
	},
	{
		Author: "Douglas AdamsMostly Harmless",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "Nothing travels faster than the speed of light, with the possible exception of bad news, which obeys its own special laws.",
	},
	{
		Author: "China MiévilleKing Rat",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243988363p2/33918.jpg",
		Content: "A trap is only a trap if you don't know about it. If you know about it, it's a challenge.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Clary felt suddenly annoyed. \"When the self-congratulatory part of the evening is over, maybe we could get back to saving my best friend from being exsanguinated to death?\"\"Exsanguinated,\" said Jace, impressed. \"That's a big word.\"\"And you're a big-\"\"Tsk tsk,\" he interupted. \"No swearing in church.",
	},
	{
		Author: "Norman Vincent Peale",
		AuthorAvatar: "https://images.gr-assets.com/authors/1293010421p2/8435.jpg",
		Content: "Shoot for the moon. Even if you miss, you'll land among the stars.",
	},
	{
		Author: "Chuck PalahniukSurvivor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "A girl calls and asks, \"Does it hurt very much to die?\"\"Well, sweetheart,\" I tell her, \"yes, but it hurts a lot more to keep living.",
	},
	{
		Author: "Steve MaraboliLife, the Truth, and Being Free",
		AuthorAvatar: "https://images.gr-assets.com/authors/1515015443p2/4491185.jpg",
		Content: "Incredible change happens in your life when you decide to take control of what you do have power over instead of craving control over what you don't.",
	},
	{
		Author: "Ira Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214341880p2/113989.jpg",
		Content: "Nobody tells this to people who are beginners, I wish someone told me. All of us who do creative work, we get into it because we have good taste. But there is this gap. For the first couple years you make stuff, it’s just not that good. It’s trying to be good, it has potential, but it’s not. But your taste, the thing that got you into the game, is still killer. And your taste is why your work disappoints you. A lot of people never get past this phase, they quit. Most people I know who do interesting, creative work went through years of this. We know our work doesn’t have this special thing that we want it to have. We all go through this. And if you are just starting out or you are still in this phase, you gotta know its normal and the most important thing you can do is do a lot of work. Put yourself on a deadline so that every week you will finish one story. It is only by going through a volume of work that you will close that gap, and your work will be as good as your ambitions. And I took longer to figure out how to do this than anyone I’ve ever met. It’s gonna take awhile. It’s normal to take awhile. You’ve just gotta fight your way through.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Sebastian just smiled. “I could hear your heart beating,",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "One day you will ask me which is more important? My life or yours? I will say mine and you will walk away not knowing that you are my life.",
	},
	{
		Author: "Lemony SnicketThe Slippery Slope",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "Well-read people are less likely to be evil.",
	},
	{
		Author: "Oscar WildeA Woman of No Importance",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "After a good dinner one can forgive anybody, even one's own relations.",
	},
	{
		Author: "William Shakespeare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "My only love sprung from my only hate!Too early seen unknown, and known too late!Prodigious birth of love it is to me,That I must love a loathed enemy.",
	},
	{
		Author: "Anne EnrightThe Gathering",
		AuthorAvatar: "https://images.gr-assets.com/authors/1426012522p2/52832.jpg",
		Content: "People do not change, they are merely revealed.",
	},
	{
		Author: "Gregory MaguireWicked: The Life and Times of the Wicked Witch of the West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1319068553p2/7025.jpg",
		Content: "Remember this: Nothing is written in the stars. Not these stars, nor any others. No one controls your destiny.",
	},
	{
		Author: "Louisa May AlcottLittle Women",
		AuthorAvatar: "https://images.gr-assets.com/authors/1200326665p2/1315.jpg",
		Content: "I am not afraid of storms, for I am learning how to sail my ship.",
	},
	{
		Author: "Anne SextonThe Complete Poems",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525637942p2/26814.jpg",
		Content: "Watch out for intellect,because it knows so much it knows nothingand leaves you hanging upside down,mouthing knowledge as your heartfalls out of your mouth.",
	},
	{
		Author: "Jonathan Safran Foer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "I think and think and think, I‘ve thought myself out of happiness one million times, but never once into it.",
	},
	{
		Author: "Charles Baudelaire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371208630p2/13847.jpg",
		Content: "Always be a poet, even in prose.",
	},
	{
		Author: "Ben Okri",
		AuthorAvatar: "https://images.gr-assets.com/authors/1294662937p2/31425.jpg",
		Content: "The most authentic thing about us is our capacity to create, to overcome, to endure, to transform, to love and to be greater than our suffering.",
	},
	{
		Author: "Paulo CoelhoEleven Minutes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "Passion makes a person stop eating, sleeping, working, feeling at peace. A lot of people are frightened because, when it appears, it demolishes all the old things it finds in its path. No one wants their life thrown into chaos. That is why a lot of people keep that threat under control, and are somehow capable of sustaining a house or a structure that is already rotten. They are the engineers of the superseded. Other people think exactly the opposite: they surrender themselves without a second thought, hoping to find in passion the solutions to all their problems. They make the other person responsible for their happiness and blame them for their possible unhappiness. They are either euphoric because something marvelous has happened or depressed because something unexpected has just ruined everything. Keeping passion at bay or surrendering blindly to it - which of these two attitudes is the least destructive? I don't know.",
	},
	{
		Author: "Leigh BardugoSix of Crows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437274413p2/4575289.jpg",
		Content: "Kaz leaned back. \"What's the easiest way to steal a man's wallet?\"\"Knife to the throat?\" asked Inej.\"Gun to the back?\" said Jesper.\"Poison in his cup?\" suggested Nina.\"You're all horrible,\" said Matthias.",
	},
	{
		Author: "Judy Blume",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430843076p2/12942.jpg",
		Content: "Our finger prints don't fade from the lives we touch.",
	},
	{
		Author: "Henry David Thoreau",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "The question is not what you look at, but what you see.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "The earth laughs in flowers.",
	},
	{
		Author: "Cassandra ClareClockwork Prince",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "If no one cares for you at all, do you even really exist?",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "I clench his hands to the point of pain. \"Stay with me.\"His pupils contract to pinpoints, dialate again rapidly, and then return to something resembling normalcy. \"Always,\" he murmurs.",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "In the midst of winter, I found there was, within me, an invincible summer.And that makes me happy. For it says that no matter how hard the world pushes against me, within me, there’s something stronger – something better, pushing right back.",
	},
	{
		Author: "Audrey NiffeneggerThe Time Traveler's Wife",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367342548p2/498072.jpg",
		Content: "Love the world and yourself in it, move through it as though it offers no resistance, as though the world is your natural element.",
	},
	{
		Author: "Elizabeth GilbertEat, pray, love: one woman's search for everything",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "To lose balance sometimes for love is part of living a balancedlife.",
	},
	{
		Author: "Stephen KingSkeleton Crew",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "A short story is a different thing altogether – a short story is like a quick kiss in the dark from a stranger.",
	},
	{
		Author: "Leo Tolstoy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "Only people who are capable of loving strongly can also suffer great sorrow, but this same necessity of loving serves to counteract their grief and heals them.",
	},
	{
		Author: "James Lecesne",
		AuthorAvatar: "",
		Content: "This is how it works. I love the people in my life, and I do for my friends whatever they need me to do for them, again and again, as many times as is necessary. For example, in your case you always forgot who you are and how much you're loved. So what I do for you as your friend is remind you who you are and tell you how much I love you. And this isn't any kind of burden for me, because I love who you are very much. Every time I remind you, I get to remember with you, which is my pleasure.",
	},
	{
		Author: "Will Rogers",
		AuthorAvatar: "https://images.gr-assets.com/authors/1250044478p2/132444.jpg",
		Content: "Never miss a good chance to shut up.",
	},
	{
		Author: "Markus ZusakThe Book Thief",
		AuthorAvatar: "https://images.gr-assets.com/authors/1376268260p2/11466.jpg",
		Content: "A small but noteworthy note. I've seen so many young men over the years who think they're running at other young men. They are not. They are running at me.",
	},
	{
		Author: "Stephen R. CoveyThe 7 Habits of Highly Effective People: Powerful Lessons in Personal Change",
		AuthorAvatar: "https://images.gr-assets.com/authors/1321654785p2/1538.jpg",
		Content: "But until a person can say deeply and honestly, \"I am what I am today because of the choices I made yesterday,\" that person cannot say, \"I choose otherwise.",
	},
	{
		Author: "Christopher MooreThe Stupidest Angel: A Heartwarming Tale of Christmas Terror",
		AuthorAvatar: "https://images.gr-assets.com/authors/1460399391p2/16218.jpg",
		Content: "People, generally, suck.",
	},
	{
		Author: "R.J. PalacioWonder",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351246501p2/4859212.jpg",
		Content: "I think there should be a rule that everyone in the world should get a standing ovation at least once in their lives.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "I want you always to remember me. Will you remember that I existed, and that I stood next to you here like this?",
	},
	{
		Author: "Abraham Lincoln",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198518158p2/229.jpg",
		Content: "When I do good, I feel good. When I do bad, I feel bad. That's my religion.",
	},
	{
		Author: "Charles DarwinThe Life \u0026 Letters of Charles Darwin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398693802p2/12793.jpg",
		Content: "A man who dares to waste one hour of time has not discovered the value of life.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "If I am not good to myself, how can I expect anyone else to be good to me?",
	},
	{
		Author: "Emily Dickinson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536260p2/7440.jpg",
		Content: "If I read a book and it makes my whole body so cold no fire can ever warm me, I know that is poetry.",
	},
	{
		Author: "Pablo PicassoPablo Picasso: Metamorphoses of the Human Form : Graphic Works, 1895-1972",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198536109p2/3253.jpg",
		Content: "Others have seen what is and asked why. I have seen what could be and asked why not. ",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Ask for what you want and be prepared to get it!",
	},
	{
		Author: "E.E. Cummings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "Whenever you think or you believe or you know, you're a lot of other people: but the moment you feel, you're nobody-but-yourself.",
	},
	{
		Author: "John Grisham",
		AuthorAvatar: "https://images.gr-assets.com/authors/1413390525p2/721.jpg",
		Content: "In life, finding a voice is speaking and living the truth. Each of you is an original. Each of you has a distinctive voice. When you find it, your story will be told. You will be heard.",
	},
	{
		Author: "Oscar Wilde",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Whenever people agree with me I always feel I must be wrong.",
	},
	{
		Author: "Judith McNaughtRemember When",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246151270p2/9885.jpg",
		Content: "There will be a few times in your life when all your instincts will tell you to do something, something that defies logic, upsets your plans, and may seem crazy to others. When that happens, you do it. Listen to your instincts and ignore everything else. Ignore logic, ignore the odds, ignore the complications, and just go for it.",
	},
	{
		Author: "Rainer Maria RilkeLetters to a Young Poet",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493785350p2/7906.jpg",
		Content: "Therefore, dear Sir, love your solitude and try to sing out with the pain it causes you. For those who are near you are far away... and this shows that the space around you is beginning to grow vast.... be happy about your growth, in which of course you can't take anyone with you, and be gentle with those who stay behind; be confident and calm in front of them and don't torment them with your doubts and don't frighten them with your faith or joy, which they wouldn't be able to comprehend. Seek out some simple and true feeling of what you have in common with them, which doesn't necessarily have to alter when you yourself change again and again; when you see them, love life in a form that is not your own and be indulgent toward those who are growing old, who are afraid of the aloneness that you trust.... and don't expect any understanding; but believe in a love that is being stored up for you like an inheritance, and have faith that in this love there is a strength and a blessing so large that you can travel as far as you wish without having to step outside it.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "The most terrible poverty is loneliness, and the feeling of being unloved.",
	},
	{
		Author: "George Washington BurnapThe Sphere and Duties of Woman: A Course of Lectures",
		AuthorAvatar: "",
		Content: "The grand essentials to happiness in this life are something to do, something to love, and something to hope for.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Whoever is careless with the truth in small matters cannot be trusted with important matters",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Man often becomes what he believes himself to be. If I keep on saying to myself that I cannot do a certain thing, it is possible that I may end by really becoming incapable of doing it. On the contrary, if I have the belief that I can do it, I shall surely acquire the capacity to do it even if I may not have it at the beginning.",
	},
	{
		Author: "Bob Newhart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1221616466p2/3829.jpg",
		Content: "I think you should be a child for as long as you can. I have been successful for 74 years being able to do that.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "God alert!\" Blackjack yelled. \"It's the wine dude!Mr. D sighed in exasperation. \"The next person, or horse, who calls me the 'wine dude' will end up in a bottle of Merlot!",
	},
	{
		Author: "Chaim PotokThe Chosen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1472717039p2/7385.jpg",
		Content: "I've begun to realize that you can listen to silence and learn from it. It has a quality and a dimension all its own.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Most men would rather deny a hard truth than face it.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Children begin by loving their parents; as they grow older they judge them; sometimes they forgive them.",
	},
	{
		Author: "C.S. LewisA Grief Observed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "No one ever told me that grief felt so like fear.",
	},
	{
		Author: "John GreenThe Fault in Our Stars",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "What else? She is so beautiful. You don’t get tired of looking at her. You never worry if she is smarter than you: You know she is. She is funny without ever being mean. I love her. I am so lucky to love her, Van Houten. You don’t get to choose if you get hurt in this world, old man, but you do have some say in who hurts you. I like my choices. I hope she likes hers.",
	},
	{
		Author: "Gena ShowalterSeduce the Darkness",
		AuthorAvatar: "https://images.gr-assets.com/authors/1447177534p2/48192.jpg",
		Content: "I don't hate you.. I just don't like that you exist",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "I believe in God, but not as one thing, not as an old man in the sky. I believe that what people call God is something in all of us. I believe that what Jesus and Mohammed and Buddha and all the rest said was right. It's just that the translations have gone wrong.",
	},
	{
		Author: "E.E. CummingsPoems, 1923-1954",
		AuthorAvatar: "https://images.gr-assets.com/authors/1512865727p2/10547.jpg",
		Content: "I will take the sun in my mouthand leap into the ripe air Alive with closed eyesto dash against darkness",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "I like people too much or not at all. I've got to go down deep, to fall into people, to really know them.",
	},
	{
		Author: "عمر طاهر",
		AuthorAvatar: "https://images.gr-assets.com/authors/1476909836p2/1569475.jpg",
		Content: "مش مهم انك تغير الكون... المهم انك تخلي الكون ما يغيركش",
	},
	{
		Author: "Bill WattersonThe Calvin And Hobbes:  Tenth Anniversary Book",
		AuthorAvatar: "https://images.gr-assets.com/authors/1374016829p2/13778.jpg",
		Content: "People think it must be fun to be a super genius, but they don't realize how hard it is to put up with all the idiots in the world.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "The best way out is always through.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "لا تجزع من جرحك، وإلا فكيف للنور أن يتسلل إلى باطنك؟",
	},
	{
		Author: "Albert Camus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Blessed are the hearts that can bend; they shall never be broken.",
	},
	{
		Author: "William NicholsonShadowlands",
		AuthorAvatar: "https://images.gr-assets.com/authors/1290447569p2/46130.jpg",
		Content: "I pray because I can't help myself. I pray because I'm helpless. I pray because the need flows out of me all the time- waking and sleeping. It doesn't change God- it changes me.",
	},
	{
		Author: "Lemony Snicket",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "If you are a student you should always get a good nights sleep unless you have come to the good part of your book, and then you should stay up all night and let your schoolwork fall by the wayside, a phrase which means 'flunk'.",
	},
	{
		Author: "Douglas AdamsThe Hitchhiker's Guide to the Galaxy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1189120061p2/4.jpg",
		Content: "He felt that his whole life was some kind of dream and he sometimes wondered whose it was and whether they were enjoying it.",
	},
	{
		Author: "George Bernard Shaw",
		AuthorAvatar: "https://images.gr-assets.com/authors/1271683549p2/5217.jpg",
		Content: "Those who cannot change their minds cannot change anything.",
	},
	{
		Author: "Percy Bysshe Shelley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212686605p2/45882.jpg",
		Content: "The sunlight claps the earth, and the moonbeams kiss the sea: what are all these kissings worth, if thou kiss not me?",
	},
	{
		Author: "Alan MooreV for Vendetta",
		AuthorAvatar: "https://images.gr-assets.com/authors/1304944713p2/3961.jpg",
		Content: "Behind this mask there is more than just flesh. Beneath this mask there is an idea... and ideas are bulletproof.",
	},
	{
		Author: "Erica Jong",
		AuthorAvatar: "https://images.gr-assets.com/authors/1286537368p2/6085.jpg",
		Content: "I have accepted fear as part of life – specifically the fear of change... I have gone ahead despite the pounding in the heart that says: turn back....",
	},
	{
		Author: "Leigh BardugoSix of Crows",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437274413p2/4575289.jpg",
		Content: "Many boys will bring you flowers. But someday you'll meet a boy who will learn your favorite flower, your favorite song, your favorite sweet. And even if he is too poor to give you any of them, it won't matter because he will have taken the time to know you as no one else does. Only that boy earns your heart.",
	},
	{
		Author: "Sylvia PlathAriel: The Restored Edition",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "If the moon smiled, she would resemble you.You leave the same impression Of something beautiful, but annihilating.",
	},
	{
		Author: "Woody Allen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501608198p2/10356.jpg",
		Content: "Life doesn't imitate art, it imitates bad television.",
	},
	{
		Author: "Brian SelznickThe Invention of Hugo Cabret",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201028714p2/38120.jpg",
		Content: "I address you all tonight for who you truly are: wizards, mermaids, travelers, adventurers, and magicians. You are the true dreamers.",
	},
	{
		Author: "Douglas CouplandLife After God",
		AuthorAvatar: "https://images.gr-assets.com/authors/1489283217p2/1886.jpg",
		Content: "And then I felt sad because I realized that once people are broken in certain ways, they can't ever be fixed, and this is something nobody ever tells you when you are young and it never fails to surprise you as you grow older as you see the people in your life break one by one. You wonder when your turn is going to be, or if it's already happened.",
	},
	{
		Author: "LeVar Burton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1516355258p2/160676.jpg",
		Content: "For me, literacy means freedom. For the individual and for society.",
	},
	{
		Author: "Fyodor DostoyevskyNotes from Underground, White Nights, The Dream of a Ridiculous Man, and Selections from The House of the Dead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "Man only likes to count his troubles; he doesn't calculate his happiness.",
	},
	{
		Author: "Ernest HemingwayA Farewell to Arms",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406040005p2/1455.jpg",
		Content: "All thinking men are atheists.",
	},
	{
		Author: "Maxim GorkyThe Lower Depths and Other Plays",
		AuthorAvatar: "https://images.gr-assets.com/authors/1431018607p2/3072763.jpg",
		Content: "Happiness always looks small while you hold it in your hands, but let it go, and you learn at once how big and precious it is.",
	},
	{
		Author: "Simone de Beauvoir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1382722690p2/5548.jpg",
		Content: "I am too intelligent, too demanding, and too resourceful for anyone to be able to take charge of me entirely. No one knows me or loves me completely. I have only myself",
	},
	{
		Author: "Anaïs NinHenry and June: From \"A Journal of Love\"--The Unexpurgated Diary of Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "I hate men who are afraid of women's strength.",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "There's a bluebird in my heart that wants to get outbut I'm too tough for him,I say, stay in there, I'm not going to let anybody see you.",
	},
	{
		Author: "Nicholas SparksThe Guardian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "You always have a choice. It's just that some people make the wrong one.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Silence is so freaking loud",
	},
	{
		Author: "Stephen King",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362814142p2/3389.jpg",
		Content: "Alone. Yes, that's the key word, the most awful word in the English tongue. Murder doesn't hold a candle to it and hell is only a poor synonym.",
	},
	{
		Author: "Suzanne CollinsCatching Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "The bird, the pin, the song, the berries, the watch, the cracker, the dress that burst into flames. I am the mockingjay. The one that survived despite the Capitol's plans. The symbol of the rebellion.",
	},
	{
		Author: "Sigrid UndsetKristin Lavransdatter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1352776655p2/4203.jpg",
		Content: "All my days I have longed equally to travel the right road and to take my own errant path.",
	},
	{
		Author: "Kahlil Gibran",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353732571p2/6466154.jpg",
		Content: "لا تجالس أنصاف العشاق، ولا تصادق أنصاف الأصدقاء، لا تقرأ لأنصاف الموهوبين،لا تعش نصف حياة، ولا تمت نصف موت،لا تختر نصف حل، ولا تقف في منتصف الحقيقة، لا تحلم نصف حلم، ولا تتعلق بنصف أمل، إذا صمتّ.. فاصمت حتى النهاية، وإذا تكلمت.. فتكلّم حتى النهاية، لا تصمت كي تتكلم، ولا تتكلم كي تصمت. إذا رضيت فعبّر عن رضاك، لا تصطنع نصف رضا، وإذا رفضت.. فعبّر عن رفضك، لأن نصف الرفض قبول.. النصف هو حياة لم تعشها، وهو كلمة لم تقلها،وهو ابتسامة أجّلتها، وهو حب لم تصل إليه، وهو صداقة لم تعرفها.. النصف هو ما يجعلك غريباً عن أقرب الناس إليك، وهو ما يجعل أقرب الناس إليك غرباء عنك. النصف هو أن تصل وأن لاتصل، أن تعمل وأن لا تعمل،أن تغيب وأن تحضر.. النصف هو أنت، عندما لا تكون أنت.. لأنك لم تعرف من أنت، النصف هو أن لا تعرف من أنت.. ومن تحب ليس نصفك الآخر.. هو أنت في مكان آخر في الوقت نفسه. نصف شربة لن تروي ظمأك، ونصف وجبة لن تشبع جوعك،نصف طريق لن يوصلك إلى أي مكان، ونصف فكرة لن تعطي لك نتيجة النصف هو لحظة عجزك وأنت لست بعاجز.. لأنك لست نصف إنسان. أنت إنسان وجدت كي تعيش الحياة، وليس كي تعيش نصف حياة ليست حقيقة الإنسان بما يظهره لك.. بل بما لا يستطيع أن يظهره، لذلك.. إذا أردت أن تعرفه فلا تصغي إلى ما يقوله .. بل إلى ما لا يقوله.",
	},
	{
		Author: "Nicholas SparksThe Notebook",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "In times of grief and sorrow I will hold you and rock you and take your grief and make it my own. When you cry I cry and when you hurt I hurt. And together we will try to hold back the floods to tears and despair and make it through the potholed street of life",
	},
	{
		Author: "Anna Akhmatova",
		AuthorAvatar: "https://images.gr-assets.com/authors/1487037121p2/99703.jpg",
		Content: "You will hear thunder and remember me,and think: she wanted storms...",
	},
	{
		Author: "Patricia Cornwell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1486394658p2/1025097.jpg",
		Content: "Do no harm and leave the world a better place than you found it.",
	},
	{
		Author: "Sarah Kay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1327230533p2/11377.jpg",
		Content: "If I should have a daughter…“Instead of “Mom",
	},
	{
		Author: "Cassandra Clare",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "It's the mortal cup Jace, not the mortal toilet bowl.",
	},
	{
		Author: "Jonathan Safran FoerEverything Is Illuminated",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "She was a genius of sadness, immersing herself in it, separating its numerous strands, appreciating its subtle nuances. She was a prism through which sadness could be divided into its infinite spectrum.",
	},
	{
		Author: "Charles DickensA Tale of Two Cities",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "It is a far, far better thing that I do, than I have ever done; it is a far, far better rest that I go to than I have ever known.",
	},
	{
		Author: "Oscar WildeThe Picture of Dorian Gray",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "When one is in love, one always begins by deceiving one's self, and one always ends by deceiving others. That is what the world calls a romance.",
	},
	{
		Author: "Neil GaimanM Is for Magic",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "Stories you read when you're the right age never quite leave you. You may forget who wrote them or what the story was called. Sometimes you'll forget precisely what happened, but if a story touches you it will stay with you, haunting the places in your mind that you rarely ever visit.",
	},
	{
		Author: "Richelle MeadFrostbite",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "There's nothing worse than waiting and not knowing what'll happen to you. Your own imagination can be crueler than any captor.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Accio Brain!",
	},
	{
		Author: "Cheris Kramarae",
		AuthorAvatar: "https://images.gr-assets.com/authors/1252536090p2/83683.jpg",
		Content: "Feminism is the radical notion that women are human beings.",
	},
	{
		Author: "Leo TolstoyAnna Karenina",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "Spring is the time of plans and projects.",
	},
	{
		Author: "Plato",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393978633p2/879.jpg",
		Content: "The price good men pay for indifference to public affairs is to be ruled by evil men.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "That's the thing about girls. Every time they do something pretty, even if they're not much to look at, or even if they're sort of stupid, you fall in love with them, and then you never know where the hell you are. Girls. Jesus Christ. They can drive you crazy. They really can.",
	},
	{
		Author: "Neil deGrasse Tyson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1423292563p2/12855.jpg",
		Content: "The good thing about science is that it's true whether or not you believe in it.",
	},
	{
		Author: "Dalai Lama XIV",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241989386p2/570218.jpg",
		Content: "My religion is very simple. My religion is kindness.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "But Dumbledore says he doesn't care what they do as long as they don't take him off the Chocolate Frog cards.",
	},
	{
		Author: "Aldous Huxley",
		AuthorAvatar: "https://images.gr-assets.com/authors/1499982247p2/3487.jpg",
		Content: "Maybe this world is another planet’s hell.",
	},
	{
		Author: "Robert J. SawyerCalculating God",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224975910p2/25883.jpg",
		Content: "Learning to ignore things is one of the great paths to inner peace.",
	},
	{
		Author: "Ken KeseyOne Flew Over the Cuckoo's Nest",
		AuthorAvatar: "https://images.gr-assets.com/authors/1363273797p2/7285.jpg",
		Content: "Man, when you lose your laugh you lose your footing.",
	},
	{
		Author: "Walt Whitman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392303683p2/1438.jpg",
		Content: "Keep your face always toward the sunshine - and shadows will fall behind you.",
	},
	{
		Author: "نزار قباني",
		AuthorAvatar: "https://images.gr-assets.com/authors/1260699865p2/850723.jpg",
		Content: "In the summerI stretch out on the shoreAnd think of you. Had I told the seaWhat I felt for you,It would have left its shores,Its shells,Its fish,And followed me.",
	},
	{
		Author: "Aristotle",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390143800p2/2192.jpg",
		Content: "Educating the mind without educating the heart is no education at all.",
	},
	{
		Author: "Neil GaimanPreludes \u0026 Nocturnes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234150163p2/1221698.jpg",
		Content: "You get what anybody gets - you get a lifetime.",
	},
	{
		Author: "John GreenAn Abundance of Katherines",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I figured something out. The future is unpredictable.",
	},
	{
		Author: "John Updike",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419249254p2/6878.jpg",
		Content: "Dreams come true. Without that possibility, nature would not incite us to have them. ",
	},
	{
		Author: "Nelson DeMille",
		AuthorAvatar: "https://images.gr-assets.com/authors/1498066748p2/1238.jpg",
		Content: "The problem with doing nothing is not knowing when you are finished.",
	},
	{
		Author: "Nicholas SparksNights in Rodanthe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "While I sleep, I dream of you, and when I wake, I long to hold you in my arms. If anything, our time apart has only made me more certain that I want to spend my nights by your side, and my days with your heart.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "The greatness of a nation and its moral progress can be judged by the way its animals are treated.",
	},
	{
		Author: "Cormac McCarthyThe Road",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414695980p2/4178.jpg",
		Content: "You forget what you want to remember, and you remember what you want to forget.",
	},
	{
		Author: "Ray Bradbury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "If we listened to our intellect we'd never have a love affair. We'd never have a friendship. We'd never go in business because we'd be cynical: \"It's gonna go wrong.\" Or \"She's going to hurt me.\" Or,\"I've had a couple of bad love affairs, so therefore . . .\" Well, that's nonsense. You're going to miss life. You've got to jump off the cliff all the time and build your wings on the way down.",
	},
	{
		Author: "Hayley Williams",
		AuthorAvatar: "https://images.gr-assets.com/authors/1329247990p2/4119881.jpg",
		Content: "Sometimes it takes a good fall to really know where you stand",
	},
	{
		Author: "George Orwell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "The most effective way to destroy people is to deny and obliterate their own understanding of their history.",
	},
	{
		Author: "Bruce Lee",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454507945p2/32579.jpg",
		Content: "I’m not in this world to live up to your expectations and you’re not in this world to live up to mine.",
	},
	{
		Author: "Mary RenaultThe Persian Boy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1378247630p2/38185.jpg",
		Content: "One must live as if it would be forever, and as if one might die each moment. Always both at once.",
	},
	{
		Author: "Dalai Lama XIV",
		AuthorAvatar: "https://images.gr-assets.com/authors/1241989386p2/570218.jpg",
		Content: "There is a saying in Tibetan, 'Tragedy should be utilized as a source of strength.'No matter what sort of difficulties, how painful experience is, if we lose our hope, that's our real disaster.",
	},
	{
		Author: "Ludwig van Beethoven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1243246871p2/40589.jpg",
		Content: "Music is ... A higher revelation than all Wisdom \u0026 Philosophy",
	},
	{
		Author: "Ned VizziniIt's Kind of a Funny Story",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341737392p2/11672.jpg",
		Content: "I can't eat and I can't sleep. I'm not doing well in terms of being a functional human, you know?",
	},
	{
		Author: "Margaret Atwood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1282859073p2/3472.jpg",
		Content: "Men are afraid that women will laugh at them. Women are afraid that men will kill them.",
	},
	{
		Author: "J.D. SalingerThe Catcher in the Rye",
		AuthorAvatar: "https://images.gr-assets.com/authors/1288777679p2/819789.jpg",
		Content: "The mark of the immature man is that he wants to die nobly for a cause, while the mark of the mature man is that he wants to live humbly for one.",
	},
	{
		Author: "Jodi PicoultMercy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1475775448p2/7128.jpg",
		Content: "You know it's never fifty-fifty in a marriage. It's always seventy-thirty, or sixty-forty. Someone falls in love first. Someone puts someone else up on a pedestal. Someone works very hard to keep things rolling smoothly; someone else sails along for the ride.",
	},
	{
		Author: "C.S. Lewis",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "I didn’t go to religion to make me happy. I always knew a bottle of Port would do that. If you want a religion to make you feel really comfortable, I certainly don’t recommend Christianity.",
	},
	{
		Author: "Johnny Depp",
		AuthorAvatar: "https://images.gr-assets.com/authors/1317565242p2/88280.jpg",
		Content: "if you love two people at the same time, choose the second. Because if you really loved the first one, you wouldn't have fallen for the second.",
	},
	{
		Author: "Marcus AureliusMeditations",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430676293p2/17212.jpg",
		Content: "The happiness of your life depends upon the quality of your thoughts.",
	},
	{
		Author: "Mary  StewartThe Hollow Hills",
		AuthorAvatar: "https://images.gr-assets.com/authors/1210367214p2/15590.jpg",
		Content: "Every life has death and every light has shadow. Be content to stand in the light and let the shadow fall where it will.",
	},
	{
		Author: "Napoléon Bonaparte",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397332087p2/210910.jpg",
		Content: "Courage isn't having the strength to go on - it is going on when you don't have strength.",
	},
	{
		Author: "John AnsterThe First Part of Goethe's Faust",
		AuthorAvatar: "",
		Content: "Whatever you can do or dream you can, begin it.Boldness has genius, power and magic in it!",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "It's so strange how life works: You want something and you wait and wait and feel like it's taking forever to come. Then it happens and it's over and all you want to do is curl back up in that moment before things changed.",
	},
	{
		Author: "Albert CamusA Happy Death",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "When I look at my life and its secret colours, I feel like bursting into tears.",
	},
	{
		Author: "Franz Kafka",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495464914p2/5223.jpg",
		Content: "I think we ought to read only the kind of books that wound or stab us. If the book we're reading doesn't wake us up with a blow to the head, what are we reading for? So that it will make us happy, as you write? Good Lord, we would be happy precisely if we had no books, and the kind of books that make us happy are the kind we could write ourselves if we had to. But we need books that affect us like a disaster, that grieve us deeply, like the death of someone we loved more than ourselves, like being banished into forests far from everyone, like a suicide. A book must be the axe for the frozen sea within us. That is my belief.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "There is no such thing as a \"broken family.\" Family is family, and is not determined by marriage certificates, divorce papers, and adoption documents. Families are made in the heart. The only time family becomes null is when those ties in the heart are cut. If you cut those ties, those people are not your family. If you make those ties, those people are your family. And if you hate those ties, those people will still be your family because whatever you hate will always be with you.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "The things we love destroy us every time, lad. Remember that.",
	},
	{
		Author: "Allen Ginsberg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1421583811p2/4261.jpg",
		Content: "I don't think there is any truth. There are only points of view. ",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "If the person you are talking to doesn't appear to be listening, be patient. It may simply be that he has a small piece of fluff in his ear.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "Throw your dreams into space like a kite, and you do not know what it will bring back, a new life, a new friend, a new love, a new country.",
	},
	{
		Author: "Walt Disney Company",
		AuthorAvatar: "https://images.gr-assets.com/authors/1289112683p2/3510823.jpg",
		Content: "The way to get started is to quit talking and begin doing. ",
	},
	{
		Author: "Joseph Joubert",
		AuthorAvatar: "https://images.gr-assets.com/authors/1249311953p2/313376.jpg",
		Content: "The worst thing about new books is that they keep us from reading the old ones.",
	},
	{
		Author: "Sue GraftonM is for Malice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1319510564p2/9559.jpg",
		Content: "Ghosts don't haunt us. That's not how it works. They're present among us because we won't let go of them.",
	},
	{
		Author: "Richard WrightBlack Boy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209848296p2/9657.jpg",
		Content: "Whenever my environment had failed to support or nourish me, I had clutched at books...",
	},
	{
		Author: "Robin Williams",
		AuthorAvatar: "",
		Content: "See, the problem is that God gives men a brain and a penis, and only enough blood to run one at a time.",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "There's no need to clarify my finger snap,\" said Magnus. \"The implication was clear in the snap itself.",
	},
	{
		Author: "Richard Evans",
		AuthorAvatar: "",
		Content: "It is often in the darkest skies that we see the brightest stars.",
	},
	{
		Author: "Maria Montessori",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232159605p2/34106.jpg",
		Content: "Imagination does not become great until human beings, given the courage and the strength, use it to create.",
	},
	{
		Author: "Agostinho da Silva",
		AuthorAvatar: "https://images.gr-assets.com/authors/1345064734p2/1223934.jpg",
		Content: "I am not interested in being original. I am interested in being true.",
	},
	{
		Author: "Cassandra ClareCity of Fallen Angels",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "You know, some people think Shadowhunters are just myths. Like mummies and genies.\" Kyle grinned at Jace. \"Can you grant wishes?\"\"That depends,\" he said. \"Do you wish to be punched in the face?",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You care so much you feel as though you will bleed to death with the pain of it.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "That was enterprising,\" Will sounded nearly impressed.Nate smiled. Tess shot him a furious look. \"Don't look pleased with yourself. When Will says 'enterprising' he means 'morally deficient.'\"\"No, I mean enterprising,\" said Will. \"When I mean morally deficient, I say, 'Now, that's something I would have done'.",
	},
	{
		Author: "Jane AustenSense and Sensibility",
		AuthorAvatar: "https://images.gr-assets.com/authors/1380085320p2/1265.jpg",
		Content: "It isn't what we say or think that defines us, but what we do.",
	},
	{
		Author: "Lewis Carroll",
		AuthorAvatar: "https://images.gr-assets.com/authors/1192735053p2/8164.jpg",
		Content: "One of the deep secrets of life is that all that is really worth the doing is what we do for others.",
	},
	{
		Author: "Sarah J. MaasA Court of Mist and Fury",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "He thinks he'll be remembered as the villain in the story. But I forgot to tell him that the villain is usually the person who locks up the maiden and throws away the key. He was the one who let me out.",
	},
	{
		Author: "Jojo MoyesMe Before You",
		AuthorAvatar: "https://images.gr-assets.com/authors/1400624880p2/281810.jpg",
		Content: "Push yourself. Don't Settle. Just live well. Just LIVE.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Black hair and blue eyes are my favorite combination.",
	},
	{
		Author: "Hunter S. Thompson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1206560814p2/5237.jpg",
		Content: "So we shall let the reader answer this question for himself: who is the happier man, he who has braved the storm of life and lived or he who has stayed securely on shore and merely existed?",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "It is not the length of life, but the depth.",
	},
	{
		Author: "John Muir",
		AuthorAvatar: "https://images.gr-assets.com/authors/1398092241p2/5297.jpg",
		Content: "The mountains are calling and I must go.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "Dress shabbily and they remember the dress; dress impeccably and they remember the woman.",
	},
	{
		Author: "C. JoyBell C.",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501913788p2/4114218.jpg",
		Content: "The unhappiest people in this world, are those who care the most about what other people think.",
	},
	{
		Author: "Ayn RandThe Fountainhead",
		AuthorAvatar: "https://images.gr-assets.com/authors/1168729178p2/432.jpg",
		Content: "Freedom (n.): To ask nothing. To expect nothing. To depend on nothing.",
	},
	{
		Author: "Sarah J. MaasHeir of Fire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "She was the heir of ash and fire, and she would bow to no one.",
	},
	{
		Author: "Becca FitzpatrickHush, Hush",
		AuthorAvatar: "https://images.gr-assets.com/authors/1390505291p2/2876763.jpg",
		Content: "Coach: \"All right, Patch. let's say you're at a party. the room is full of girls of all shapes and sizes. You see blondes, brunettes, redheads, a few girl with black hair. Some are talkive, while other appear shy. You've one girl who fits your profile - attractive, intelligent and vulnerable. Dow do you let her know you're interested?\"Patch: \"Single her out. Talk to her.\"Coach: \"Good. Now for the big question - how do you know if she's game or if she wants you to move on?\"Patch: \"I study her. I figure out what she's thinking and feeling. She's not gonig to come right out and tell me, which is why i have to pay attention. Does she turn her body toward mine? Does she hold me eyes, then look away? Does she bite her lip and play with her hair, the way Nora is doing right now?",
	},
	{
		Author: "Gabriel García MárquezLove in the Time of Cholera",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502310670p2/13450.jpg",
		Content: "He allowed himself to be swayed by his conviction that human beings are not born once and for all on the day their mothers give birth to them, but that life obliges them over and over again to give birth to themselves.",
	},
	{
		Author: "Napoleon HillThink and Grow Rich: The Landmark Bestseller Now Revised and Updated for the 21st Century",
		AuthorAvatar: "https://images.gr-assets.com/authors/1265381018p2/399.jpg",
		Content: "The starting point of all achievement is DESIRE. Keep this constantly in mind. Weak desire brings weak results, just as a small fire makes a small amount of heat.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "Be awesome! Be a book nut!",
	},
	{
		Author: "Gustave Flaubert",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198541369p2/1461.jpg",
		Content: "Be steady and well-ordered in your life so that you can be fierce and original in your work.",
	},
	{
		Author: "Mark Twain",
		AuthorAvatar: "https://images.gr-assets.com/authors/1322103868p2/1244.jpg",
		Content: "The easy confidence with which I know another man's religion is folly teaches me to suspect that my own is also.",
	},
	{
		Author: "Lauren OliverDelirium",
		AuthorAvatar: "https://images.gr-assets.com/authors/1416335442p2/2936493.jpg",
		Content: "Love: It will kill you and save you, both",
	},
	{
		Author: "Cassandra ClareCity of Lost Souls",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I don't care,\" Clary said. \"He'd do it for me. Tell me he wouldn't. If I were missing-\" \"He'd burn the whole world down till he could dig you out of the ashes. I know,\" Alec said.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "One more time? For the audience?\" he says. His voice isn't angry. It's hollow, which is worse. Already the boy with the bread is slipping away from me.I take his hand, holding on tightly, preparing for the cameras, and dreading the moment when I will finally have to let go.",
	},
	{
		Author: "Toni Morrison",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494211316p2/3534.jpg",
		Content: "You think because he doesn't love you that you are worthless. You think that because he doesn't want you anymore that he is right -- that his judgement and opinion of you are correct. If he throws you out, then you are garbage. You think he belongs to you because you want to belong to him. Don't. It's a bad word, 'belong.' Especially when you put it with somebody you love. Love shouldn't be like that. Did you ever see the way the clouds love a mountain? They circle all around it; sometimes you can't even see the mountain for the clouds. But you know what? You go up top and what do you see? His head. The clouds never cover the head. His head pokes through, beacuse the clouds let him; they don't wrap him up. They let him keep his head up high, free, with nothing to hide him or bind him. You can't own a human being. You can't lose what you don't own. Suppose you did own him. Could you really love somebody who was absolutely nobody without you? You really want somebody like that? Somebody who falls apart when you walk out the door? You don't, do you? And neither does he. You're turning over your whole life to him. Your whole life, girl. And if it means so little to you that you can just give it away, hand it to him, then why should it mean any more to him? He can't value you more than you value yourself.",
	},
	{
		Author: "Woody Allen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501608198p2/10356.jpg",
		Content: "The difference between sex and love is that sex relieves tension and love causes it.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Don't let the muggles get you down.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "The desire to reach for the stars is ambitious. The desire to reach hearts is wise.",
	},
	{
		Author: "André BretonWhat is Surrealism?: Selected Writings",
		AuthorAvatar: "https://images.gr-assets.com/authors/1260768105p2/54133.jpg",
		Content: "My wish is that you may be loved to the point of madness.",
	},
	{
		Author: "J.K. RowlingThe Tales of Beedle the Bard",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "To hurt is as human as to breathe.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "The advantage of a bad memory is that one enjoys several times the same good things for the first time.",
	},
	{
		Author: "Andy Rooney",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209913186p2/35924.jpg",
		Content: "Yes, we praise women over 40 for a multitude of reasons. Unfortunately, it's not reciprocal. For every stunning, smart, well-coiffed, hot woman over 40, there is a bald, paunchy relic in yellow pants making a fool of himself with some 22-year old waitress. Ladies, I apologize. For all those men who say, \"Why buy the cow when you can get the milk for free?\", here's an update for you. Nowadays 80% of women are against marriage. Why? Because women realize it's not worth buying an entire pig just to get a little sausage!",
	},
	{
		Author: "John Lennon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207677340p2/19968.jpg",
		Content: "If everyone demanded peace instead of another television set, then there'd be peace.",
	},
	{
		Author: "Tiffanie DeBartoloHow to Kill a Rock Star",
		AuthorAvatar: "https://images.gr-assets.com/authors/1492733789p2/65959.jpg",
		Content: "Did you really want to die?\"\"No one commits suicide because they want to die.\"\"Then why do they do it?\"\"Because they want to stop the pain.",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Katniss, the girl who was on fire!",
	},
	{
		Author: "Omar KhayyámThe Rubáiyát of Omar Khayyám",
		AuthorAvatar: "https://images.gr-assets.com/authors/1525037614p2/2742325.jpg",
		Content: "Be happy for this moment. This moment is your life.",
	},
	{
		Author: "Virginia Woolf",
		AuthorAvatar: "https://images.gr-assets.com/authors/1419596619p2/6765.jpg",
		Content: "I can only note that the past is beautiful because one never realises an emotion at the time. It expands later, and thus we don't have complete emotions about the present, only about the past.",
	},
	{
		Author: "W.B. Yeats",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440689155p2/29963.jpg",
		Content: "I have spread my dreams under your feet.Tread softly because you tread on my dreams.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "In the end that was the choice you made, and it doesn't matter how hard it was to make it. It matters that you did.",
	},
	{
		Author: "Rumi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1217355543p2/875661.jpg",
		Content: "Lovers don't finally meet somewhere. They're in each other all along.",
	},
	{
		Author: "Chuck Palahniuk",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "If you love something set it free, but don't be surprised if it comes back with herpes.",
	},
	{
		Author: "L.J. SmithSecret Vampire",
		AuthorAvatar: "https://images.gr-assets.com/authors/1266517290p2/50873.jpg",
		Content: "You don't love a girl because of beauty. You love her because she sings a song only you can understand.",
	},
	{
		Author: "Fyodor DostoyevskyNotes from Underground",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506003555p2/3137322.jpg",
		Content: "I say let the world go to hell, but I should always have my tea.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "I was born with an enormous need for affection, and a terrible need to give it.",
	},
	{
		Author: "Anthony Burgess",
		AuthorAvatar: "https://images.gr-assets.com/authors/1484769559p2/5735.jpg",
		Content: "Laugh and the world laughs with you, snore and you sleep alone.",
	},
	{
		Author: "Sarah DessenJust Listen",
		AuthorAvatar: "https://images.gr-assets.com/authors/1372181953p2/2987.jpg",
		Content: "Because this is what happens when you try to run from the past. It just doesn’t catch up, it overtakes … blotting out the future.",
	},
	{
		Author: "Colleen HooverSlammed",
		AuthorAvatar: "https://images.gr-assets.com/authors/1464032240p2/5430144.jpg",
		Content: "Question everything. Your love, your religion, your passion. If you don't have questions, you'll never find answers.",
	},
	{
		Author: "Stieg LarssonThe Girl with the Dragon Tattoo",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246466225p2/706255.jpg",
		Content: "Everyone has secrets. It's just a matter of finding out what they are.",
	},
	{
		Author: "Suzanne CollinsMockingjay",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Ally.\" Peeta says the words slowly, tasting it. \"Friend. Lover. Victor. Enemy. Fiancee. Target. Mutt. Neighbor. Hunter. Tribute. Ally. I'll add it to the list of words I use to try to figure you out. The problem is, I can't tell what's real anymore, and what's made up.",
	},
	{
		Author: "Nicholson Baker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1247028956p2/15882.jpg",
		Content: "Books: a beautifully browsable invention that needs no electricity and exists in a readable form no matter what happens.",
	},
	{
		Author: "Ann Richards",
		AuthorAvatar: "https://images.gr-assets.com/authors/1242787158p2/126242.jpg",
		Content: "After all, Ginger Rogers did everything that Fred Astaire did. She just did it backwards and in high heels.",
	},
	{
		Author: "C.S. LewisThe Weight of Glory, and Other Addresses",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367519078p2/1069006.jpg",
		Content: "It would seem that Our Lord finds our desires not too strong, but too weak. We are half-hearted creatures, fooling about with drink and sex and ambition when infinite joy is offered us, like an ignorant child who wants to go on making mud pies in a slum because he cannot imagine what is meant by the offer of a holiday at the sea. We are far too easily pleased.",
	},
	{
		Author: "Richelle MeadFrostbite",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "what's your name?\"what?\" i asked, squinting at the light.your name.\" I reconized Dr. Olendzki peering over me.you know my name.\"I want you to tell me.\"Rose. Rose Hathaway.\"Do you know your birthday?\"Of course I do. Why are you asking me such stupid things? Did you lose my records?\"Dr. Olendzki gave an exasperated sigh and walked off, taking the annoying light with her. \"I think she's fine,",
	},
	{
		Author: "William Ernest HenleyEchoes of Life and Death;",
		AuthorAvatar: "https://images.gr-assets.com/authors/1228003614p2/718319.jpg",
		Content: "It matters not how strait the gate, How charged with punishments the scroll, I am the master of my fate: I am the captain of my soul.",
	},
	{
		Author: "Charlotte BrontëJane Eyre",
		AuthorAvatar: "https://images.gr-assets.com/authors/1335001351p2/1036615.jpg",
		Content: "Jane, be still; don't struggle so like a wild, frantic bird, that is rending its own plumage in its desperation.\"\"I am no bird; and no net ensnares me; I am a free human being, with an independent will; which I now exert to leave you.",
	},
	{
		Author: "Isaac Asimov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1341965730p2/16667.jpg",
		Content: "If my doctor told me I had only six minutes to live, I wouldn't brood. I'd type a little faster.",
	},
	{
		Author: "Eoin ColferArtemis Fowl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1254336426p2/10896.jpg",
		Content: "It's like learning to ride a unicorn. You never forget.",
	},
	{
		Author: "G.K. Chesterton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1365860649p2/7014283.jpg",
		Content: "Literature is a luxury; fiction is a necessity.",
	},
	{
		Author: "Kiran DesaiThe Inheritance of Loss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232769561p2/31428.jpg",
		Content: "The present changes the past. Looking back you do not find what you left behind.",
	},
	{
		Author: "Pete Wentz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1213747212p2/1267442.jpg",
		Content: "Here's to the kids.The kids who would rather spend their night with a bottle of coke \u0026 Patrick or Sonny playing on their headphones than go to some vomit-stained high school party.Here's to the kids whose 11:11 wish was wasted on one person who will never be there for them.Here's to the kids whose idea of a good night is sitting on the hood of a car, watching the stars.Here's to the kids who never were too good at life, but still were wicked cool.Here's to the kids who listened to Fall Out boy and Hawthorne Heights before they were on MTV...and blame MTV for ruining their life.Here's to the kids who care more about the music than the haircuts. Here's to the kids who have crushes on a stupid lush.Here's to the kids who hum \"A Little Less 16 Candles, A Little More Touch Me\" when they're stuck home, dateless, on a Saturday night.Here's to the kids who have ever had a broken heart from someone who didn't even know they existed.Here's to the kids who have read The Perks of Being a Wallflower \u0026 didn't feel so alone after doing so.Here's to the kids who spend their days in photobooths with their best friend(s). Here's to the kids who are straight up smartasses \u0026 just don't care.Here's to the kids who speak their mind.Here's to the kids who consider screamo their lullaby for going to sleep.Here's to the kids who second guess themselves on everything they do.Here's to the kids who will never have 100 percent confidence in anything they do, and to the kids who are okay with that.Here's to the kids.This one's not for the kids,who always get what they want,But for the ones who never had it at all. It's not for the ones who never got caught, But for the ones who always try and fall. This one's for the kids who didnt make it, We were the kids who never made it. The Overcast girls and the Underdog Boys.Not for the kids who had all their joys. This one's for the kids who never faked it. We're the kids who didn't make it. They say \"Breaking hearts is what we do best,\"And, \"We'll make your heart be ripped of your chest\" The only heart that I broke was mine, When I got My Hopes up too too high.We were the kids who didnt make it. We are the kids who never made it.",
	},
	{
		Author: "Terry PratchettReaper Man",
		AuthorAvatar: "https://images.gr-assets.com/authors/1235562205p2/1654.jpg",
		Content: "Five exclamation marks, the sure sign of an insane mind.",
	},
	{
		Author: "Alysha Speer",
		AuthorAvatar: "https://images.gr-assets.com/authors/1352932197p2/3119687.jpg",
		Content: "Laugh, even when you feel too sick or too worn out or tired. Smile, even when you're trying not to cry and the tears are blurring your vision. Sing, even when people stare at you and tell you your voice is crappy. Trust, even when your heart begs you not to. Twirl, even when your mind makes no sense of what you see. Frolick, even when you are made fun of. Kiss, even when others are watching. Sleep, even when you're afraid of what the dreams might bring. Run, even when it feels like you can't run any more.And, always, remember, even when the memories pinch your heart. Because the pain of all your experience is what makes you the person you are now. And without your experience---you are an empty page, a blank notebook, a missing lyric. What makes you brave is your willingness to live through your terrible life and hold your head up high the next day. So don't live life in fear. Because you are stronger now, after all the crap has happened, than you ever were back before it started.",
	},
	{
		Author: "Tim O'Brien",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232136886p2/2330.jpg",
		Content: "That's what fiction is for. It's for getting at the truth when the truth isn't sufficient for the truth.",
	},
	{
		Author: "Elizabeth StroutAbide with Me",
		AuthorAvatar: "https://images.gr-assets.com/authors/1361387789p2/97313.jpg",
		Content: "I suspect the most we can hope for, and it's no small hope, is that we never give up, that we never stop giving ourselves permission to try to love and receive love.",
	},
	{
		Author: "Chuck PalahniukFight Club",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "Today is the sort of day where the sun only comes up to humiliate you.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "The role of a writer is not to say what we can all say, but what we are unable to say.",
	},
	{
		Author: "Alfred TennysonThe Complete Poetical Works of Tennyson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1454788521p2/13638502.jpg",
		Content: "I am a part of all that I have met.",
	},
	{
		Author: "Nora Roberts",
		AuthorAvatar: "https://images.gr-assets.com/authors/1505847251p2/625.jpg",
		Content: "If you don't go after what you want, you'll never have it. If you don't ask, the answer is always no. If you don't step forward, you're always in the same place.",
	},
	{
		Author: "Herman MelvilleMoby-Dick or, The Whale",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495029910p2/1624.jpg",
		Content: "I know not all that may be coming, but be it what it will, I'll go to it laughing.",
	},
	{
		Author: "Simone de BeauvoirThe Mandarins",
		AuthorAvatar: "https://images.gr-assets.com/authors/1382722690p2/5548.jpg",
		Content: "She was ready to deny the existence of space and time rather than admit that love might not be eternal.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I walk around the school hallways and look at the people. I look at the teachers and wonder why they're here. If they like their jobs. Or us. And I wonder how smart they were when they were fifteen. Not in a mean way. In a curious way. It's like looking at all the students and wondering who's had their heart broken that day, and how they are able to cope with having three quizzes and a book report due on top of that. Or wondering who did the heart breaking. And wondering why.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "Peace cannot be kept by force; it can only be achieved by understanding.",
	},
	{
		Author: "أحمد خالد توفيق",
		AuthorAvatar: "https://images.gr-assets.com/authors/1316682283p2/1479015.jpg",
		Content: "أسوء تعذيب فى العالم هو الشخص المُصر على الكلام بينما أنت مُثقل بالهموم , ترغب فى أن تبقى صامتاً وأن تصغى لأفكارك.",
	},
	{
		Author: "Oliver GoldsmithThe Citizen of the World, Or, Letters from a Chinese Philosopher, Residing in London, to His Friends in the Country, by Dr. Goldsmith",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234787093p2/65124.jpg",
		Content: "Our greatest glory is not in never falling, but in rising every time we fall.",
	},
	{
		Author: "Elizabeth GilbertEat Pray Love: One Woman's Search for Everything Across Italy, India and Indonesia",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "Your emotions are the slaves to your thoughts, and you are the slave to your emotions.",
	},
	{
		Author: "Czesław Miłosz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414843475p2/84259.jpg",
		Content: "Not that I want to be a god or a hero. Just to change into a tree, grow for ages, not hurt anyone.",
	},
	{
		Author: "John Henry Newman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1494112023p2/24706.jpg",
		Content: "I sought to hear the voice of God and climbed the topmost steeple, but God declared: \"Go down again - I dwell among the people.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Poetry is what gets lost in translation.",
	},
	{
		Author: "Richelle MeadShadow Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "He stepped back and threw his arms out.\"I'm always crazy around you Rose. Here, I'm going to write an impromptu poem for you.\"He tipped his head back and shouted to the sky:\"Rose is in redBut never in blueSharp as a thornFights like one too.",
	},
	{
		Author: "Ray BradburyFahrenheit 451",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445955959p2/1630.jpg",
		Content: "We need not to be let alone. We need to be really bothered once in a while. How long is it since you were really bothered? About something important, about something real?",
	},
	{
		Author: "William ShakespeareA Midsummer Night's Dream",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "Lord, what fools these mortals be!",
	},
	{
		Author: "Leo Tolstoy",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091388p2/128382.jpg",
		Content: "If you want to be happy, be.",
	},
	{
		Author: "Rebecca WalkerBaby Love: Choosing Motherhood After a Lifetime of Ambivalence",
		AuthorAvatar: "https://images.gr-assets.com/authors/1230086720p2/49532.jpg",
		Content: ". . . when it comes down to it, that’s what life is all about: showing up for the people you love, again and again, until you can’t show up anymore.",
	},
	{
		Author: "Carl Sandburg",
		AuthorAvatar: "https://images.gr-assets.com/authors/1225504449p2/16380.jpg",
		Content: "I don't know where I'm going, but I'm on my way.",
	},
	{
		Author: "Mahmoud Darwish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1414535868p2/75055.jpg",
		Content: "في داخلي شُرْفَةٌلا يَمُرُّ بها أَحَدٌ للتَّحيَّة.",
	},
	{
		Author: "Maggie StiefvaterShiver",
		AuthorAvatar: "https://images.gr-assets.com/authors/1460918410p2/1330292.jpg",
		Content: "You're beautiful and sad,\" I said finally, not looking at him when I did. \"Just like your eyes. You're like a song that I heard when I was a little kid but forgot I knew until I heard it again.\" For a long moment there was only the whirring sound of the tires on the road, and then Sam said softly, \"Thank you.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Dare to live the life you have dreamed for yourself. Go forward and make your dreams come true.",
	},
	{
		Author: "Confucius",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407613261p2/15321.jpg",
		Content: "He who knows all the answers has not been asked all the questions.",
	},
	{
		Author: "Adeline Yen MahChinese Cinderella: The True Story of an Unwanted Daughter",
		AuthorAvatar: "https://images.gr-assets.com/authors/1232136046p2/3878.jpg",
		Content: "Please believe that one single positive dream is more important than a thousand negative realities.",
	},
	{
		Author: "F. Scott FitzgeraldThe Great Gatsby",
		AuthorAvatar: "https://images.gr-assets.com/authors/1517864008p2/3190.jpg",
		Content: "In my younger and more vulnerable years my father gave me some advice that I've been turning over in my mind ever since.\"Whenever you feel like criticizing any one,\" he told me, \"just remember that all the people in this world haven't had the advantages that you've had.",
	},
	{
		Author: "William ShakespeareA Midsummer Night's Dream",
		AuthorAvatar: "https://images.gr-assets.com/authors/1424313573p2/947.jpg",
		Content: "The course of true love never did run smooth.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "Death is so terribly final, while life is full of possibilities.",
	},
	{
		Author: "Kate Chopin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1224805370p2/5132.jpg",
		Content: "She wanted something to happen - something, anything: she did not know what.",
	},
	{
		Author: "Jack KerouacOn the Road",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "What is that feeling when you're driving away from people and they recede on the plain till you see their specks dispersing? - it's the too-huge world vaulting us, and it's good-bye. But we lean forward to the next crazy venture beneath the skies.",
	},
	{
		Author: "Marianne Moore",
		AuthorAvatar: "https://images.gr-assets.com/authors/1238338879p2/82947.jpg",
		Content: "Your thorns are the best part of you.",
	},
	{
		Author: "Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "Whatever you do, you need courage. Whatever course you decide upon, there is always someone to tell you that you are wrong. There are always difficulties arising that tempt you to believe your critics are right. To map out a course of action and follow it to an end requires some of the same courage that a soldier needs. Peace has its victories, but it takes brave men and women to win them.",
	},
	{
		Author: "John GreenPaper Towns",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "Here's what's not beautiful about it: from here, you can't see the rust or the cracked paint or whatever, but you can tell what the place really is. You can see how fake it all is. It's not even hard enough to be made out of plastic. It's a paper town. I mean, look at it, Q: look at all those culs-de-sac, those streets that turn in on themselves, all the houses that were built to fall apart. All those paper people living in their paper houses, burning the future to stay warm. All the paper kids drinking beer some bum bought for them at the paper convenience store. Everyone demented with the mania of owning things. All the things paper-thin and paper-frail. And all the people, too. I've lived here for eighteen years and I have never once in my life come across anyone who cares about anything that matters.",
	},
	{
		Author: "Jeanette Winterson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1406177070p2/9399.jpg",
		Content: "Book collecting is an obsession, an occupation, a disease, an addiction, a fascination, an absurdity, a fate. It is not a hobby. Those who do it must do it. Those who do not do it, think of it as a cousin of stamp collecting, a sister of the trophy cabinet, bastard of a sound bank account and a weak mind.",
	},
	{
		Author: "Irvine WelshPorno",
		AuthorAvatar: "https://images.gr-assets.com/authors/1461084249p2/5687.jpg",
		Content: "You can't lie to your soul.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Where is Wood?\" said Harry, suddenly realizing he wasn't there.\"Still in the showers,\" said Fred. \"We think he's trying to drown himself.",
	},
	{
		Author: "Friedrich Nietzsche",
		AuthorAvatar: "https://images.gr-assets.com/authors/1455294131p2/1938.jpg",
		Content: "Man is the cruelest animal.",
	},
	{
		Author: "Sharon Pollock",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397765649p2/339819.jpg",
		Content: "There is nobility in the struggle, you don't have to win.",
	},
	{
		Author: "Anaïs Nin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "There were always in me, two women at least, one woman desperate and bewildered, who felt she was drowning and another who would leap into a scene, as upon a stage, conceal her true emotions because they were weaknesses, helplessness, despair, and present to the world only a smile, an eagerness, curiosity, enthusiasm, interest.",
	},
	{
		Author: "Jamie McGuireBeautiful Disaster",
		AuthorAvatar: "https://images.gr-assets.com/authors/1479315727p2/4464118.jpg",
		Content: "You can’t tell mewhat to do anymore, Travis! I don’t belong to you!",
	},
	{
		Author: "Tamora PierceLady Knight",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209044273p2/8596.jpg",
		Content: "Threats are the last resort of a man with no vocabulary.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "How long was I asleep?\" she whispered. He didn't respond. \"How long was I asleep?\" she asked again, and noticed a hint of red in his cheeks. \"You were asleep, too?\" \"Until you began drooling on my shoulder.",
	},
	{
		Author: "George R.R. MartinA Game of Thrones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351944410p2/346732.jpg",
		Content: "When the snows fall and the white winds blow, the lone wolf dies but the pack survives.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I just did some calculations and I've been able to determine that you're full of shit.",
	},
	{
		Author: "Jonathan Safran FoerExtremely Loud and Incredibly Close",
		AuthorAvatar: "https://images.gr-assets.com/authors/1466172069p2/2617.jpg",
		Content: "There were things I wanted to tell him. But I knew they would hurt him. So I buried them, and let them hurt me.",
	},
	{
		Author: "Samuel Beckett",
		AuthorAvatar: "https://images.gr-assets.com/authors/1522050561p2/1433597.jpg",
		Content: "We are all born mad. Some remain so.",
	},
	{
		Author: "Sarah J. MaasThrone of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1269281353p2/3433047.jpg",
		Content: "My name is Celaena Sardothien. But it makes no difference if my name's Celaena or Lillian or Bitch, because I'd still beat you, no matter what you call me.",
	},
	{
		Author: "Stephenie MeyerNew Moon",
		AuthorAvatar: "https://images.gr-assets.com/authors/1477924095p2/941441.jpg",
		Content: "The bond forged between us was not one that could be broken by absence, distance, or time. And no matter how much more special or beautiful or brilliant or perfect than me he might be, he was as irreversibly altered as I was. As I would always belong to him, so would he always be mine.",
	},
	{
		Author: "Cassandra ClareCity of Glass",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "If you ever meet the man who could take advantage of Isabelle, you’ll have to let me know. I’d like to shake his hand. Or run away from him very fast, I’m not sure which.",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "If we have no peace, it is because we have forgotten that we belong to each other.",
	},
	{
		Author: "Eleanor Roosevelt",
		AuthorAvatar: "https://images.gr-assets.com/authors/1195764180p2/44566.jpg",
		Content: "No matter how plain a woman may be, if truth and honesty are written across her face, she will be beautiful.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "Just tell me how to be different in a way that makes sense.",
	},
	{
		Author: "Barack Obama",
		AuthorAvatar: "https://images.gr-assets.com/authors/1313512321p2/6356.jpg",
		Content: "The best way to not feel hopeless is to get up and do something. Don’t wait for good things to happen to you. If you go out and make some good things happen, you will fill the world with hope, you will fill yourself with hope.",
	},
	{
		Author: "Shel Silverstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201029128p2/435477.jpg",
		Content: "The VoiceThere is a voice inside of youThat whispers all day long,\"I feel this is right for me,I know that this is wrong.\"No teacher, preacher, parent, friendOr wise man can decideWhat's right for you--just listen toThe voice that speaks inside.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "In the middle of difficulty lies opportunity",
	},
	{
		Author: "Alexander McCall SmithThe No. 1 Ladies' Detective Agency",
		AuthorAvatar: "https://images.gr-assets.com/authors/1272995820p2/4738.jpg",
		Content: "You can go through life and make new friends every year - every month practically - but there was never any substitute for those friendships of childhood that survive into adult years. Those are the ones in which we are bound to one another with hoops of steel.",
	},
	{
		Author: "David LevithanThe Lover's Dictionary",
		AuthorAvatar: "https://images.gr-assets.com/authors/1426529210p2/11664.jpg",
		Content: "livid, adj.Fuck You for cheating on me. Fuck you for reducing it to the word cheating. As if this were a card game, and you sneaked a look at my hand. Who came up with the term cheating, anyway? A cheater, I imagine. Someone who thought liar was too harsh. Someone who thought devastator was too emotional. The same person who thought, oops, he’d gotten caught with his hand in the cookie jar. Fuck you. This isn’t about slipping yourself an extra twenty dollars of Monopoly money. These are our lives. You went and broke our lives. You are so much worse than a cheater. You killed something. And you killed it when its back was turned.",
	},
	{
		Author: "Stephen ChboskyThe Perks of Being a Wallflower",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199070302p2/12898.jpg",
		Content: "I just want you to know that you’re very special… and the only reason I’m telling you is that I don’t know if anyone else ever has.",
	},
	{
		Author: "Sylvia PlathThe Unabridged Journals of Sylvia Plath",
		AuthorAvatar: "https://images.gr-assets.com/authors/1373572652p2/4379.jpg",
		Content: "God, but life is loneliness, despite all the opiates, despite the shrill tinsel gaiety of \"parties\" with no purpose, despite the false grinning faces we all wear. And when at last you find someone to whom you feel you can pour out your soul, you stop in shock at the words you utter - they are so rusty, so ugly, so meaningless and feeble from being kept in the small cramped dark inside you so long. Yes, there is joy, fulfillment and companionship - but the loneliness of the soul in its appalling self-consciousness is horrible and overpowering.",
	},
	{
		Author: "Lemony SnicketThe Hostile Hospital",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "The sad truth is the truth is sad.",
	},
	{
		Author: "Charles DickensOur Mutual Friend",
		AuthorAvatar: "https://images.gr-assets.com/authors/1387078070p2/239579.jpg",
		Content: "And O there are days in this life, worth life and worth death.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "If you try to fail, and succeed, which have you done?",
	},
	{
		Author: "Gautama Buddha",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356856893p2/2167493.jpg",
		Content: "There is nothing more dreadful than the habit of doubt. Doubt separates people. It is a poison that disintegrates friendships and breaks up pleasant relations. It is a thorn that irritates and hurts; it is a sword that kills.",
	},
	{
		Author: "أحلام مستغانميcom نسيان",
		AuthorAvatar: "https://images.gr-assets.com/authors/1351458215p2/1109606.jpg",
		Content: "لا تقدم ابداً شروحاً لأحد.. أصدقاؤك الحقيقيون ليسوا فى حاجة إليها و أعداؤك لن يصدقوها",
	},
	{
		Author: "Rodney Dangerfield",
		AuthorAvatar: "https://images.gr-assets.com/authors/1208829061p2/78470.jpg",
		Content: "What a kid I got, I told him about the birds and the bees and he told me about the butcher and my wife.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "A question that sometimes drives me hazy: am I or are the others crazy?",
	},
	{
		Author: "Paulo CoelhoBrida",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "None of us knows what might happen even the next minute, yet still we go forward. Because we trust. Because we have Faith.",
	},
	{
		Author: "Susan Polis Schutz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1297527636p2/33395.jpg",
		Content: "Love is . . . Being happy for the other person when they are happy, Being sad for the person when they are sad, Being together in good times, And being together in bad times.LOVE IS THE SOURCE OF STRENGTH.Love is . . . Being honest with yourself at all times, Being honest with the other person at all times, Telling, listening, respecting the truth, And never pretending.LOVE IS THE SOURCE OF REALITY.Love is . . . An understanding so complete that you feel as if you are a part of the other person, Accepting the other person just the way they are, And not trying to change them to be something else.LOVE IS THE SOURCE OF UNITY.Love is . . . The freedom to pursue your own desires while sharing your experiences with the other person, The growth of one individual alongside of and together with the growth of another individual. LOVE IS THE SOURCE OF SUCCESS.Love is . . . The excitement of planning things together, The excitement of doing things together.LOVE IS THE SOURCE OF THE FUTURE.Love is . . . The fury of the storm, The calm in the rainbow.LOVE IS THE SOURCE OF PASSION.Love is . . . Giving and taking in a daily situation, Being patient with each other's needs and desires.LOVE IS THE SOURCE OF SHARING.Love is . . . Knowing that the other person will always be with you regardless of what happens, Missing the other person when they are away but remaining near in heart at all times.LOVE IS THE SOURCE OF SECURITY.LOVE IS . . . THE SOURCE OF LIFE!",
	},
	{
		Author: "James  JonesFrom Here to Eternity",
		AuthorAvatar: "https://images.gr-assets.com/authors/1227111357p2/3999.jpg",
		Content: "That was one of the virtues of being a pessimist: nothing was ever as bad as you thought it would be.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Order of the Phoenix",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "From now on, I don't care if my tea leaves spell 'Die, Ron, Die,' I'm chucking them in the bin where they belong.",
	},
	{
		Author: "Paulo CoelhoThe Alchemist",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201840056p2/566.jpg",
		Content: "This is what we call love. When you are loved, you can do anything in creation. When you are loved, there's no need at all to understand what's happening, because everything happens within you.",
	},
	{
		Author: "George Washington",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204292208p2/4356.jpg",
		Content: "It is better to offer no excuse than a bad one.",
	},
	{
		Author: "Taylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "To me, “FEARLESS",
	},
	{
		Author: "Albert CamusNotebooks 1935-1942",
		AuthorAvatar: "https://images.gr-assets.com/authors/1506091612p2/957894.jpg",
		Content: "Real generosity towards the future lies in giving all to the present.",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "It is always sad when someone leaves home, unless they are simply going around the corner and will return in a few minutes with ice-cream sandwiches.",
	},
	{
		Author: "Adolf Hitler",
		AuthorAvatar: "https://images.gr-assets.com/authors/1502617950p2/30691.jpg",
		Content: "If you win, you need not have to explain...If you lose, you should not be there to explain!",
	},
	{
		Author: "Anaïs NinThe Diary of Anaïs Nin, Vol. 1: 1931-1934",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468617897p2/7190.jpg",
		Content: "You live like this, sheltered, in a delicate world, and you believe you are living. Then you read a book… or you take a trip… and you discover that you are not living, that you are hibernating. The symptoms of hibernating are easily detectable: first, restlessness. The second symptom (when hibernating becomes dangerous and might degenerate into death): absence of pleasure. That is all. It appears like an innocuous illness. Monotony, boredom, death. Millions live like this (or die like this) without knowing it. They work in offices. They drive a car. They picnic with their families. They raise children. And then some shock treatment takes place, a person, a book, a song, and it awakens them and saves them from death. Some never awaken.",
	},
	{
		Author: "Galileo Galilei",
		AuthorAvatar: "https://images.gr-assets.com/authors/1225678415p2/14190.jpg",
		Content: "I have never met a man so ignorant that I couldn't learn something from him.",
	},
	{
		Author: "Cat Cora",
		AuthorAvatar: "https://images.gr-assets.com/authors/1459696345p2/398330.jpg",
		Content: "Even when you have doubts, take that step. Take chances. Mistakes are never a failure—they can be turned into wisdom.",
	},
	{
		Author: "Hilary MantelWolf Hall",
		AuthorAvatar: "https://images.gr-assets.com/authors/1334862633p2/58851.jpg",
		Content: "It is the absence of facts that frightens people: the gap you open, into which they pour their fears, fantasies, desires.",
	},
	{
		Author: "Epicurus",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430555926p2/114041.jpg",
		Content: "Do not spoil what you have by desiring what you have not; remember that what you now have was once among the things you only hoped for.",
	},
	{
		Author: "Maya Angelou",
		AuthorAvatar: "https://images.gr-assets.com/authors/1379017377p2/3503.jpg",
		Content: "Stepping onto a brand-new path is difficult, but not more difficult than remaining in a situation, which is not nurturing to the whole woman.",
	},
	{
		Author: "Marie Curie",
		AuthorAvatar: "https://images.gr-assets.com/authors/1212528295p2/126903.jpg",
		Content: "Nothing in life is to be feared, it is only to be understood. Now is the time to understand more, so that we may fear less.",
	},
	{
		Author: "Mahatma Gandhi",
		AuthorAvatar: "https://images.gr-assets.com/authors/1356810912p2/5810891.jpg",
		Content: "Each night, when I go to sleep, I die. And the next morning, when I wake up, I am reborn.",
	},
	{
		Author: "George Carlin",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214234683p2/22782.jpg",
		Content: "The main reason Santa is so jolly is because he knows where all the bad girls live.",
	},
	{
		Author: "Oscar WildeLady Windermere's Fan",
		AuthorAvatar: "https://images.gr-assets.com/authors/1521044377p2/3565.jpg",
		Content: "Life is far too important a thing ever to talk seriously about.",
	},
	{
		Author: "George Orwell1984",
		AuthorAvatar: "https://images.gr-assets.com/authors/1450573063p2/3706.jpg",
		Content: "It was a bright cold day in April, and the clocks were striking thirteen.",
	},
	{
		Author: "Rick RiordanThe Titan's Curse",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "He cleared his throat and held up one hand dramatically.“Green grass breaks through snow.\tArtemis pleads for my help.\tI am so cool.",
	},
	{
		Author: "Amelia Earhart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1252510868p2/367092.jpg",
		Content: "Adventure is worthwhile in itself.",
	},
	{
		Author: "Steven Moffat",
		AuthorAvatar: "https://images.gr-assets.com/authors/1307757543p2/812234.jpg",
		Content: "Demons run when a good man goes to warNight will fall and drown the sunWhen a good man goes to warFriendship dies and true love liesNight will fall and the dark will riseWhen a good man goes to warDemons run, but count the costThe battle's won, but the child is lost",
	},
	{
		Author: "Stéphane Mallarmé",
		AuthorAvatar: "https://images.gr-assets.com/authors/1491623064p2/5798517.jpg",
		Content: "Everything in the world exists in order to end up as a book.",
	},
	{
		Author: "Nicholas SparksSafe Haven",
		AuthorAvatar: "https://images.gr-assets.com/authors/1468247396p2/2345.jpg",
		Content: "Every couple needs to argue now and then. Just to prove that the relationship is strong enough to survive. Long-term relationships, the ones that matter, are all about weathering the peaks and the valleys.",
	},
	{
		Author: "Gilles DeleuzeA Thousand Plateaus: Capitalism and Schizophrenia",
		AuthorAvatar: "https://images.gr-assets.com/authors/1493078891p2/13009.jpg",
		Content: "A concept is a brick. It can be used to build a courthouse of reason. Or it can be thrown through the window.",
	},
	{
		Author: "Cassandra ClareClockwork Angel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "So you're a Shadowhunter,' Nate said. 'De Quincey told me that you lot were monsters.''Was that before or after he tried to eat you?' Will inquired.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "I used to believe in forever, but forever's too good to be true",
	},
	{
		Author: "Elizabeth GilbertEat, Pray, Love",
		AuthorAvatar: "https://images.gr-assets.com/authors/1440718929p2/11679.jpg",
		Content: "The Bhagavad Gita--that ancient Indian Yogic text--says that it is better to live your own destiny imperfectly than to live an imitation of somebody else's life with perfection.",
	},
	{
		Author: "Lorraine Hansberry",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234149147p2/3732.jpg",
		Content: "Never be afraid to sit awhile and think.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "I found myself thinking about President William McKinley, the third American president to be assassinated. He lived for several days after he was shot, and towards the end, his wife started crying and screaming, \"I want to go too! I want to go too!\" And with his last measure of strength, McKinley turned to her and spoke his last words: \"We are all going.",
	},
	{
		Author: "Jack Kerouac",
		AuthorAvatar: "https://images.gr-assets.com/authors/1430512644p2/1742.jpg",
		Content: "My fault, my failure, is not in the passions I have, but in my lack of control of them.",
	},
	{
		Author: "Ingmar Bergman",
		AuthorAvatar: "https://images.gr-assets.com/authors/1214845804p2/125406.jpg",
		Content: "Only someone who is well prepared has the opportunity to improvise.",
	},
	{
		Author: "John BanvilleThe Sea",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407383277p2/91.jpg",
		Content: "The past beats inside me like a second heart.",
	},
	{
		Author: "A.A. MilneWinnie-the-Pooh",
		AuthorAvatar: "https://images.gr-assets.com/authors/1204664899p2/81466.jpg",
		Content: "I'm not lost for I know where I am. But however, where I am may be lost.",
	},
	{
		Author: "Arthur Conan DoyleThe Boscombe Valley Mystery",
		AuthorAvatar: "https://images.gr-assets.com/authors/1495008883p2/2448.jpg",
		Content: "There is nothing more deceptive than an obvious fact.",
	},
	{
		Author: "Edith Sitwell",
		AuthorAvatar: "https://images.gr-assets.com/authors/1251402186p2/65646.jpg",
		Content: "Winter is the time for comfort, for good food and warmth, for the touch of a friendly hand and for a talk beside the fire: it is the time for home.",
	},
	{
		Author: "Robert Frost",
		AuthorAvatar: "https://images.gr-assets.com/authors/1183232004p2/7715.jpg",
		Content: "Half the world is composed of people who have something to say and can't, and the other half who have nothing to say and keep on saying it.",
	},
	{
		Author: "Sophie KinsellaThe Undomestic Goddess",
		AuthorAvatar: "https://images.gr-assets.com/authors/1245821549p2/6160.jpg",
		Content: "sometimes you don't need a goal in life, you don't need to know the big picture. you just need to know what you're going to do next!",
	},
	{
		Author: "Lemony SnicketHorseradish",
		AuthorAvatar: "https://images.gr-assets.com/authors/1199734355p2/36746.jpg",
		Content: "It is difficult, when faced with a situation you cannot control, to admit you can do nothing.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Prisoner of Azkaban",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "You think the dead we loved truly ever leave us? You think that we don't recall them more clearly in times of great trouble?",
	},
	{
		Author: "Charles M. Schulz",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198590750p2/209672.jpg",
		Content: "Sometimes I lie awake at night, and I ask, 'Where have I gone wrong'.Then a voice says to me, 'This is going to take more than one night.",
	},
	{
		Author: "Audrey NiffeneggerThe Time Traveler's Wife",
		AuthorAvatar: "https://images.gr-assets.com/authors/1367342548p2/498072.jpg",
		Content: "I won't ever leave you, even though you're always leaving me.",
	},
	{
		Author: "Bram Stoker",
		AuthorAvatar: "https://images.gr-assets.com/authors/1202438456p2/6988.jpg",
		Content: "Listen to them, the children of the night. What music they make!",
	},
	{
		Author: "Lloyd AlexanderThe Black Cauldron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353430382p2/8924.jpg",
		Content: "Child, child, do you not see? For each of us comes a time when we must be more than what we are.",
	},
	{
		Author: "Mae West",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198551937p2/259666.jpg",
		Content: "I'll try anything once, twice if I like it, three times to make sure.",
	},
	{
		Author: "Benjamin Alire SáenzAristotle and Dante Discover the Secrets of the Universe",
		AuthorAvatar: "https://images.gr-assets.com/authors/1386527200p2/4841310.jpg",
		Content: "Words were different when they lived inside of you.",
	},
	{
		Author: "Audrey Hepburn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1362614211p2/692403.jpg",
		Content: "I believe in being strong when everything seems to be going wrong. I believe that happy girls are the prettiest girls. I believe that tomorrow is another day, and I believe in miracles",
	},
	{
		Author: "Mother Teresa",
		AuthorAvatar: "https://images.gr-assets.com/authors/1263888735p2/838305.jpg",
		Content: "What can you do to promote world peace? Go home and love your family.",
	},
	{
		Author: "Rick RiordanThe Last Olympian",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "The world was collapsing, and the only thing that really mattered to me was that she was alive.",
	},
	{
		Author: "Coco Chanel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1340706964p2/3004479.jpg",
		Content: "In order to be irreplaceacle, one must always be different.",
	},
	{
		Author: "Yves Saint-Laurent",
		AuthorAvatar: "https://images.gr-assets.com/authors/1211910000p2/132725.jpg",
		Content: "The most beautiful makeup of a woman is passion. But cosmetics are easier to buy.",
	},
	{
		Author: "John GreenLooking for Alaska",
		AuthorAvatar: "https://images.gr-assets.com/authors/1353452301p2/1406384.jpg",
		Content: "That didn’t happen, of course. Things never happened the way I imagined them.",
	},
	{
		Author: "Amos Bronson AlcottConcord Days",
		AuthorAvatar: "https://images.gr-assets.com/authors/1244049004p2/138631.jpg",
		Content: "Stay is a charming word in a friend's vocabulary.",
	},
	{
		Author: "Vladimir Nabokov",
		AuthorAvatar: "https://images.gr-assets.com/authors/1482502806p2/5152.jpg",
		Content: "Our imagination flies -- we are its shadow on the earth.",
	},
	{
		Author: "Francis BaconThe Advancement Of Learning",
		AuthorAvatar: "https://images.gr-assets.com/authors/1397100712p2/50964.jpg",
		Content: "If a man will begin with certainties, he shall end in doubts; but if he will be content to begin with doubts, he shall end in certainties.",
	},
	{
		Author: "Haruki MurakamiNorwegian Wood",
		AuthorAvatar: "https://images.gr-assets.com/authors/1470611596p2/3354.jpg",
		Content: "I was always hungry for love. Just once, I wanted to know what it was like to get my fill of it -- to be fed so much love I couldn't take any more. Just once. ",
	},
	{
		Author: "Suzanne CollinsThe Hunger Games",
		AuthorAvatar: "https://images.gr-assets.com/authors/1394819770p2/153394.jpg",
		Content: "Kind people have a way of working their way inside me and rooting there.",
	},
	{
		Author: "Rick RiordanThe Battle of the Labyrinth",
		AuthorAvatar: "https://images.gr-assets.com/authors/1383677264p2/15872.jpg",
		Content: "Rachel: You're a half-blood, too?Annabeth: Shhh! Just announce it to the world, how about?Rachel: Okay. Hey, everybody! These two aren't human! They're half Greek god!...They don't seem to care.",
	},
	{
		Author: "J.K. RowlingHarry Potter and the Chamber of Secrets",
		AuthorAvatar: "https://images.gr-assets.com/authors/1510435123p2/1077326.jpg",
		Content: "Well, you're expelling us aren't you?\" said Ron.\"Not today, Mr. Weasley.\"Snape looked as though Christmas had been canceled.",
	},
	{
		Author: "Natalie BabbittTuck Everlasting",
		AuthorAvatar: "https://images.gr-assets.com/authors/1226361832p2/1954.jpg",
		Content: "Like all magnificent things, it's very simple.",
	},
	{
		Author: "Henry David Thoreau",
		AuthorAvatar: "https://images.gr-assets.com/authors/1392432620p2/10264.jpg",
		Content: "Dreams are the touchstones of our characters.",
	},
	{
		Author: "George Gordon Byron",
		AuthorAvatar: "https://images.gr-assets.com/authors/1357459330p2/44407.jpg",
		Content: "In secret we metIn silence I grieve,That thy heart could forget,Thy spirit deceive.",
	},
	{
		Author: "Harvey Pekar",
		AuthorAvatar: "https://images.gr-assets.com/authors/1278951274p2/5125.jpg",
		Content: "Ordinary life is pretty complex stuff.",
	},
	{
		Author: "Stephanie PerkinsAnna and the French Kiss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1407443106p2/3095893.jpg",
		Content: "I love you as certain dark things are loved, secretly, between the shadow and the soul.",
	},
	{
		Author: "Richelle MeadBlood Promise",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270374609p2/137902.jpg",
		Content: "Dimitri: \"Why did you come here?\"Rose: \"Because you hit me on the head and dragged me here.",
	},
	{
		Author: "David T.W. McCord",
		AuthorAvatar: "https://images.gr-assets.com/quotes/1405411782p2/1286369.jpg",
		Content: "Books fall open, you fall in.",
	},
	{
		Author: "Horton Foote",
		AuthorAvatar: "https://images.gr-assets.com/authors/1251494049p2/98565.jpg",
		Content: "I’ve known people that the world has thrown everything at to discourage them...to break their spirit. And yet something about them retains a dignity. They face life and don’t ask quarters.",
	},
	{
		Author: "Cassandra ClareCity of Bones",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "I forgot that's what gets you all hot and bothered, Jace, girls killing things.\"\"I like anyone killing things, especially me.\" he said with a smile.",
	},
	{
		Author: "Kathleen Winsor",
		AuthorAvatar: "https://images.gr-assets.com/authors/1234929032p2/3620.jpg",
		Content: "Charm is the ability to make someone else think that both of you are pretty wonderful.",
	},
	{
		Author: "Studs Terkel",
		AuthorAvatar: "https://images.gr-assets.com/authors/1445205508p2/33716.jpg",
		Content: "Work is about a search for daily meaning as well as daily bread, for recognition as well as cash, for astonishment rather than torpor; in short, for a sort of life rather than a Monday through Friday sort of dying.",
	},
	{
		Author: "Curtis SittenfeldPrep",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371752798p2/6429.jpg",
		Content: "I always worried someone would notice me, and then when no one did, I felt lonely.",
	},
	{
		Author: "Edward LearThe Owl and the Pussycat",
		AuthorAvatar: "https://images.gr-assets.com/authors/1209058926p2/142.jpg",
		Content: "And hand in hand, on the edge of the sand, They danced by the light of the moon.",
	},
	{
		Author: "Taylor SwiftTaylor Swift",
		AuthorAvatar: "https://images.gr-assets.com/authors/1371139510p2/1036517.jpg",
		Content: "Your lucky enough to be different, never change",
	},
	{
		Author: "Ted HughesLetters of Ted Hughes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1215068226p2/996.jpg",
		Content: "The only calibration that counts is how much heart people invest, how much they ignore their fears of being hurt or caught out or humiliated. And the only thing people regret is that they didn't live boldly enough, that they didn't invest enough heart, didn't love enough. Nothing else really counts at all.",
	},
	{
		Author: "Chuck PalahniukInvisible Monsters",
		AuthorAvatar: "https://images.gr-assets.com/authors/1391203076p2/2546.jpg",
		Content: "Parents are like God because you wanna know they're out there, and you want them to think well of you, but you really only call when you need something.",
	},
	{
		Author: "Dr. Seuss",
		AuthorAvatar: "https://images.gr-assets.com/authors/1193930952p2/61105.jpg",
		Content: "I’m glad we had the times together just to laugh and sing a song, seems like we just got started and then before you know it, the times we had together were gone.",
	},
	{
		Author: "Masashi KishimotoNaruto, Vol. 18: Tsunade's Choice",
		AuthorAvatar: "https://images.gr-assets.com/authors/1246487533p2/11047.jpg",
		Content: "She's strong! And scary...I bet she's single...I'd put money on it..",
	},
	{
		Author: "Charles Bukowski",
		AuthorAvatar: "https://images.gr-assets.com/authors/1501509674p2/13275.jpg",
		Content: "there is a place in the heart thatwill never be filleda spaceand even during thebest momentsandthe greatest timestimeswe will know itwe will know itmore thaneverthere is a place in the heart thatwill never be filledandwe will waitandwaitin that space.",
	},
	{
		Author: "Matt Groening",
		AuthorAvatar: "https://images.gr-assets.com/authors/1198682012p2/30808.jpg",
		Content: "I know all those words, but that sentence makes no sense to me.",
	},
	{
		Author: "Rainbow RowellFangirl",
		AuthorAvatar: "https://images.gr-assets.com/authors/1342324527p2/4208569.jpg",
		Content: "In new situations, all the trickiest rules are the ones nobody bothers to explain to you. (And the ones you can't Google.)",
	},
	{
		Author: "Cassandra ClareCity of Ashes",
		AuthorAvatar: "https://images.gr-assets.com/authors/1270502031p2/150038.jpg",
		Content: "Is there some particular reason that you're here?\" ...\"Not this again.\"\"Not what again?\" said Clary.\"Every time I annoy him, he retreats into his No Mundanes Allowed tree house.\" Simon pointed at Jace.",
	},
	{
		Author: "James Branch CabellThe Silver Stallion",
		AuthorAvatar: "https://images.gr-assets.com/authors/1207156655p2/92665.jpg",
		Content: "The optimist proclaims that we live in the best of all possible worlds; and the pessimist fears this is true.",
	},
	{
		Author: "Chris Abani",
		AuthorAvatar: "https://images.gr-assets.com/authors/1242037144p2/2928223.jpg",
		Content: "What I've come to learn is that the world is never saved in grand messianic gestures, but in the simple accumulation of gentle, soft, almost invisible acts of compassion.",
	},
	{
		Author: "Albert Einstein",
		AuthorAvatar: "https://images.gr-assets.com/authors/1429114964p2/9810.jpg",
		Content: "The pursuit of truth and beauty is a sphere of activity in which we are permitted to remain children all our lives.",
	},
	{
		Author: "Dennis LehaneSacred",
		AuthorAvatar: "https://images.gr-assets.com/authors/1227580381p2/10289.jpg",
		Content: "There's something ugly about the flawless.",
	},
	{
		Author: "Sue Grafton",
		AuthorAvatar: "https://images.gr-assets.com/authors/1319510564p2/9559.jpg",
		Content: "Ideas are easy. It's the execution of ideas that really separates the sheep from the goats.",
	},
	{
		Author: "Ralph Waldo EmersonThe Complete Prose Works of Ralph Waldo Emerson",
		AuthorAvatar: "https://images.gr-assets.com/authors/1393555704p2/12080.jpg",
		Content: "It is easy in the world to live after the world's opinion; it is easy in solitude to live after our own; but the great man is he who in the midst of the crowd keeps with perfect sweetness the independence of solitude.",
	},
	{
		Author: "Judith ViorstLove \u0026 Guilt \u0026 The Meaning Of Life, Etc",
		AuthorAvatar: "https://images.gr-assets.com/authors/1201030839p2/3080.jpg",
		Content: "Strength is the capacity to break a Hershey bar into four pieces with your bare hands - and then eat just one of the pieces.",
	},
	{
		Author: "Armistead MaupinMore Tales of the City",
		AuthorAvatar: "https://images.gr-assets.com/authors/1221811837p2/10022.jpg",
		Content: "Laugh all you want and cry all you want and whistle at pretty men in the street and to hell with anybody who thinks you're a damned fool!",
	},
	{
		Author: "Elinor Glyn",
		AuthorAvatar: "https://images.gr-assets.com/authors/1205206634p2/249153.jpg",
		Content: "Romance is the glamour which turns the dust of everyday life into a golden haze. ",
	},
	{
		Author: "Cornelia FunkeInkheart",
		AuthorAvatar: "https://images.gr-assets.com/authors/1437000100p2/15873.jpg",
		Content: "If you take a book with you on a journey,\" Mo had said when he put the first one in her box, \"an odd thing happens: The book begins collecting your memories. And forever after you have only to open that book to be back where you first read it. It will all come into your mind with the very first words: the sights you saw in that place, what it smelled like, the ice cream you ate while you were reading it... yes, books are like flypaper—memories cling to the printed page better than anything else.",
	},
}